// Copyright (C) 2021 CGI France
//
// This file is part of the jsonline library.
//
// The jsonline library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The jsonline library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the jsonline library.  If not, see <http://www.gnu.org/licenses/>.
//
// Linking this library statically or dynamically with other modules is
// making a combined work based on this library.  Thus, the terms and
// conditions of the GNU General Public License cover the whole
// combination.
//
// As a special exception, the copyright holders of this library give you
// permission to link this library with independent modules to produce an
// executable, regardless of the license terms of these independent
// modules, and to copy and distribute the resulting executable under
// terms of your choice, provided that you also meet, for each linked
// independent module, the terms and conditions of the license of that
// module.  An independent module is a module which is not derived from
// or based on this library.  If you modify this library, you may extend
// this exception to your version of the library, but you are not
// obligated to do so.  If you do not wish to do so, delete this
// exception statement from your version.

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
	Auto                    // no specific format enforced
	Hidden                  // will not be exported in jsonline
)

type Value interface {
	GetFormat() format

	Raw() interface{}
	Export() (interface{}, error)
	Import(interface{}) error

	json.Unmarshaler
	json.Marshaler
	fmt.Stringer
}

type value struct {
	raw interface{}
	f   format
}

func NewValue(v interface{}, f format) Value {
	return &value{
		raw: v,
		f:   f,
	}
}

func NewValueNil() Value {
	return &value{
		raw: nil,
		f:   Auto,
	}
}

func NewValueAuto(v interface{}) Value {
	return &value{
		raw: v,
		f:   Auto,
	}
}

func NewValueString(v interface{}) Value {
	return &value{
		raw: v,
		f:   String,
	}
}

func NewValueNumeric(v interface{}) Value {
	return &value{
		raw: v,
		f:   Numeric,
	}
}

func NewValueBoolean(v interface{}) Value {
	return &value{
		raw: v,
		f:   Boolean,
	}
}

func NewValueBinary(v interface{}) Value {
	return &value{
		raw: v,
		f:   Binary,
	}
}

func NewValueDateTime(v interface{}) Value {
	return &value{
		raw: v,
		f:   DateTime,
	}
}

func NewValueTime(v interface{}) Value {
	return &value{
		raw: v,
		f:   Time,
	}
}

func NewValueTimestamp(v interface{}) Value {
	return &value{
		raw: v,
		f:   Timestamp,
	}
}

func NewValueHidden(v interface{}) Value {
	return &value{
		raw: v,
		f:   Hidden,
	}
}

func CloneValue(v Value) Value {
	return NewValue(v.Raw(), v.GetFormat())
}

func (v *value) GetFormat() format {
	return v.f
}

func (v *value) Raw() interface{} {
	return v.raw
}

func (v *value) Export() (interface{}, error) {
	val := v.raw

	switch v.f {
	case String:
		return exportToString(val), nil
	case Numeric:
		return exportToNumeric(val)
	case Boolean:
		return exportToBoolean(val)
	case Binary:
		return exportToBinary(val), nil
	case DateTime:
		return exportToDateTime(val)
	case Time:
		return exportToTime(val)
	case Timestamp:
		return exportToTimeStamp(val)
	case Auto, Hidden:
		return v.raw, nil
	}

	return nil, fmt.Errorf("%w: %#v", ErrUnsupportedFormat, v.f)
}

func (v *value) Import(val interface{}) error {
	var err error

	switch v.f {
	case String:
		v.raw = importToString(val)
	case Numeric:
		v.raw = val
	case Boolean:
		v.raw = val
	case Binary:
		v.raw = val
	case DateTime:
		v.raw, err = importToDateTime(val)
	case Time:
		v.raw = val
	case Timestamp:
		v.raw = val
	case Auto, Hidden:
		v.raw = val
	default:
		return fmt.Errorf("%w: %#v", ErrUnsupportedFormat, v.f)
	}

	return err
}

func (v *value) MarshalJSON() ([]byte, error) {
	e, err := v.Export()
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("can't marshal value %v to json: %w", e, err)
	}

	return b, nil
}

func (v *value) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &v.raw) //nolint
}

func (v *value) String() string {
	b, err := v.MarshalJSON()
	if err != nil {
		return fmt.Sprintf("ERROR: %v", err)
	}

	return string(b)
}
