package jsonline

import (
	"encoding/json"
	"fmt"
)

type format int8

const (
	String  format = iota // default representation, e.g. : "hello", "2.4", "true"
	Numeric               // integer or decimal, e.g. : 2.4, 1
	Boolean               // true or false
	Binary                // base64 encoded data
	Hidden                // will never appear in the output
)

type Value interface {
	SetExportFormat(format)

	Export() interface{}

	json.Marshaler
	fmt.Stringer
}

type value struct {
	raw    interface{}
	export format
}

func NewValue(v interface{}) Value {
	return &value{
		raw:    v,
		export: String,
	}
}

func (v *value) SetExportFormat(f format) {
	v.export = f
}

func (v *value) Export() interface{} {
	val := v.raw

	switch v.export {
	case String:
		return toString(val)
	case Numeric:
		return toNumeric(val)
	case Boolean:
		return toBoolean(val)
	case Binary:
	case Hidden:
		panic(fmt.Errorf("%w: %v", ErrExportHiddenValueFobidden, v))
	}

	return v.raw
}

func (v *value) MarshalJSON() (res []byte, err error) {
	e := v.Export()

	b, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("can't marshal value %v to json: %w", e, err)
	}

	return b, nil
}

func (v *value) String() string {
	b, err := v.MarshalJSON()
	if err != nil {
		return fmt.Sprintf("ERROR: %v", err)
	}

	return string(b)
}
