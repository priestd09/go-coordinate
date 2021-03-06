// Copyright 2015-2016 Diffeo, Inc.
// This software is released under an MIT/X11 open source license.

package postgres

import (
	"database/sql"
	"fmt"
	"github.com/diffeo/go-coordinate/coordinate"
	"github.com/lib/pq"
	"time"
)

type attempt struct {
	unit   *workUnit
	worker *worker
	id     int
}

// Attempt interface

func (a *attempt) WorkUnit() coordinate.WorkUnit {
	return a.unit
}

func (a *attempt) Worker() coordinate.Worker {
	return a.worker
}

func (a *attempt) Status() (coordinate.AttemptStatus, error) {
	a.Coordinate().Expiry.Do(a)

	var status string
	err := withTx(a, true, func(tx *sql.Tx) error {
		return tx.QueryRow("SELECT status FROM attempt WHERE id=$1", a.id).Scan(&status)
	})
	if err == sql.ErrNoRows {
		err = coordinate.ErrGone
	}
	if err != nil {
		return 0, err
	}
	switch status {
	case "pending":
		return coordinate.Pending, nil
	case "expired":
		return coordinate.Expired, nil
	case "finished":
		return coordinate.Finished, nil
	case "failed":
		return coordinate.Failed, nil
	case "retryable":
		return coordinate.Retryable, nil
	}
	return 0, fmt.Errorf("invalid attempt status in database %v", status)
}

func (a *attempt) Data() (map[string]interface{}, error) {
	var result map[string]interface{}
	err := withTx(a, true, func(tx *sql.Tx) error {
		var dataBytes []byte
		row := tx.QueryRow("SELECT data FROM attempt WHERE id=$1", a.id)
		err := row.Scan(&dataBytes)
		if err == sql.ErrNoRows {
			err = coordinate.ErrGone
		}
		if err != nil {
			return err
		}
		if dataBytes == nil {
			// null data in the attempt; get the unmodified
			// work unit data
			row = tx.QueryRow("SELECT data FROM work_unit WHERE id=$1", a.unit.id)
			err = row.Scan(&dataBytes)
			if err == sql.ErrNoRows {
				err = coordinate.ErrGone
			}
			if err != nil {
				return err
			}
		}
		result, err = bytesToMap(dataBytes)
		return err
	})
	if err != nil {
		return nil, err
	}
	return result, err
}

func (a *attempt) StartTime() (result time.Time, err error) {
	err = withTx(a, true, func(tx *sql.Tx) error {
		return tx.QueryRow("SELECT start_time FROM attempt WHERE id=$1", a.id).Scan(&result)
	})
	if err == sql.ErrNoRows {
		err = coordinate.ErrGone
	}
	return
}

func (a *attempt) EndTime() (time.Time, error) {
	a.Coordinate().Expiry.Do(a)

	var nt pq.NullTime
	err := withTx(a, true, func(tx *sql.Tx) error {
		return tx.QueryRow("SELECT end_time FROM attempt WHERE id=$1", a.id).Scan(&nt)
	})
	if err == sql.ErrNoRows {
		err = coordinate.ErrGone
	}
	if err != nil {
		return time.Time{}, err
	}
	result := nullTimeToTime(nt)
	return result, nil
}

func (a *attempt) ExpirationTime() (result time.Time, err error) {
	a.Coordinate().Expiry.Do(a)

	err = withTx(a, true, func(tx *sql.Tx) error {
		return tx.QueryRow("SELECT expiration_time FROM attempt WHERE id=$1", a.id).Scan(&result)
	})
	if err == sql.ErrNoRows {
		err = coordinate.ErrGone
	}
	return
}

func (a *attempt) Renew(extendDuration time.Duration, data map[string]interface{}) error {
	// TODO(dmaze): check valid state and active status
	now := a.Coordinate().clock.Now()
	params := queryParams{}
	fields := fieldList{}
	fields.Add(&params, "expiration_time", now.Add(extendDuration))
	if data != nil {
		dataBytes, err := mapToBytes(data)
		if err != nil {
			return err
		}
		fields.Add(&params, "data", dataBytes)
	}
	query := buildUpdate(attemptTable, fields.UpdateChanges(), []string{
		isAttempt(&params, a.id),
	})
	return execInTx(a, query, params, true)
}

func (a *attempt) Expire(data map[string]interface{}) error {
	return withTx(a, false, func(tx *sql.Tx) error {
		return a.complete(tx, data, "expired")
	})
}

func (a *attempt) Finish(data map[string]interface{}) error {
	// Mark the attempt finished, then create any new work units
	// declared in an "output" key.
	//
	// These do not have to happen atomically.  So first, just mark
	// the attempt as done.
	err := withTx(a, false, func(tx *sql.Tx) error {
		return a.complete(tx, data, "finished")
	})
	if err != nil {
		return err
	}

	// A fast path: if we have a data dictionary and there is
	// no "output", stop.
	if data != nil {
		if _, present := data["output"]; !present {
			return nil
		}
	}

	// Otherwise we maybe have "output".  Do one query to the
	// database that gets back the work unit data (if we need it)
	// and the matching next work spec.  A join could fail, which
	// would result in nothing coming back, which would be okay.
	// This also depends on this attempt still being the active
	// attempt, which again, we can check in the query.
	params := queryParams{}
	outputs := []string{
		"next.id",
		"next.name",
	}
	tables := []string{
		workUnitTable,
		workSpecTable,
		workSpecTable + " next",
	}
	conditions := []string{
		isWorkUnit(&params, a.unit.id),
		workUnitHasAttempt(&params, a.id),
		workUnitInThisSpec,
		workSpecNextWorkSpec + "=next.name",
		workSpecNamespace + "=next.namespace_id",
	}
	if data == nil {
		// We need both the most recent attempt data and
		// the original unit data
		outputs = append(outputs, workUnitData, attemptData)
		tables = append(tables, attemptTable)
		conditions = append(conditions, attemptThisWorkUnit)
	}
	query := buildSelect(outputs, tables, conditions)
	spec := workSpec{namespace: a.unit.spec.namespace}
	err = withTx(a, true, func(tx *sql.Tx) (err error) {
		row := tx.QueryRow(query, params...)
		if data == nil {
			var unitData, attemptData []byte
			err = row.Scan(&spec.id, &spec.name, &unitData, &attemptData)
			if err == nil {
				if attemptData != nil {
					data, err = bytesToMap(attemptData)
				} else if unitData != nil {
					data, err = bytesToMap(unitData)
				} else {
					data = map[string]interface{}{}
				}
			}
		} else {
			err = row.Scan(&spec.id, &spec.name)
		}
		return
	})

	// Now, either that query failed, or we have both work unit
	// data and a next work spec.
	if err == sql.ErrNoRows {
		// As a reminder:
		// * a isn't the active attempt; or
		// * spec["then"] points nowhere
		// In any case, no outputs and we're done
		return nil
	}
	if err != nil {
		return err
	}

	units := coordinate.ExtractWorkUnitOutput(data["output"], a.Coordinate().clock.Now())
	if units == nil {
		return nil // nothing to do
	}
	for name, item := range units {
		var dataBytes []byte
		dataBytes, err = mapToBytes(item.Data)
		if err != nil {
			return err
		}
		_, err = spec.addWorkUnit(name, dataBytes, item.Meta)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *attempt) Fail(data map[string]interface{}) error {
	return withTx(a, false, func(tx *sql.Tx) error {
		return a.complete(tx, data, "failed")
	})
}

func (a *attempt) Retry(data map[string]interface{}, delay time.Duration) error {
	return withTx(a, false, func(tx *sql.Tx) error {
		err := a.complete(tx, data, "retryable")
		if err == nil {
			// Also update the "not before" time on the work unit
			then := a.Coordinate().clock.Now().Add(delay)
			params := queryParams{}
			fields := fieldList{}
			fields.Add(&params, "not_before", then)
			query := buildUpdate(workUnitTable,
				fields.UpdateChanges(),
				[]string{
					isWorkUnit(&params, a.unit.id),
				})
			_, err = tx.Exec(query, params...)
		}
		return err
	})
}

func (a *attempt) complete(tx *sql.Tx, data map[string]interface{}, status string) error {
	// TODO(dmaze): check valid state transitions
	// TODO(dmaze): check if attempt is active if required

	// Mark the attempt as completed
	params := queryParams{}
	fields := fieldList{}
	fields.AddDirect("active", "FALSE")
	fields.Add(&params, "status", status)
	fields.Add(&params, "end_time", a.Coordinate().clock.Now())
	if data != nil {
		dataBytes, err := mapToBytes(data)
		if err != nil {
			return err
		}
		fields.Add(&params, "data", dataBytes)
	}
	query := buildUpdate(attemptTable, fields.UpdateChanges(), []string{
		isAttempt(&params, a.id),
	})
	result, err := tx.Exec(query, params...)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err == nil && count == 0 {
		err = coordinate.ErrGone
	}
	if err != nil {
		return err
	}

	// If it was the active attempt, and this is a non-terminal
	// resolution, also reset that
	if status == "retryable" || status == "expired" {
		query = buildUpdate(workUnitTable, []string{
			"active_attempt_id=NULL",
		}, []string{
			"active_attempt_id=$1",
		})
		_, err = tx.Exec(query, a.id)
	}

	return err
}

// WorkUnit attempt functions

func (unit *workUnit) ActiveAttempt() (coordinate.Attempt, error) {
	unit.Coordinate().Expiry.Do(unit)
	w := worker{namespace: unit.spec.namespace}
	a := attempt{unit: unit, worker: &w}
	query := buildSelect([]string{
		"attempt.id",
		"worker.id",
		"worker.name",
	}, []string{
		"attempt",
		"worker",
		"work_unit",
	}, []string{
		"work_unit.id=$1",
		"attempt.id=work_unit.active_attempt_id",
		"worker.id=attempt.worker_id",
	})
	err := withTx(unit, true, func(tx *sql.Tx) error {
		return tx.QueryRow(query, unit.id).Scan(&a.id, &w.id, &w.name)
	})
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (unit *workUnit) ClearActiveAttempt() error {
	params := queryParams{}
	query := buildUpdate("work_unit", []string{
		"active_attempt_id=NULL",
	}, []string{
		isWorkUnit(&params, unit.id),
	})
	return execInTx(unit, query, params, true)
}

func (unit *workUnit) NumAttempts() (int, error) {
	num := 0
	var err error
	withTx(unit, true, func(tx *sql.Tx) error {
		num, err = unit.countAttempts(tx)
		return err
	})
	return num, err
}

func (unit *workUnit) Attempts() ([]coordinate.Attempt, error) {
	params := queryParams{}
	query := buildSelect([]string{
		attemptID,
		workerID,
		workerName,
	}, []string{
		attemptTable,
		workerTable,
	}, []string{
		attemptForUnit(&params, unit.id),
		attemptThisWorker,
	})
	var result []coordinate.Attempt
	err := queryAndScan(unit, query, params, func(rows *sql.Rows) error {
		w := worker{namespace: unit.spec.namespace}
		a := attempt{worker: &w, unit: unit}
		err := rows.Scan(&a.id, &w.id, &w.name)
		if err == nil {
			result = append(result, &a)
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (unit *workUnit) countAttempts(tx *sql.Tx) (int, error) {
	params := queryParams{}
	query := buildSelect(
		[]string{"COUNT(*)"},
		[]string{attemptTable},
		[]string{attemptForUnit(&params, unit.id)},
	)
	var count int
	err := tx.QueryRow(query, params...).Scan(&count)
	return count, err
}

// Worker attempt functions

func (w *worker) RequestAttempts(req coordinate.AttemptRequest) ([]coordinate.Attempt, error) {
	var (
		specs map[string]*workSpec
		metas map[string]*coordinate.WorkSpecMeta
		name  string
		err   error
		spec  *workSpec
		meta  *coordinate.WorkSpecMeta
	)

	// Run system-global expiry.
	w.Coordinate().Expiry.Do(w)

	// Collect the set of candidate work specs and metadata outside
	// the main transaction.  This is pretty expensive to collect
	// and we want to avoid retrying it if possible.
	//
	// There is a possible race condition on a bad day.  It is
	// possible that this returns work specs with positive
	// available units, but while we're deciding what to do,
	// another worker picks those up.  That means the scheduler
	// could pick something but we then fail to get any work from
	// it.
	for {
		err = withTx(w, true, func(tx *sql.Tx) (err error) {
			specs, metas, err = w.namespace.allMetas(tx, true)
			return
		})
		if err != nil {
			return nil, err
		}

		// Now pick something (this is stateless, but see TODO above)
		// (If this picks nothing, we're done)
		metas = coordinate.LimitMetasToNames(metas, req.WorkSpecs)
		metas = coordinate.LimitMetasToRuntimes(metas, req.Runtimes)
		now := w.Coordinate().clock.Now()
		name, err = coordinate.SimplifiedScheduler(metas, now, req.AvailableGb)
		if err == coordinate.ErrNoWork {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		spec = specs[name]
		meta = metas[name]

		// Then get some attempts
		attempts, err := w.requestAttemptsForSpec(req, spec, meta)
		if err != nil {
			return nil, err
		}

		// If that returned non-zero attempts, we're done
		if len(attempts) > 0 {
			result := make([]coordinate.Attempt, len(attempts))
			for i, a := range attempts {
				result[i] = a
			}
			return result, nil
		}
		// Otherwise reloop
	}
}

func (w *worker) requestAttemptsForSpec(
	req coordinate.AttemptRequest,
	spec *workSpec,
	meta *coordinate.WorkSpecMeta,
) ([]*attempt, error) {
	var (
		attempts []*attempt
		count    int
		err      error
	)

	// Adjust the work unit count based on what's possible here
	count = req.NumberOfWorkUnits
	if count < 1 {
		count = 1
	}
	if meta.MaxAttemptsReturned > 0 && count > meta.MaxAttemptsReturned {
		count = meta.MaxAttemptsReturned
	}
	if meta.MaxRunning > 0 && count > meta.MaxRunning-meta.PendingCount {
		count = meta.MaxRunning - meta.PendingCount
	}

	continuous := false
	length := time.Duration(15) * time.Minute
	err = withTx(w, false, func(tx *sql.Tx) error {
		var err error
		now := w.Coordinate().clock.Now()

		// Take an advisory lock for the work spec.  It doesn't
		// make sense for multiple concurrent actors to progress
		// beyond this point (they will hit the same work units
		// in chooseAndMakeAttempts() and one will roll back) and
		// this reduces the database load.  We are still protected
		// by standard SQL transactionality.
		params := queryParams{}
		query := "SELECT pg_advisory_xact_lock(0, " + params.Param(spec.id) + ")"
		_, err = tx.Exec(query, params...)
		if err != nil {
			return err
		}

		// Try to create attempts from pre-existing work units
		// (assuming we expect there to be some)
		if meta.AvailableCount > 0 {
			attempts, err = w.chooseAndMakeAttempts(
				tx, spec, count, now, length)
		}
		if err != nil || len(attempts) > 0 {
			return err
		}

		// If there were none, but the selected work spec is
		// continuous, maybe we can create a work unit and an
		// attempt
		if meta.CanStartContinuous(now) {
			var unit *workUnit
			var a *attempt
			continuous = true
			unit, err = w.createContinuousUnit(tx, spec, meta, now)
			if err == nil && unit != nil {
				a, err = makeAttempt(tx, unit, w, length)
			}
			if err == nil && a != nil {
				attempts = []*attempt{a}
			}
		}

		// Whatever happened, end of the road
		return err
	})
	// On a very bad day, we could have gone to create continuous
	// work units, but an outright INSERT failed from a duplicate
	// key.  If this happened, just pretend we didn't actually get
	// any attempts back, which will trigger the retry loop in our
	// caller.
	if continuous && err != nil && isDuplicateUnitName(err) {
		attempts = nil
		err = nil
	}
	// If we got attempts, but for a work spec with a max-retries
	// limit, recheck whether we need to fail some of those attempts.
	// (If this fails _some_ of the attempts, return less than the
	// maximum, that's okay; if this fails _all_ of the attempts,
	// that will cause RequestAttempts to try picking a work spec
	// again.)
	//
	// (Do this after the transaction above to cause the advisory
	// lock to get released and let other workers try to get work.
	// At this point we definitively do have attempts for these
	// work units, we just need to decide if we want to kill some
	// of them off preemptively.)
	if err == nil && meta.MaxRetries > 0 && len(attempts) > 0 {
		// At this point we _have_ the attempts.  (They are
		// committed in the database and everything.)  If
		// there is a database error at this point, it's
		// better to err on the side of returning them to the
		// caller and having them retried an extra time.
		_ = withTx(w, false, func(tx *sql.Tx) error {
			var err error
			attempts, err = w.maybeFailAttempts(
				tx, attempts, meta.MaxRetries)
			return err
		})
	}

	return attempts, err
}

// chooseMakeFailAttempts finds work units to do for a specific work
// spec, honoring the work spec's MaxRetries field.  It doesn't make
// sense to call this with maxRetries as zero; call
// chooseAndMakeAttempts instead.
func (w *worker) maybeFailAttempts(
	tx *sql.Tx,
	moreAttempts []*attempt,
	maxRetries int,
) ([]*attempt, error) {
	var attempts []*attempt
	// For each of the (new) attempts, count the number of
	// existing attempts for the work unit and maybe fail it.
	// (It might be nice to do this in a batch?)
	for _, a := range moreAttempts {
		count, err := a.unit.countAttempts(tx)
		if err != nil {
			return nil, err
		}
		if count > maxRetries {
			err = a.complete(tx,
				map[string]interface{}{
					"traceback": "too many retries",
				},
				"failed")
			if err != nil {
				return nil, err
			}
			continue
			// and drop this attempt
		}
		attempts = append(attempts, a)
	}
	return attempts, nil
}

// chooseAndMakeAttempts, in one SQL query, finds work units to do for
// a specific work spec, creates attempts for them, and returns the
// corresponding attempt objects.
func (w *worker) chooseAndMakeAttempts(
	tx *sql.Tx,
	spec *workSpec,
	numUnits int,
	now time.Time,
	length time.Duration,
) ([]*attempt, error) {
	params := queryParams{}

	choose := buildSelect([]string{
		workUnitID,
		workUnitName,
	}, []string{
		workUnitTable,
	}, []string{
		workUnitInSpec(&params, spec.id),
		workUnitHasNoAttempt,
		"NOT " + workUnitTooSoon(&params, now),
	})
	choose += " ORDER BY priority DESC, name ASC"
	choose += fmt.Sprintf(" LIMIT %v", numUnits)

	expiration := now.Add(length)
	whatToInsert := buildSelect([]string{
		"id",
		params.Param(spec.id),
		params.Param(w.id),
		params.Param(now),
		params.Param(expiration),
	}, []string{
		"u",
	}, []string{})
	attempts := "INSERT INTO " + attemptTable +
		"(work_unit_id, work_spec_id, worker_id, start_time, expiration_time) " +
		whatToInsert +
		" RETURNING id, work_unit_id"

	update := "UPDATE " + workUnitTable + " " +
		"SET active_attempt_id=a.id " +
		"FROM a, u " +
		"WHERE " + workUnitID + "=u.id AND a.work_unit_id=u.id " +
		"RETURNING u.id, u.name, a.id"

	query := "WITH u AS (" + choose + "), a AS (" + attempts + ") " + update

	rows, err := tx.Query(query, params...)
	if err != nil {
		return nil, err
	}
	var result []*attempt
	err = scanRows(rows, func() error {
		unit := workUnit{spec: spec}
		attempt := attempt{unit: &unit, worker: w}
		if err := rows.Scan(&unit.id, &unit.name, &attempt.id); err == nil {
			result = append(result, &attempt)
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// createContinuousUnit tries to create exactly one continuous work
// unit, and returns it.
func (w *worker) createContinuousUnit(tx *sql.Tx, spec *workSpec, meta *coordinate.WorkSpecMeta, now time.Time) (*workUnit, error) {
	// Ideally only one worker is attempting to create this work
	// unit.  Start by updating the next-continuous time; this
	// creates a modification that will limit other transactions.
	params := queryParams{}
	fields := fieldList{}
	fields.Add(&params, "next_continuous", now.Add(meta.Interval))
	query := buildUpdate(workSpecTable,
		fields.UpdateChanges(),
		[]string{
			isWorkSpec(&params, spec.id),
			workSpecNotTooSoon(&params, now),
		})
	res, err := tx.Exec(query, params...)
	if err != nil {
		return nil, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		// This probably means somebody else advanced the
		// next-continuous time into the future under us.  Not
		// an error, just stop (and reloop in the caller).
		return nil, nil
	}

	// We should consider repeating the query at this point to
	// ensure there is no available or pending work unit for the
	// work spec.  (Or roll it into the query side of the
	// preceding UPDATE.)  Especially if meta.Interval is zero,
	// this can protect against multiple continuous units being
	// created on top of each other.

	// Create the work unit
	seconds := now.Unix()
	nano := now.Nanosecond()
	milli := nano / 1000000
	name := fmt.Sprintf("%d.%03d", seconds, milli)
	dataBytes, err := mapToBytes(map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	// The unit *shouldn't* exist, though on a bad day with lots
	// of workers it could.  We will probably see the consistency
	// error in the commit.  :-/
	unit, err := spec.insertWorkUnit(tx, name, dataBytes, coordinate.WorkUnitMeta{})
	if err != nil {
		return nil, err
	}

	// We have the single work unit we want to do.
	return unit, nil
}

func (w *worker) MakeAttempt(cUnit coordinate.WorkUnit, length time.Duration) (coordinate.Attempt, error) {
	unit, ok := cUnit.(*workUnit)
	if !ok {
		return nil, coordinate.ErrWrongBackend
	}
	var a *attempt
	var err error
	err = withTx(w, false, func(tx *sql.Tx) error {
		a, err = makeAttempt(tx, unit, w, length)
		return err
	})
	if err != nil {
		return nil, err
	}
	return a, nil
}

func makeAttempt(tx *sql.Tx, unit *workUnit, w *worker, length time.Duration) (*attempt, error) {
	a := attempt{unit: unit, worker: w}

	now := a.Coordinate().clock.Now()
	expiration := now.Add(length)
	params := queryParams{}
	fields := fieldList{}
	fields.Add(&params, "work_unit_id", unit.id)
	fields.Add(&params, "work_spec_id", unit.spec.id)
	fields.Add(&params, "worker_id", w.id)
	fields.Add(&params, "start_time", now)
	fields.Add(&params, "expiration_time", expiration)
	query := fields.InsertStatement(attemptTable) + " RETURNING id"
	row := tx.QueryRow(query, params...)
	err := row.Scan(&a.id)
	if err != nil {
		return nil, err
	}

	params = queryParams{}
	fields = fieldList{}
	fields.Add(&params, "active_attempt_id", a.id)
	query = buildUpdate(workUnitTable, fields.UpdateChanges(), []string{
		isWorkUnit(&params, unit.id),
	})
	_, err = tx.Exec(query, params...)

	return &a, err
}

func (w *worker) ActiveAttempts() ([]coordinate.Attempt, error) {
	qp := queryParams{}
	return w.findAttempts([]string{
		attemptByWorker(&qp, w.id),
		attemptIsActive,
	}, &qp, false)
}

func (w *worker) AllAttempts() ([]coordinate.Attempt, error) {
	qp := queryParams{}
	return w.findAttempts([]string{
		attemptByWorker(&qp, w.id),
	}, &qp, false)
}

func (w *worker) ChildAttempts() ([]coordinate.Attempt, error) {
	qp := queryParams{}
	return w.findAttempts([]string{
		attemptThisWorker,
		attemptIsActive,
		workerHasParent(&qp, w.id),
	}, &qp, true)
}

func (w *worker) findAttempts(conditions []string, qp *queryParams, forOtherWorkers bool) ([]coordinate.Attempt, error) {
	outputs := []string{
		attemptID,
		workUnitID,
		workUnitName,
		workSpecID,
		workSpecName,
	}
	tables := []string{
		attemptTable,
		workUnitTable,
		workSpecTable,
	}
	conditions = append(conditions,
		attemptThisWorkUnit,
		workUnitInThisSpec,
	)
	if forOtherWorkers {
		outputs = append(outputs, workerID, workerName)
		tables = append(tables, workerTable)
		conditions = append(conditions, attemptThisWorker)
	}
	query := buildSelect(outputs, tables, conditions)
	var result []coordinate.Attempt
	err := queryAndScan(w, query, *qp, func(rows *sql.Rows) error {
		spec := workSpec{namespace: w.namespace}
		unit := workUnit{spec: &spec}
		a := attempt{worker: w, unit: &unit}
		theWorker := worker{namespace: w.namespace}
		var err error
		if forOtherWorkers {
			a.worker = &theWorker
			err = rows.Scan(&a.id,
				&unit.id, &unit.name,
				&spec.id, &spec.name,
				&theWorker.id, &theWorker.name)
		} else {
			err = rows.Scan(&a.id,
				&unit.id, &unit.name,
				&spec.id, &spec.name)
		}
		if err == nil {
			result = append(result, &a)
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// coordinable interface

func (a *attempt) Coordinate() *pgCoordinate {
	return a.worker.namespace.coordinate
}
