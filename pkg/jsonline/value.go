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

	"github.com/cgi-fr/jsonline/pkg/cast"
)

type (
	Format  int8
	RawType interface{}
)

const (
	String    Format = iota // String representation, e.g. : "hello", "2.4", "true".
	Numeric                 // Numeric (integer or decimal), e.g. : 2.4, 1.
	Boolean                 // Boolean : true or false.
	Binary                  // Binary representation encoded as base64.
	DateTime                // DateTime as RFC3339, e.g. : "2006-01-02T15:04:05Z", "2006-01-02T15:04:05+07:00".
	Timestamp               // Timestamp the number of seconds since 1970 ("UNIX time").
	Auto                    // Auto columns have no specific format enforced.
	Hidden                  // Hidden columns will not be exported in jsonline.
)

type Value interface {
	GetFormat() Format
	GetRawType() RawType

	Raw() interface{}

	Export() (interface{}, error)
	Import(interface{}) error

	json.Unmarshaler
	json.Marshaler
	fmt.Stringer

	DebugString() string
}

type value struct {
	raw interface{}
	f   Format
	typ RawType
}

func NewValue(v interface{}, f Format, rawtype RawType) Value {
	return &value{
		raw: v,
		f:   f,
		typ: rawtype,
	}
}

func NewValueNil(f Format, rawtype RawType) Value {
	return &value{
		raw: nil,
		f:   f,
		typ: rawtype,
	}
}

func NewValueAuto(v interface{}) Value {
	return &value{
		raw: v,
		f:   Auto,
		typ: nil,
	}
}

func NewValueString(v interface{}) Value {
	return &value{
		raw: v,
		f:   String,
		typ: nil,
	}
}

func NewValueNumeric(v interface{}) Value {
	return &value{
		raw: v,
		f:   Numeric,
		typ: nil,
	}
}

func NewValueBoolean(v interface{}) Value {
	return &value{
		raw: v,
		f:   Boolean,
		typ: nil,
	}
}

func NewValueBinary(v interface{}) Value {
	return &value{
		raw: v,
		f:   Binary,
		typ: nil,
	}
}

func NewValueDateTime(v interface{}) Value {
	return &value{
		raw: v,
		f:   DateTime,
		typ: nil,
	}
}

func NewValueTimestamp(v interface{}) Value {
	return &value{
		raw: v,
		f:   Timestamp,
		typ: nil,
	}
}

func NewValueHidden(v interface{}) Value {
	return &value{
		raw: v,
		f:   Hidden,
		typ: nil,
	}
}

func CloneValue(v Value) Value {
	return NewValue(v.Raw(), v.GetFormat(), v.GetRawType())
}

func (v *value) GetFormat() Format {
	return v.f
}

func (v *value) GetRawType() RawType {
	return v.typ
}

func (v *value) Raw() interface{} {
	return v.raw
}

func (v *value) Export() (interface{}, error) {
	if v.raw == nil {
		return nil, nil
	}

	switch v.f {
	case String:
		return exportToString(v.raw)
	case Numeric:
		return exportToNumber(v.raw)
	case Boolean:
		return exportToBool(v.raw)
	case Binary:
		return exportToBinary(v.raw)
	case DateTime:
		return exportToDateTime(v.raw)
	case Timestamp:
		return exportToTimestamp(v.raw)
	case Auto, Hidden:
		return v.raw, nil
	}

	return nil, fmt.Errorf("%w: %#v", ErrUnsupportedFormat, v.f)
}

//nolint:cyclop
func (v *value) Import(val interface{}) error {
	if val == nil {
		v.raw = nil

		return nil
	}

	if value, ok := val.(Value); ok {
		v.f = value.GetFormat()
		v.raw = value.Raw()
		v.typ = value.GetRawType()

		return nil
	}

	var err error

	switch v.f {
	case String:
		v.raw, err = importFromString(val, v.typ)
	case Numeric:
		v.raw, err = importFromNumeric(val, v.typ)
	case Boolean:
		v.raw, err = importFromBoolean(val, v.typ)
	case Binary:
		v.raw, err = importFromBinary(val, v.typ)
	case DateTime:
		v.raw, err = importFromDateTime(val, v.typ)
	case Timestamp:
		v.raw, err = importFromTimestamp(val, v.typ)
	case Auto, Hidden:
		v.raw, err = cast.To(v.typ, val)
	default:
		err = fmt.Errorf("%w: %#v", ErrUnsupportedFormat, v.f)
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
	var raw interface{}

	err := json.Unmarshal(data, &raw)
	if err != nil {
		return fmt.Errorf("can't unmarshal JSON data: %w", err)
	}

	err = v.Import(raw)
	if err != nil {
		return fmt.Errorf("can't unmarshal JSON data: %w", err)
	}

	return nil
}

func (v *value) String() string {
	b, err := v.MarshalJSON()
	if err != nil {
		return fmt.Sprintf("ERROR: %v", err)
	}

	return string(b)
}

func (v *value) DebugString() string {
	return fmt.Sprintf("%T(%#v)", v.raw, v.raw)
}
