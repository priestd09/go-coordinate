// Copyright 2015 Diffeo, Inc.
// This software is released under an MIT/X11 open source license.

// Package cborrpc defines the CBOR-RPC format used by the Python
// Coordinate daemon.
package cborrpc

import (
	"github.com/satori/go.uuid"
	"github.com/ugorji/go/codec"
	"reflect"
)

// Request defines the fields of a CBOR-RPC request.
type Request struct {
	// Name of the RPC method to invoke.
	Method string
	// Sequential, non-unique identifier for this request.
	ID uint
	// List of arbitrary parameters.
	Params []interface{}
}

// Response defines the fields of a CBOR-RPC response
type Response struct {
	// Sequential, non-unique identifier for this response.  This should
	// always match the identifier from the corresponding Request.
	ID uint
	// Arbitrary response object; should be nil on error.
	Result interface{}
	// Error message on failure; should be empty on success.
	Error string
}

// Wrapper type to avoid codec "fast path" encoding
type keyOrValue interface{}

// Actual "wire format" representation for top-level CBOR-RPC messages.
type wireFormat []keyOrValue

// MapBySlice is a marker for the codec library to indicate this is
// actually a map.
func (w wireFormat) MapBySlice() {}

// (NB: the read end of this is pretty close to mapstructure's intended
// use case, including the sloppy bytes/chars in the wire format because
// Python 2.  We really want to use the encoding/json marshaler to get
// a dict from an object and then feed that into CBOR, but this seems
// inseparable from the JSON wire format.  For that matter, it looks like
// the codec package this is already using supports reflection-based
// marshaling, so we should just use that.)

// Codec extension plugin to convert Request.
type reqExt struct {
	cbor *codec.CborHandle
}

// Export a Request in some format.
func (x reqExt) ConvertExt(v interface{}) interface{} {
	request := v.(Request)

	// Assemble the "over-the-wire" response dictionary
	wire := wireFormat{
		[]byte("method"),
		[]byte(request.Method),
		[]byte("id"),
		uint64(request.ID),
		[]byte("params"),
		request.Params,
	}

	// Convert it to CBOR bytes
	var resp []byte
	encoder := codec.NewEncoderBytes(&resp, x.cbor)
	encoder.MustEncode(wire)
	return resp
}

// Unpackage some format into a Request.
func (x reqExt) UpdateExt(dest interface{}, v interface{}) {
	data := v.([]byte)
	decoder := codec.NewDecoderBytes(data, x.cbor)
	var wire map[string]interface{}
	decoder.MustDecode(&wire)

	result := dest.(*Request)
	result.Method = string(wire["method"].([]byte))
	result.ID = uint(wire["id"].(uint64))
	result.Params = wire["params"].([]interface{})
}

// Codec extension plugin to convert Response.
type respExt struct {
	cbor *codec.CborHandle
}

// Export a Response in some format.
func (x respExt) ConvertExt(v interface{}) interface{} {
	response := v.(Response)

	// Assemble the "over-the-wire" response dictionary
	wire := wireFormat{
		[]byte("id"),
		uint64(response.ID),
	}
	if response.Result != nil {
		wire = append(wire, []byte("result"), response.Result)
	}
	if response.Error != "" {
		errorDict := make(map[string]string)
		errorDict["message"] = response.Error
		wire = append(wire, []byte("error"), errorDict)
	}

	// Convert it to CBOR bytes
	var resp []byte
	encoder := codec.NewEncoderBytes(&resp, x.cbor)
	encoder.MustEncode(wire)
	return resp
}

// Unpackage some format into a Response.
func (x respExt) UpdateExt(dest interface{}, v interface{}) {
	data := v.([]byte)
	decoder := codec.NewDecoderBytes(data, x.cbor)
	var wire map[string]interface{}
	decoder.MustDecode(&wire)
	response := dest.(*Response)
	response.ID = uint(wire["id"].(uint64))
	response.Result = wire["result"]
	if wire["error"] != nil {
		errorDict := wire["error"].(map[string]interface{})
		response.Error = string(errorDict["message"].([]byte))
	}
}

// PythonTuple is a simple Go wrapper representing a Python tuple.
// This is a fixed-length, ordered, immutable sequence of arbitrary
// items.
type PythonTuple struct {
	Items []interface{}
}

// pythonTupleExt is a codec extension plugin to encode and decode
// PythonTuple objects.
type pythonTupleExt struct {
	cbor *codec.CborHandle
}

func (x pythonTupleExt) ConvertExt(v interface{}) interface{} {
	var ptuple *PythonTuple
	ptuple, ok := v.(*PythonTuple)
	if !ok {
		tuple := v.(PythonTuple)
		ptuple = &tuple
	}
	if ptuple.Items != nil {
		return ptuple.Items
	}
	return []interface{}{}
}

func (x pythonTupleExt) UpdateExt(dest interface{}, v interface{}) {
	items := v.([]interface{})
	tuple := dest.(*PythonTuple)
	*tuple = PythonTuple{items}
}

// uuidExt is a codec extension plugin to encode and decode UUID
// objects.
type uuidExt struct {
	cbor *codec.CborHandle
}

func (x uuidExt) ConvertExt(v interface{}) interface{} {
	uuid := v.(uuid.UUID)
	return uuid.Bytes()
}

func (x uuidExt) UpdateExt(dest interface{}, v interface{}) {
	bytes := v.([]byte)
	if len(bytes) != 16 {
		panic("encoded UUID must have 16 bytes")
	}
	uuidp := dest.(*uuid.UUID)
	*uuidp = uuid.UUID{}
	copy(uuidp[:], bytes)
}

// SetExts sets up the CBOR codec to understand the other objects in
// this package.
func SetExts(cbor *codec.CborHandle) error {
	reqExt := new(reqExt)
	reqExt.cbor = cbor
	var req Request
	if err := cbor.SetInterfaceExt(reflect.TypeOf(req), 24, reqExt); err != nil {
		return err
	}

	respExt := new(respExt)
	respExt.cbor = cbor
	var resp Response
	if err := cbor.SetInterfaceExt(reflect.TypeOf(resp), 24, respExt); err != nil {
		return err
	}

	if err := cbor.SetInterfaceExt(reflect.TypeOf(uuid.UUID{}), 37, &uuidExt{cbor}); err != nil {
		return err
	}

	tupleExt := pythonTupleExt{cbor}
	var tuple PythonTuple
	if err := cbor.SetInterfaceExt(reflect.TypeOf(tuple), 128, &tupleExt); err != nil {
		return err
	}
	return nil
}
