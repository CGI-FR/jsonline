package jsonline

import (
	"encoding/json"
	"fmt"
)

type format int8

const (
	String    format = iota // default representation, e.g. : "hello", "2.4", "true"
	Numeric                 // integer or decimal, e.g. : 2.4, 1
	Boolean                 // true or false
	Binary                  // base64 encoded data
	DateTime                // date and time as RFC3339, e.g. : "2006-01-02T15:04:05Z", "2006-01-02T15:04:05+07:00"
	Time                    // time with timezone, e.g. : "15:04:05Z", "15:04:05-07:00"
	Timestamp               // milliseconds since 1970
	None                    // no specific format enforced
)

type Value interface {
	Export() interface{}
	Import(interface{}) Value

	json.Marshaler
	fmt.Stringer
}

type value struct {
	raw interface{}
	f   format
}

func NewValue(f format) Value {
	return &value{
		raw: nil,
		f:   f,
	}
}

func NewRaw(v interface{}) Value {
	return &value{
		raw: v,
		f:   None,
	}
}

func NewString(v interface{}) Value {
	return &value{
		raw: v,
		f:   String,
	}
}

func NewNumeric(v interface{}) Value {
	return &value{
		raw: v,
		f:   Numeric,
	}
}

func NewBoolean(v interface{}) Value {
	return &value{
		raw: v,
		f:   Boolean,
	}
}

func NewBinary(v interface{}) Value {
	return &value{
		raw: v,
		f:   Binary,
	}
}

func NewDateTime(v interface{}) Value {
	return &value{
		raw: v,
		f:   DateTime,
	}
}

func NewTime(v interface{}) Value {
	return &value{
		raw: v,
		f:   Time,
	}
}

func NewTimestamp(v interface{}) Value {
	return &value{
		raw: v,
		f:   Timestamp,
	}
}

func (v *value) Export() interface{} {
	val := v.raw

	switch v.f {
	case String:
		return toString(val)
	case Numeric:
		return toNumeric(val)
	case Boolean:
		return toBoolean(val)
	case Binary:
		return toBinary(val)
	case DateTime:
		return toDateTime(val)
	case Time:
		return toTime(val)
	case Timestamp:
		return toTimeStamp(val)
	case None:
		return v.raw
	}

	panic(fmt.Errorf("%w: %#v", ErrUnsupportedFormat, v.f))
}

func (v *value) Import(val interface{}) Value {
	v.raw = val

	return v
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
