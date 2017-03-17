// Code generated by go-bindata.
// sources:
// migrations/20150927-core.sql
// migrations/20151002-mingb.sql
// migrations/20151006-work-unit-type.sql
// migrations/20151013-index.sql
// migrations/20151014-index.sql
// migrations/20151019-worker-mode.sql
// migrations/20151028-index.sql
// migrations/20151216-work-spec-runtime.sql
// migrations/20160104-not-before.sql
// migrations/20160125-index.sql
// migrations/20160217-attempt-spec.sql
// migrations/20160328-index.sql
// migrations/20160329-index.sql
// migrations/20170316-index.sql
// DO NOT EDIT!

package postgres

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _migrations20150927CoreSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x96\xcd\x72\xe2\x38\x10\xc7\xef\x7e\x8a\xbe\x01\xbb\xe1\x05\xc2\x49\xc1\x4a\xe2\x5a\x63\xb3\xb2\x9c\x2d\xf6\xe2\x52\x70\x87\xa8\x06\x64\x47\x96\xf3\xf1\xf6\x53\x36\x18\x1b\x7f\x30\x90\xaa\x39\x8c\x0f\x94\x91\x5a\xfd\xf1\xd7\xaf\x1b\xa6\x53\x98\xfe\x35\x85\x5d\x12\xe3\x2d\x64\x6f\xdb\x59\xf1\x31\x4d\x75\x12\xe7\x6b\x73\x0b\x69\x92\x99\x8d\xc6\xac\x30\xb2\xa6\x53\xf8\x7b\x27\x37\x5a\x18\x84\x30\xb5\xe6\x8c\x12\x4e\x81\x93\x3b\x97\x82\x12\x3b\xcc\x52\xb1\xc6\xb1\x05\xfb\x47\xc6\x10\x50\xe6\x10\x17\x96\xcc\x59\x10\xb6\x82\x7f\xe8\xea\xa6\xda\x2d\xec\xe1\x89\xb0\xf9\x23\x61\x10\x7a\xce\xbf\x21\x05\xcf\xe7\xe0\x85\xae\x6b\x4d\x66\xd6\xa9\xf7\x8f\x44\xff\x88\xb2\x14\xd7\x57\x78\x2f\xb3\x89\x64\x0c\x8e\xc7\xe9\x03\x65\xb5\x7b\xe8\x79\x18\xbd\xa7\x8c\x7a\x73\x1a\x34\x4a\x91\xf1\x04\x7c\x0f\x6c\xea\x52\x4e\x61\x4e\x82\x39\xb1\x69\x7f\x09\x95\xf3\xe3\x6e\x2c\x8c\x80\xbb\x15\xa7\xa4\xbb\x97\x6a\x99\x68\x69\xbe\x3a\xa9\x1d\x2d\x3e\x50\x6e\x5e\xcd\xf0\x7e\x2a\xf2\x0c\x63\xb8\xf3\x7d\x97\x12\xaf\xbb\xbf\x4e\x94\x91\x2a\x4f\xf2\xec\x8c\x8d\x50\xd1\x33\x46\x97\x98\x4a\x65\x50\xbf\x8b\x6d\x99\x10\x7b\x22\x6e\xd7\x44\xe1\xa7\x69\xfa\xe2\xce\x82\x06\x9c\x2c\x96\xf0\x9f\xc3\x1f\xcb\xaf\xf0\xbf\xef\xd5\xf2\xed\xc4\x67\xa4\x73\xa5\xa4\xda\x0c\xd7\x59\x18\x09\x63\x70\x97\x9a\x2c\xd2\x68\x72\xad\xb0\x7b\xa3\xa7\x49\x1c\x61\x89\xce\xdf\x50\xcb\x38\xd5\x58\x86\x19\x56\x61\xee\x7b\x01\x67\xc4\xf1\x78\x0d\x64\x94\x2b\xf9\x96\xe3\x3e\xd4\x9e\xe3\x71\x93\xbe\x9b\x12\x93\xc9\x00\xd2\xb9\x92\xe6\x42\xa4\xeb\x88\xdf\x40\xba\xee\x9f\xdf\x8e\xb4\xed\x87\x45\x79\x4b\x46\xe7\x4e\xe0\xf8\x3d\x32\x8a\xb5\x91\xef\x58\x5d\x6b\xa3\x9c\x41\xa1\x0b\x99\xfa\x84\x6e\x6a\x72\x56\x68\xd4\x7f\xce\xe0\x48\x85\x46\x55\xb7\x7e\xeb\x16\x51\xb7\x9c\x07\xb4\x57\xde\x61\x8a\x8b\x51\x3f\xdc\x41\xe7\xee\x18\x3f\x53\xa9\x85\x91\x89\x1a\xec\xee\xee\xa1\xad\xc8\x4c\x94\xa7\x71\xf1\xa3\x71\xf9\xa9\x16\x01\xa8\xbf\xd3\x67\xab\x25\x85\x8a\xb2\xcc\x08\x93\x67\x40\x02\xa0\x5e\xb8\x18\x8f\x52\x54\xb1\x54\x9b\xd1\x0d\x8c\xca\xb2\x30\x2e\x5e\x5f\xa4\x92\xd9\x6b\xf1\x5e\x25\x32\x7a\x11\x72\xbb\xdf\xd4\x68\xf4\x97\x78\xde\xe2\xa8\xc3\xd8\x21\xca\x35\xad\x5c\x32\xfd\xdd\x56\x2e\xe7\xc6\x79\xc8\x0e\xb2\x5d\x14\xe1\x57\x8c\xb5\x7d\x1f\xc4\x6c\x69\x5b\xf9\x07\x9b\xde\x93\xd0\xe5\x50\x8b\xdc\xa5\xab\xe9\x4b\x9b\xc8\xc8\xdd\x35\x78\xa0\x8a\xcf\x1f\xe9\x61\xf6\xda\x18\x03\x7d\x74\xac\x8e\xb3\x90\x96\xb8\x11\x97\x53\xd6\x9e\xea\x07\x2f\xc4\xb6\xfb\xa7\x59\x6b\x08\xbe\x8b\xad\x8c\x0f\x67\xee\x7d\x46\x9d\x07\xaf\x40\x06\xc6\x9d\x61\x39\x69\xde\x56\xc5\x9d\x8c\x27\x87\xc3\xdd\xc1\x30\xb3\x4e\xfe\xb8\xd9\xc9\x87\xea\x4f\x19\x6c\xe6\x2f\x2f\xcf\x76\x66\x95\xf6\x27\x0d\x50\xad\x75\x3b\xef\xc4\x7a\x0f\x59\x67\xa9\x8c\xd5\x5d\x2d\x06\xfc\xc9\xea\xb1\xef\x67\xd6\xcf\x00\x00\x00\xff\xff\x3d\x99\xdb\xbf\xbf\x0a\x00\x00")

func migrations20150927CoreSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20150927CoreSql,
		"migrations/20150927-core.sql",
	)
}

func migrations20150927CoreSql() (*asset, error) {
	bytes, err := migrations20150927CoreSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20150927-core.sql", size: 2751, mode: os.FileMode(420), modTime: time.Unix(1487206567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20151002MingbSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\xcd\xc1\x0a\xc2\x20\x00\x87\xf1\xbb\x4f\xf1\x3f\x17\x42\xe7\xed\x64\xd3\x60\x60\x3a\x96\x9e\x47\x6d\x32\x46\x39\x4d\x17\xa3\xb7\x8f\x1d\x83\xba\x7c\xa7\x0f\x7e\x94\x82\xee\x28\x7c\x18\x5c\x81\xfc\x7c\x94\x5b\x68\x4c\x61\x78\xf5\x4b\x81\x18\xf2\x32\x26\x97\xb7\x89\x50\x8a\xbd\x9f\xc6\x74\x5d\x1c\x6c\x24\x4c\x1a\xd1\xc2\xb0\xa3\x14\x58\x43\xba\x77\x39\xba\x1e\x8c\x73\x54\x5a\xda\xb3\x82\x9f\xe6\xce\x3b\x1f\xd2\xbb\x1b\x6f\xe0\xda\x6e\x67\xd3\x8a\xaa\xbe\xd4\x5a\x41\x69\x03\x65\xa5\x04\x17\x27\x66\xa5\xc1\xa1\x24\x5f\x06\x0f\xeb\xfc\x47\xe1\xad\x6e\x7e\x32\x25\xf9\x04\x00\x00\xff\xff\x98\xf6\x73\x37\xd2\x00\x00\x00")

func migrations20151002MingbSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20151002MingbSql,
		"migrations/20151002-mingb.sql",
	)
}

func migrations20151002MingbSql() (*asset, error) {
	bytes, err := migrations20151002MingbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20151002-mingb.sql", size: 210, mode: os.FileMode(420), modTime: time.Unix(1487206567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20151006WorkUnitTypeSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x90\x4f\x4f\xc2\x40\x10\xc5\xef\xfd\x14\xef\x86\xff\x96\x33\x81\xd3\x02\x45\x4d\x2a\x9a\xb2\x35\xe1\x44\x6a\x3b\xb6\x8d\xee\x4e\xdd\x9d\x8a\x7c\x7b\xb3\x60\x62\x34\x9e\x9c\xc3\x1c\x26\xef\xcd\xef\xe5\x29\x05\x75\xa1\x60\xb9\xa6\x29\xc2\xdb\xeb\x2c\x2e\xd5\x7b\xae\x87\x4a\xa6\xe8\x39\x48\xe3\x29\x44\x51\xa2\x14\x4c\xdb\x05\x54\x6d\xe9\x1a\x0a\x90\x96\xb0\x67\xff\xb2\x1b\x5c\x27\x63\x57\x5a\x82\x1c\x7a\xc2\xb3\x67\x8b\x47\x9d\x2f\x6e\x74\x0e\x61\xcc\xb7\x26\xd5\x63\x40\xbb\x03\xe8\xa3\x0b\xd2\xb9\x26\x3e\x8b\x5e\x44\x6f\x40\xe9\x09\x65\x08\x83\xa5\x3a\x3a\x9e\x08\x85\x59\xa9\x09\xc8\x55\x5c\x53\x3d\x4e\xa2\xfe\xd2\x76\x8d\x2f\x85\x50\xf4\x89\xce\x4c\x9a\xc3\xe8\x79\x96\x7e\x67\xc0\xe9\xba\xb8\xcf\x8a\xbb\x35\x8e\x81\x36\xa9\xc1\x52\x1b\x0d\xb3\x7d\x48\x4f\x49\x12\x1c\xa7\xd8\xdc\xae\xaf\x51\xb1\x7b\x27\x2f\x3b\xe1\xb3\xa8\xbf\xc2\xa8\x30\xab\xc9\xe8\x7c\xf6\x93\xb8\xe4\xbd\xfb\x27\xf3\xab\x88\x3f\xa9\xb1\xa9\xdf\xdc\xcf\x00\x00\x00\xff\xff\xd4\xb3\x9b\xee\x92\x01\x00\x00")

func migrations20151006WorkUnitTypeSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20151006WorkUnitTypeSql,
		"migrations/20151006-work-unit-type.sql",
	)
}

func migrations20151006WorkUnitTypeSql() (*asset, error) {
	bytes, err := migrations20151006WorkUnitTypeSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20151006-work-unit-type.sql", size: 402, mode: os.FileMode(420), modTime: time.Unix(1487206567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20151013IndexSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\xcc\xb1\xce\x82\x30\x14\xc5\xf1\xbd\x4f\x71\xc6\xef\x53\xef\x0b\xc0\x44\x68\x07\x17\x30\xa0\x89\x1b\x21\xd2\x90\x46\xdb\x5b\x6f\x4b\x88\x6f\x6f\x98\x8c\x83\xcb\x19\x4e\xfe\xf9\x11\x81\x76\x04\xcf\x93\x2d\x90\x9e\x8f\x72\x1b\x8a\xc2\xd3\x72\xcb\x05\x22\xa7\x3c\x8b\x4d\x5b\xa4\x88\xb0\xf7\x6e\x96\x31\x5b\x5c\xa2\xaa\x3b\x53\x9d\x0d\x8e\x8d\x36\x57\xac\x2c\xf7\x61\x09\x2e\x0f\x2c\x93\x15\x17\x66\xb4\xcd\xe7\xfd\x8b\xe2\x58\x5c\x7e\x41\x9b\xbe\x3e\x20\x8c\xde\xa2\xea\xeb\xff\x52\x7d\xb1\x9a\xd7\xa0\x74\xd7\x9e\x7e\xb2\xa5\x7a\x07\x00\x00\xff\xff\x3f\x3c\x19\x8a\xb3\x00\x00\x00")

func migrations20151013IndexSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20151013IndexSql,
		"migrations/20151013-index.sql",
	)
}

func migrations20151013IndexSql() (*asset, error) {
	bytes, err := migrations20151013IndexSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20151013-index.sql", size: 179, mode: os.FileMode(420), modTime: time.Unix(1487206567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20151014IndexSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd5\x55\xd0\xd5\xd2\x55\xc8\xcd\x4f\x49\xb5\x52\x28\x2e\xcc\xb1\x06\x11\xba\x05\x45\xf9\x29\xa5\xc9\x25\x56\x0a\x05\xf9\xc5\x25\xe9\x45\xa9\xc5\x20\x45\x5c\xba\xba\x0a\xda\xb9\x99\xe9\x45\x89\x25\xa9\x0a\xa1\x05\x5c\xce\x41\xae\x8e\x21\xae\x0a\x9e\x7e\x2e\xae\x11\x0a\xe5\xf9\x45\xd9\xf1\xa5\x79\x99\x25\xf1\xc5\x05\xa9\xc9\x0a\xfe\x7e\x08\x11\x0d\x30\x0b\x24\x1c\x9f\x99\xa2\x69\x8d\x4f\x5f\x7c\x62\x49\x49\x6a\x6e\x41\x09\x6e\xfd\x3a\x0a\x89\xc9\x25\x99\x65\xa9\x30\x95\x10\x23\x51\x9c\xe6\x92\x5f\x9e\xc7\xe5\x12\xe4\x1f\x80\xd7\x0a\x6b\xdc\x4a\xac\xb9\x00\x01\x00\x00\xff\xff\x75\x53\xe1\x4c\x16\x01\x00\x00")

func migrations20151014IndexSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20151014IndexSql,
		"migrations/20151014-index.sql",
	)
}

func migrations20151014IndexSql() (*asset, error) {
	bytes, err := migrations20151014IndexSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20151014-index.sql", size: 278, mode: os.FileMode(420), modTime: time.Unix(1487206567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20151019WorkerModeSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\xcd\x41\x4b\xc3\x40\x10\x05\xe0\x7b\x7e\xc5\xbb\x15\x94\x2d\x9e\x9b\xd3\xda\x2e\xb5\x10\xa3\xa4\x1b\xc1\x63\xc8\xae\x9b\x45\xb3\x1b\x67\x46\xea\xcf\x97\xb5\x16\xf4\x2a\x0c\x73\x18\xde\x7c\x4f\x29\xa8\x2b\x85\x39\x3b\xbf\x01\xbf\xbf\xd5\x65\xa9\x85\xb2\xfb\x18\x65\x83\x25\xb3\x04\xf2\x5c\x42\x95\x2a\x03\x3b\x45\xc6\x38\x0d\x29\x78\x86\x4c\x1e\xa7\x4c\xaf\x9e\xbe\x09\xbc\x50\x9e\x11\x93\x40\x32\x58\x28\xa6\xb0\x06\xcc\x67\x64\x89\x29\xfc\x4e\x72\xa1\x06\xf2\x70\x91\xc7\x81\x9c\x77\xeb\x1f\xff\x7a\x8e\x81\x06\xf1\xe8\x97\x4a\x37\xd6\x74\xb0\xfa\xb6\x31\x97\xe7\xf3\x69\xfb\xd0\xf4\xf7\xed\xb9\xf3\x68\x2c\x76\xda\x6a\xd8\xe7\x47\x83\x27\xdd\x6d\xef\x74\x87\xfe\x78\x68\xf7\x58\xad\xea\xea\x0f\xba\xcb\xa7\xf4\x1f\xf6\xd0\x5a\xb3\x37\x17\xf6\xa6\xae\xbe\x02\x00\x00\xff\xff\x85\x69\xfb\x23\x39\x01\x00\x00")

func migrations20151019WorkerModeSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20151019WorkerModeSql,
		"migrations/20151019-worker-mode.sql",
	)
}

func migrations20151019WorkerModeSql() (*asset, error) {
	bytes, err := migrations20151019WorkerModeSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20151019-worker-mode.sql", size: 313, mode: os.FileMode(420), modTime: time.Unix(1487206567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20151028IndexSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x90\xcd\xae\x82\x30\x10\x85\xf7\x7d\x8a\x59\x72\x7f\xfa\x02\x74\x75\x73\x61\xe1\x06\x0c\xd1\xc4\x5d\xd3\xc0\x84\x34\x5a\x5a\xdb\x41\x7c\x7c\xd3\x20\x0a\x8a\x9b\x59\x9c\xef\x9c\x93\x99\xe1\x1c\xf8\x37\x07\x63\x1b\x4c\x21\x9c\x4f\x22\x0e\xee\xbc\x6d\xfa\x9a\x52\x70\x36\x50\xeb\x31\x44\x13\xe3\x1c\x7e\x8c\x6e\xbd\x22\x84\xbd\x63\xff\x55\xfe\xb7\xcb\x61\x53\x64\xf9\x01\x14\x11\x1a\x47\x32\x90\xa2\x3e\x48\xbc\x3a\xed\x15\x69\xdb\x41\x59\x4c\x30\x19\xe1\x2f\x3c\xa9\x24\x6d\xf0\x4b\xac\x77\x0d\xd6\x1f\xd1\xcf\x0b\x46\x45\xea\xe6\x35\x12\x81\xec\x3b\x4d\xf2\x6e\x8d\xa9\x87\x98\xa8\x9a\xf4\x05\x27\x36\xe6\x17\xe7\x64\x76\xe8\x58\x56\x95\xdb\x4f\x7d\x62\x4e\x97\x0b\xae\xa2\xb7\x3f\x08\x76\x0b\x00\x00\xff\xff\x0b\x68\xa2\x93\x6b\x01\x00\x00")

func migrations20151028IndexSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20151028IndexSql,
		"migrations/20151028-index.sql",
	)
}

func migrations20151028IndexSql() (*asset, error) {
	bytes, err := migrations20151028IndexSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20151028-index.sql", size: 363, mode: os.FileMode(420), modTime: time.Unix(1487206567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20151216WorkSpecRuntimeSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\xcc\xc1\x6a\xc4\x20\x14\x85\xe1\xbd\x4f\x71\x98\xcd\x40\x8b\x7d\x80\xc9\xca\xc6\x94\x2e\x6c\x52\x44\xbb\x2d\x36\x4a\x12\x1a\xa3\x55\x43\x5e\xbf\x24\x64\xd3\xc2\xc0\xe5\xae\xfe\xf3\x51\x0a\xfa\x40\xe1\x83\x75\x37\xe4\x9f\xb9\xda\x1f\x8d\x29\xd8\xb5\x2f\x37\xc4\x90\xcb\x90\x5c\xde\x23\x42\xf7\x83\x1a\xa7\x0c\x63\x6d\x86\xc1\x25\xad\x4b\x99\xbc\xbb\xa0\x0f\xf3\xea\x17\x94\x80\x32\x3a\x6c\x21\x7d\x23\x47\xd7\xa3\x98\xaf\xd9\x3d\x9d\xd3\x47\x3f\x0d\xc9\x14\x07\x1d\x09\x13\xaa\x91\x50\xec\x59\x34\x47\xfe\x79\xe4\x8c\x73\xd4\x9d\xd0\x6f\x2d\x4e\x1a\x1f\x4c\xd6\xaf\x4c\xa2\xed\x14\x5a\x2d\x04\x78\xf3\xc2\xb4\x50\xb8\x5e\x2b\xf2\x47\xe5\x61\x5b\xee\xb8\x5c\x76\xef\xff\xe0\x8a\xfc\x06\x00\x00\xff\xff\x36\xac\x57\xed\xfc\x00\x00\x00")

func migrations20151216WorkSpecRuntimeSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20151216WorkSpecRuntimeSql,
		"migrations/20151216-work-spec-runtime.sql",
	)
}

func migrations20151216WorkSpecRuntimeSql() (*asset, error) {
	bytes, err := migrations20151216WorkSpecRuntimeSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20151216-work-spec-runtime.sql", size: 252, mode: os.FileMode(420), modTime: time.Unix(1487206567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20160104NotBeforeSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\xcc\xc1\x4a\xc4\x30\x10\xc6\xf1\x7b\x9e\xe2\x63\x8f\x4a\x7c\x80\xed\x29\x9a\x80\x0b\xed\x76\x59\xb3\x08\x5e\x96\x6c\x13\xdb\x62\x9b\xa9\xc9\x94\xbe\xbe\xb4\x08\x22\x22\x0c\x03\x03\xff\xf9\x49\x09\x79\x27\x31\x92\x0f\x7b\xe4\xcf\xa1\x58\x97\x9c\x12\xf9\xb9\xe1\x3d\x26\xca\xdc\xa6\x90\xd7\x48\xc8\x75\x60\xbb\x3e\xc3\x79\x9f\xe1\xb0\x8b\xc4\xd7\x5b\x78\xa7\x14\x76\x68\x68\x98\xc7\x08\x26\x70\x17\xb0\x50\xfa\xc0\x1c\x7b\x06\xbb\xdb\x10\x1e\xbe\xbf\xef\xc7\xbe\x4d\x8e\x03\x2e\x93\x50\xa5\x35\x67\x58\xf5\x58\x9a\x2d\xbf\x6e\xb9\xd2\x1a\x4f\x75\x79\xa9\x8e\xf8\xd1\x61\x0f\x95\x79\xb1\xaa\x3a\xe1\xf5\x60\x9f\xb7\x13\x6f\xf5\xd1\x14\xe2\x97\xaa\x69\x89\xff\xb8\xfa\x5c\x9f\xfe\xc2\x85\xf8\x0a\x00\x00\xff\xff\x57\xda\xe1\x0e\x02\x01\x00\x00")

func migrations20160104NotBeforeSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20160104NotBeforeSql,
		"migrations/20160104-not-before.sql",
	)
}

func migrations20160104NotBeforeSql() (*asset, error) {
	bytes, err := migrations20160104NotBeforeSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20160104-not-before.sql", size: 258, mode: os.FileMode(420), modTime: time.Unix(1487206567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20160125IndexSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\xd0\x3f\x4f\xeb\x30\x14\x05\xf0\x3d\x9f\xe2\x8c\xed\x7b\x2f\x1d\xde\xd8\x4c\x88\x64\xa8\x54\xb5\x50\x8a\xc4\x44\xe5\xc4\x37\xc9\x15\xfe\x87\x7d\x43\xda\x6f\x8f\x0c\x14\x89\x01\xc9\xf2\x70\x64\xfd\x7c\xee\x2d\x4b\x94\x7f\x4a\x58\xaf\x69\x8d\xf4\x6a\xaa\x7c\x95\x21\x7a\x3d\x75\xb2\x46\xf0\x49\x86\x48\x29\x3f\x2a\xca\x7c\x70\x1c\x39\x41\x69\x9d\xa0\x1c\xd8\x69\x3a\x43\x3c\xa8\xef\xb9\x63\x72\x62\x2e\xe8\xd9\x69\x28\x11\xb2\x41\x12\xda\x0b\x64\x24\x8e\x98\x7d\x7c\xc9\xc0\xe4\x58\xd2\x0a\xd8\x38\x04\x15\x85\xbb\xc9\xa8\xf8\x0f\x92\xdd\x9e\xcf\x94\xa0\x10\xa2\x6f\x0d\x59\xcc\x23\x45\x82\x26\x43\xc2\x6e\x80\x51\x71\xa0\x6c\xb8\xc9\xb6\x14\x13\x7c\x9f\x7f\xe3\x34\x92\xfe\xf0\x3f\x71\x70\x42\x32\x7e\xae\x40\xe7\x40\x91\x2d\x39\x51\xc6\x5c\xc0\x02\x43\x4a\x27\x88\xcf\xca\x7e\xe1\x9e\xff\x2f\xd1\xd2\xa8\xde\xd8\x47\x78\x97\xab\x22\xb1\x1b\x0c\xe1\xe1\x7e\x8b\xba\xd9\x36\xc7\x06\x9d\xb7\x56\x39\xbd\xfa\x5a\xc1\x5f\xcb\x43\x54\x42\x78\x0c\xc5\xed\xa1\xb9\x39\x36\xd8\xec\xea\xe6\xe9\x3a\xf4\x29\x57\x39\xe5\x2a\xd8\xef\xae\xe1\xe2\x3b\x3c\xb1\x5e\x56\xc5\x0f\xa8\xf6\xb3\x2b\xea\xc3\xfe\xee\x37\xa8\x2a\xde\x03\x00\x00\xff\xff\x2c\x4b\x17\x27\xaa\x01\x00\x00")

func migrations20160125IndexSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20160125IndexSql,
		"migrations/20160125-index.sql",
	)
}

func migrations20160125IndexSql() (*asset, error) {
	bytes, err := migrations20160125IndexSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20160125-index.sql", size: 426, mode: os.FileMode(420), modTime: time.Unix(1487206567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20160217AttemptSpecSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x90\x5f\x6b\xbb\x30\x14\x86\xef\xf3\x29\xde\xcb\xdf\x9f\xe9\x07\xa8\xf4\xc2\xe9\xe9\x36\xb0\x5a\x6c\x64\x97\xc5\x99\xb4\x0d\xd3\xc4\x99\x14\xd9\xb7\x1f\x91\xad\xae\x60\x43\x08\x9c\x27\x39\x4f\x5e\x4e\x10\x20\xf8\x17\xa0\x33\x42\xae\x60\x3f\xda\xc8\x1f\x41\x3f\x18\x71\x69\xdc\x0a\xbd\xb1\xee\x34\x48\xeb\x1f\xb1\xc0\x6f\xf0\xb3\xb2\xa8\x85\xb0\xa8\x31\x9a\xe1\xfd\x60\x7b\xd9\xa0\x31\xed\xa5\xd3\x70\x06\xee\x2c\x51\x3b\x27\xbb\xde\xc1\xd5\x6f\xad\x7c\x80\xf5\xb4\x76\x18\x25\x9a\x5a\x7b\x89\x3c\x1e\x55\xa3\xa4\x76\xed\x27\x8e\x4a\x8b\xc9\x04\x6f\xb2\x18\x95\x3b\xa3\x97\x5a\x28\x7d\x9a\x78\xc8\x7c\xcb\xff\x4e\x9d\x86\xda\x49\x54\x3d\x8b\x33\x4e\x25\x78\xfc\x98\xd1\xcf\x5f\x2c\x4e\x53\x24\x45\x56\x6d\xf3\x39\xd6\x41\x09\xbc\xe4\x9c\x9e\xa8\x64\xb8\xb3\x4a\xda\x50\x49\x79\x42\xfb\xb9\xef\x8f\x12\x7f\x51\xe4\x48\x29\x23\x4e\x48\xe2\x7d\x12\xa7\x14\xdd\x75\xb0\x6a\x97\xc6\x7c\xce\xb2\x27\x7e\x13\x62\x3d\x15\x17\xad\x5c\xf8\x1b\xb3\x4d\x59\x6c\x71\xbd\x63\xaf\xcf\x54\xd2\x5c\x87\x4a\xac\xbf\x8d\xe1\x15\x1e\x94\x88\xd8\xf2\x00\x26\xb6\x34\x02\x1f\x27\x2f\x38\xf2\x2a\xcb\xa2\xdb\x61\xa6\x66\xd4\x8b\xb6\xb4\x2c\x76\x4b\xb2\x88\x7d\x05\x00\x00\xff\xff\x3b\x93\xb0\x97\x33\x02\x00\x00")

func migrations20160217AttemptSpecSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20160217AttemptSpecSql,
		"migrations/20160217-attempt-spec.sql",
	)
}

func migrations20160217AttemptSpecSql() (*asset, error) {
	bytes, err := migrations20160217AttemptSpecSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20160217-attempt-spec.sql", size: 563, mode: os.FileMode(420), modTime: time.Unix(1487206567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20160328IndexSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x90\xc1\x6e\xdb\x30\x10\x44\xef\xfa\x8a\xb9\xb5\x69\xad\x7c\x40\x74\x0a\x2c\x01\x35\x60\x38\xa9\x9d\xa0\xbd\x19\x8c\xb8\x92\x16\x96\x48\x96\xbb\x8a\xdb\xbf\x2f\x28\x29\x68\xd3\xd4\xd7\x10\x04\x41\x2c\x87\x6f\x67\x36\xcf\x91\x7f\xca\x31\x78\x4b\x37\x90\x1f\x7d\x91\x8e\x3c\x44\x6f\xc7\x5a\x6f\x10\xbc\x68\x1b\x49\x92\x28\xcb\xd3\xc6\x43\xc7\x82\x31\x58\xa3\x24\xd0\x8e\xc0\xce\xd2\x4f\x68\x67\x14\x0d\x3b\x2b\x38\xfb\x78\xc2\xe8\x58\x05\x67\xd6\x6e\x12\x75\xdc\x76\x24\x8a\x10\xd9\x47\xd6\x5f\x89\xa4\x1e\x72\xe2\x00\xff\x4c\xf1\xcd\x27\x53\x2b\x3f\x13\x8c\x2a\x0d\x41\x65\x35\x37\xa8\x8d\xfb\xa0\x78\x22\x58\xef\xe8\x1a\xd8\x68\x02\x99\x5e\x3c\x6c\xf4\x41\x60\x10\xc9\x8e\xce\x1a\xa7\x8b\x31\xdf\xcc\x6c\x09\x54\x0b\x8c\xb3\xff\xa0\xb1\x29\xa5\x00\x37\x09\x74\x3f\xc7\x3d\x7c\xdd\xc2\x11\x59\xc1\x93\x9f\xfc\xb3\x6b\x05\x3c\xb5\x87\x91\xd3\x5c\x16\x76\x6d\x4f\x79\xc3\xd4\xdb\xb9\x17\xc9\xf5\x32\xa4\xcf\x03\xb7\xd1\x28\xe1\x31\x64\xe5\xfe\xee\x1e\x9b\x5d\x59\x7d\x9f\x8c\x1c\x53\xc8\xa3\x8f\x96\x22\xbb\xb6\xc8\xd6\xfb\xea\xf6\xa1\xba\x28\xc0\xdd\xee\x4f\xf5\xe3\xcb\xf8\x50\x56\x87\xf5\x0a\xce\x0c\x84\xdb\xc3\xfa\x2a\xc3\xbc\xbe\x7d\xa9\xf6\xd5\x12\xf0\xb8\x04\x3c\xb2\xc5\xe6\x80\xdd\xe3\x76\x5b\xfc\xdf\x4c\x1a\xcd\x8b\xba\xc8\x5e\xf9\x2f\xfd\xd9\xbd\x4b\x82\x8b\x98\xbf\xcd\xbd\x46\x4d\xb7\xe9\x99\xed\xea\x6d\xe8\xab\x22\xfb\x1d\x00\x00\xff\xff\xe1\x29\x9e\x4d\xde\x02\x00\x00")

func migrations20160328IndexSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20160328IndexSql,
		"migrations/20160328-index.sql",
	)
}

func migrations20160328IndexSql() (*asset, error) {
	bytes, err := migrations20160328IndexSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20160328-index.sql", size: 734, mode: os.FileMode(420), modTime: time.Unix(1487206567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20160329IndexSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x8e\x4d\x6a\xc3\x30\x10\x85\xf7\x3a\xc5\x5b\xf6\x4f\x3d\x40\xbc\x2a\xb5\xa1\x81\xe0\x14\x37\xa1\xdd\x19\xc5\x9a\x44\xa2\xb6\xa4\x4a\x13\xbb\xb9\x7d\x91\x13\x08\x5d\x45\x08\x21\x66\x1e\xdf\xf7\xa4\x84\x7c\x90\x18\xbc\xa6\x05\xd2\x4f\x5f\xe4\x47\x86\xe8\xf5\xb1\xe3\x05\x82\x4f\x7c\x88\x94\x72\x48\xc8\x7c\xb1\x31\x36\x41\x69\x9d\xa0\x9c\x67\x43\x11\x91\x12\x47\xdb\x31\x69\x58\xa7\xe9\x17\xec\x61\xa8\x0f\xd8\x5b\xa7\x31\xf9\xf8\x8d\x14\xa8\x4b\x98\x2c\x9b\x8c\x50\xa3\xb2\xbd\xda\xf5\x34\x2f\x9f\x30\x19\x72\x18\x94\x3b\xc1\xef\xc1\xe6\x3c\xc6\xd1\x59\x4e\x30\x6a\x24\xec\x88\x1c\x14\x33\x0d\x81\x49\xcf\x88\x3e\x92\xd2\xa7\xe7\x4b\xa9\xc7\xc1\x1e\xa2\x62\xc2\x36\x88\xd7\xa6\x7a\xd9\x54\x58\xd6\x65\xf5\x35\x93\xda\x4c\x6a\x73\x85\xf6\x6a\x5e\xd7\xd7\xdd\xdd\xfc\x9b\x03\x56\xdf\x0b\x9c\xcf\xe7\x5b\xd5\x54\x50\x1d\xdb\x91\xda\x8b\xbc\xb5\x1a\xcb\x0f\xd4\xdb\xd5\xaa\x10\xff\xc4\xa5\x9f\x9c\x28\x9b\xf5\xfb\x0d\x71\x21\xfe\x02\x00\x00\xff\xff\xab\x49\xaf\xc7\x73\x01\x00\x00")

func migrations20160329IndexSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20160329IndexSql,
		"migrations/20160329-index.sql",
	)
}

func migrations20160329IndexSql() (*asset, error) {
	bytes, err := migrations20160329IndexSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20160329-index.sql", size: 371, mode: os.FileMode(420), modTime: time.Unix(1487206567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20170316IndexSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\xcc\xbf\x6a\x87\x30\x14\xc5\xf1\x3d\x4f\x71\xc6\xfe\x21\x7d\x00\x9d\x44\x33\x14\x8a\x96\xa0\xd0\x4e\x22\x26\x48\xd0\xe4\xa6\x49\x44\xe9\xd3\x17\x31\x4b\xf9\xc1\xe5\x0e\x5f\x0e\x1f\xce\xc1\x5f\x38\x2c\x29\x5d\x20\xfe\x6c\xe5\xf5\xb8\x0f\xa4\xf6\x39\x15\xf0\x14\xd3\x12\x74\xbc\x46\x8c\x5f\x87\x4a\xa9\x08\xe3\x94\x3e\x91\x08\x07\x85\x55\x87\x37\x3f\x05\xed\x12\x22\x61\xc2\x4c\xd6\x92\xc3\x46\xb4\xee\x1e\x26\x82\x7c\x32\xd6\xfc\x6a\x95\x85\x57\x6b\x96\x30\x25\x8d\xc1\x33\x56\x4b\x51\xf5\x02\xef\x6d\x23\xbe\x50\x77\x6d\x3d\x48\x29\xda\xfe\xe3\x3b\xd3\xe3\x4d\x8f\x46\x9d\x20\x97\xe3\xd3\x1d\x9f\x4b\xf6\xcf\x6b\xe8\x70\xac\x91\xdd\x67\xe6\x1e\x84\x92\xfd\x05\x00\x00\xff\xff\x4e\xd9\x97\xa2\xf0\x00\x00\x00")

func migrations20170316IndexSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20170316IndexSql,
		"migrations/20170316-index.sql",
	)
}

func migrations20170316IndexSql() (*asset, error) {
	bytes, err := migrations20170316IndexSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20170316-index.sql", size: 240, mode: os.FileMode(420), modTime: time.Unix(1489676316, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"migrations/20150927-core.sql": migrations20150927CoreSql,
	"migrations/20151002-mingb.sql": migrations20151002MingbSql,
	"migrations/20151006-work-unit-type.sql": migrations20151006WorkUnitTypeSql,
	"migrations/20151013-index.sql": migrations20151013IndexSql,
	"migrations/20151014-index.sql": migrations20151014IndexSql,
	"migrations/20151019-worker-mode.sql": migrations20151019WorkerModeSql,
	"migrations/20151028-index.sql": migrations20151028IndexSql,
	"migrations/20151216-work-spec-runtime.sql": migrations20151216WorkSpecRuntimeSql,
	"migrations/20160104-not-before.sql": migrations20160104NotBeforeSql,
	"migrations/20160125-index.sql": migrations20160125IndexSql,
	"migrations/20160217-attempt-spec.sql": migrations20160217AttemptSpecSql,
	"migrations/20160328-index.sql": migrations20160328IndexSql,
	"migrations/20160329-index.sql": migrations20160329IndexSql,
	"migrations/20170316-index.sql": migrations20170316IndexSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"migrations": &bintree{nil, map[string]*bintree{
		"20150927-core.sql": &bintree{migrations20150927CoreSql, map[string]*bintree{}},
		"20151002-mingb.sql": &bintree{migrations20151002MingbSql, map[string]*bintree{}},
		"20151006-work-unit-type.sql": &bintree{migrations20151006WorkUnitTypeSql, map[string]*bintree{}},
		"20151013-index.sql": &bintree{migrations20151013IndexSql, map[string]*bintree{}},
		"20151014-index.sql": &bintree{migrations20151014IndexSql, map[string]*bintree{}},
		"20151019-worker-mode.sql": &bintree{migrations20151019WorkerModeSql, map[string]*bintree{}},
		"20151028-index.sql": &bintree{migrations20151028IndexSql, map[string]*bintree{}},
		"20151216-work-spec-runtime.sql": &bintree{migrations20151216WorkSpecRuntimeSql, map[string]*bintree{}},
		"20160104-not-before.sql": &bintree{migrations20160104NotBeforeSql, map[string]*bintree{}},
		"20160125-index.sql": &bintree{migrations20160125IndexSql, map[string]*bintree{}},
		"20160217-attempt-spec.sql": &bintree{migrations20160217AttemptSpecSql, map[string]*bintree{}},
		"20160328-index.sql": &bintree{migrations20160328IndexSql, map[string]*bintree{}},
		"20160329-index.sql": &bintree{migrations20160329IndexSql, map[string]*bintree{}},
		"20170316-index.sql": &bintree{migrations20170316IndexSql, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

