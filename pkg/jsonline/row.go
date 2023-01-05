// Copyright (C) 2022 CGI France
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
	"bytes"
	"container/list"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
	"time"
	"unicode"

	"github.com/cgi-fr/jsonline/pkg/cast"
)

type Row interface {
	Value

	Has(key string) bool
	Get(key string) (interface{}, bool)
	GetOrNil(key string) interface{}
	GetAtIndex(index int) (interface{}, bool)
	GetAtIndexOrNil(index int) interface{}
	GetAtPath(path string) (interface{}, bool)
	GetAtPathOrNil(path string) interface{}
	Set(key string, val interface{})
	SetAtIndex(index int, val interface{})
	Len() int
	Iter() func() (string, interface{}, bool)

	GetString(key string) string
	GetInt(key string) int
	GetInt64(key string) int64
	GetInt32(key string) int32
	GetInt16(key string) int16
	GetInt8(key string) int8
	GetUint(key string) uint
	GetUint64(key string) uint64
	GetUint32(key string) uint32
	GetUint16(key string) uint16
	GetUint8(key string) uint8
	GetFloat64(key string) float64
	GetFloat32(key string) float32
	GetBool(key string) bool
	GetBytes(key string) []byte
	GetTime(key string) time.Time

	ImportAtKey(key string, val interface{}) error
	ImportAtIndex(index int, val interface{}) error
	ImportAtPath(path string, val interface{}) error

	GetValue(key string) (Value, bool)
	GetValueAtIndex(index int) (Value, bool)
	GetValueAtPath(path string) (Value, bool)
	FindValuesAtPath(path string) ([]Value, bool)
	SetValue(key string, val Value) Row
	SetValueAtIndex(index int, val Value) Row
	IterValues() func() (string, Value, bool)

	MapTo(interface{})
}

type m map[string]Value

type row struct {
	m
	l    *list.List
	keys map[string]*list.Element
}

// NewRow create a new Row.
func NewRow() Row {
	return &row{
		m:    make(map[string]Value),
		l:    list.New(),
		keys: make(map[string]*list.Element),
	}
}

func CloneRow(r Row) Row {
	result := NewRow()

	iter := r.IterValues()

	for k, v, ok := iter(); ok; k, v, ok = iter() {
		result.SetValue(k, CloneValue(v))
	}

	return result
}

func (r *row) GetFormat() Format {
	return Auto
}

func (r *row) GetRawType() RawType {
	return nil
}

func (r *row) Raw() interface{} {
	result := map[string]interface{}{}

	iter := r.Iter()

	for k, v, ok := iter(); ok; k, v, ok = iter() {
		result[k] = v
	}

	return result
}

func (r *row) Export() (interface{}, error) {
	result := map[string]interface{}{}

	iter := r.IterValues()

	for k, v, ok := iter(); ok; k, v, ok = iter() {
		valExported, err := v.Export()
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		result[k] = valExported
	}

	return result, nil
}

func (r *row) Import(v interface{}) error {
	switch values := v.(type) {
	case []interface{}:
		for i, val := range values {
			if err := r.ImportAtIndex(i, val); err != nil {
				return err
			}
		}
	case map[string]interface{}:
		for key, val := range values {
			if err := r.ImportAtKey(key, val); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("%w", ErrUnsupportedImportType)
	}

	return nil
}

func (r *row) ImportAtKey(key string, val interface{}) error {
	if _, ok := r.m[key]; !ok {
		r.keys[key] = r.l.PushBack(key)
	}

	if value, exist := r.m[key]; exist {
		if err := value.Import(val); err != nil {
			return fmt.Errorf("%w", err)
		}

		r.m[key] = value
	} else if value, ok := val.(Value); ok {
		r.m[key] = value
	} else {
		r.m[key] = NewValueAuto(val)
	}

	return nil
}

func (r *row) ImportAtIndex(index int, val interface{}) error {
	var key string

	for cur := r.l.Front(); cur != nil; cur = cur.Next() {
		if index == 0 {
			key, _ = cur.Value.(string)

			break
		}
		index--
	}

	return r.ImportAtKey(key, val)
}

func (r *row) ImportAtPath(path string, val interface{}) error {
	if value, exist := r.GetValueAtPath(path); exist {
		if err := value.Import(val); err != nil {
			return fmt.Errorf("%w", err)
		}
	} else {
		return fmt.Errorf("%w", ErrPathNotFound)
	}

	return nil
}

func (r *row) Has(key string) bool {
	_, ok := r.m[key]

	return ok
}

func (r *row) Get(key string) (interface{}, bool) {
	val, ok := r.m[key]
	if ok {
		return val.Raw(), ok
	}

	return nil, ok
}

func (r *row) GetOrNil(key string) interface{} {
	val, ok := r.Get(key)
	if !ok {
		return nil
	}

	return val
}

func (r *row) GetAtIndex(index int) (interface{}, bool) {
	var key string

	for cur := r.l.Front(); cur != nil; cur = cur.Next() {
		if index == 0 {
			key, _ = cur.Value.(string)

			break
		}
		index--
	}

	return r.Get(key)
}

func (r *row) GetAtIndexOrNil(index int) interface{} {
	var key string

	for cur := r.l.Front(); cur != nil; cur = cur.Next() {
		if index == 0 {
			key, _ = cur.Value.(string)

			break
		}
		index--
	}

	return r.GetOrNil(key)
}

func (r *row) GetAtPath(path string) (interface{}, bool) {
	if value, ok := r.GetValueAtPath(path); ok {
		return value.Raw(), true
	}

	return nil, false
}

func (r *row) GetAtPathOrNil(path string) interface{} {
	if result, ok := r.GetAtPath(path); ok {
		return result
	}

	return nil
}

func (r *row) Set(key string, val interface{}) {
	if _, ok := r.m[key]; !ok {
		r.keys[key] = r.l.PushBack(key)
	}

	if value, exist := r.m[key]; exist {
		if raw, err := cast.To(value.GetRawType(), val); err != nil {
			r.m[key] = NewValue(raw, value.GetFormat(), value.GetRawType())
		} else {
			r.m[key] = NewValue(val, value.GetFormat(), value.GetRawType())
		}
	} else if value, ok := val.(Value); ok {
		r.m[key] = value
	} else {
		r.m[key] = NewValueAuto(val)
	}
}

func (r *row) SetAtIndex(index int, val interface{}) {
	var key string

	for cur := r.l.Front(); cur != nil; cur = cur.Next() {
		if index == 0 {
			key, _ = cur.Value.(string)

			break
		}
		index--
	}

	r.Set(key, val)
}

func (r *row) Len() int {
	return r.l.Len()
}

func (r *row) Iter() func() (string, interface{}, bool) {
	iter := r.IterValues()

	return func() (string, interface{}, bool) {
		key, val, ok := iter()
		if ok {
			return key, val.Raw(), ok
		}

		return key, val, ok
	}
}

func (r *row) GetValue(key string) (Value, bool) {
	val, ok := r.m[key]
	if ok {
		return val, ok
	}

	return nil, ok
}

func (r *row) GetValueAtIndex(index int) (Value, bool) {
	var key string

	for cur := r.l.Front(); cur != nil; cur = cur.Next() {
		if index == 0 {
			key, _ = cur.Value.(string)

			break
		}
		index--
	}

	return r.GetValue(key)
}

func (r *row) GetValueAtPath(path string) (Value, bool) {
	keys := strings.Split(path, ".")

	var row Row = r
	for _, key := range keys {
		value, exist := row.GetValue(key)
		if !exist {
			return nil, false
		}

		if cast, ok := value.(Row); ok {
			row = cast
		} else {
			return value, true
		}
	}

	return row, true
}

func (r *row) FindValuesAtPath(path string) ([]Value, bool) {
	keys := strings.SplitN(path, ".", 2)

	value, exist := r.GetValue(keys[0])
	if !exist {
		return nil, false
	}

	if cast, ok := value.(Row); ok {
		return cast.FindValuesAtPath(keys[1])
	}

	result := []Value{}

	switch typedValue := value.Raw().(type) {
	case []interface{}:
		for _, row := range typedValue {
			if cast, ok := row.(Row); ok {
				if values, exists := cast.FindValuesAtPath(keys[1]); exists {
					result = append(result, values...)
				}
			}
		}
	default:
		return []Value{value}, true
	}

	return result, true
}

func (r *row) SetValue(key string, val Value) Row {
	if _, ok := r.m[key]; !ok {
		r.keys[key] = r.l.PushBack(key)
	}

	r.m[key] = val

	return r
}

func (r *row) SetValueAtIndex(index int, val Value) Row {
	var key string

	for cur := r.l.Front(); cur != nil; cur = cur.Next() {
		if index == 0 {
			key, _ = cur.Value.(string)

			break
		}
		index--
	}

	return r.SetValue(key, val)
}

func (r *row) SetValueAtPath(path string, val Value) Row {
	panic("not implemented")
}

func (r *row) IterValues() func() (string, Value, bool) {
	e := r.l.Front()

	return func() (string, Value, bool) {
		if e != nil {
			key, _ := e.Value.(string)
			e = e.Next()

			return key, r.m[key], true
		}

		return "", nil, false
	}
}

func (r *row) MapTo(v interface{}) {
	t := reflect.TypeOf(v).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		value, exist := r.Get(LcFirst(field.Name))
		if exist {
			switch val := value.(type) {
			case int, int64, int32, int16, int8:
				i, _ := cast.ToInt64(val)
				reflect.ValueOf(v).Elem().FieldByName(field.Name).SetInt(i.(int64))
			case uint, uint64, uint32, uint16, uint8:
				i, _ := cast.ToUint64(val)
				reflect.ValueOf(v).Elem().FieldByName(field.Name).SetUint(i.(uint64))
			case float32, float64:
				i, _ := cast.ToFloat64(val)
				reflect.ValueOf(v).Elem().FieldByName(field.Name).SetFloat(i.(float64))
			case string:
				reflect.ValueOf(v).Elem().FieldByName(field.Name).SetString(val)
			case bool:
				reflect.ValueOf(v).Elem().FieldByName(field.Name).SetBool(val)
			case []byte:
				reflect.ValueOf(v).Elem().FieldByName(field.Name).SetBytes(val)
			}
		}
	}
}

func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}

	return ""
}

func (r *row) MarshalJSON() (res []byte, err error) {
	res = append(res, '{')

	for e := r.l.Front(); e != nil; e = e.Next() {
		k, _ := e.Value.(string)
		if r.m[k].GetFormat() != Hidden {
			res = append(res, fmt.Sprintf("%q:", k)...)

			var b []byte

			b, err = json.Marshal(r.m[k])
			if err != nil {
				return
			}

			res = append(res, b...)
			res = append(res, ',')
		}
	}

	if len(res) > 1 {
		res[len(res)-1] = '}'
	} else {
		res = append(res, '}')
	}

	return
}

func (r *row) String() string {
	b, err := r.MarshalJSON()
	if err != nil {
		return fmt.Sprintf("ERROR: %v", err)
	}

	return string(b)
}

func (r *row) DebugString() string {
	sb := strings.Builder{}
	iter := r.IterValues()

	sep := ""
	for k, v, ok := iter(); ok; k, v, ok = iter() {
		sb.WriteString(fmt.Sprintf(`%v%v={%v}`, sep, k, v.DebugString()))
		sep = ";"
	}

	return strings.ReplaceAll(sb.String(), `"`, "`")
}

//nolint
func (r *row) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()

	// must open with a delim token '{'
	t, err := dec.Token()
	if err != nil {
		return err
	}
	if delim, ok := t.(json.Delim); !ok || delim != '{' {
		return fmt.Errorf("expect JSON object open with '{'")
	}

	err = r.parseobject(dec)
	if err != nil {
		return err
	}

	t, err = dec.Token()
	if err != io.EOF {
		return fmt.Errorf("expect end of JSON object but got more token: %T: %v or err: %v", t, t, err)
	}

	return nil
}

//nolint
func (r *row) parseobject(dec *json.Decoder) (err error) {
	var t json.Token
	for dec.More() {
		t, err = dec.Token()
		if err != nil {
			return err
		}

		key, ok := t.(string)
		if !ok {
			return fmt.Errorf("expecting JSON key should be always a string: %T: %v", t, t)
		}

		t, err = dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		var value interface{}
		value, err = handledelim(t, dec)
		if err != nil {
			return err
		}

		if existing, ok := r.m[key]; ok {
			if err := existing.Import(value); err != nil {
				return err
			}
		} else {
			r.keys[key] = r.l.PushBack(key)
			r.m[key] = NewValueAuto(value)
		}
	}

	t, err = dec.Token()
	if err != nil {
		return err
	}
	if delim, ok := t.(json.Delim); !ok || delim != '}' {
		return fmt.Errorf("expect JSON object close with '}'")
	}

	return nil
}

//nolint
func parsearray(dec *json.Decoder) (arr []interface{}, err error) {
	var t json.Token
	arr = make([]interface{}, 0)
	for dec.More() {
		t, err = dec.Token()
		if err != nil {
			return
		}

		var value interface{}
		value, err = handledelim(t, dec)
		if err != nil {
			return
		}
		arr = append(arr, value)
	}
	t, err = dec.Token()
	if err != nil {
		return
	}
	if delim, ok := t.(json.Delim); !ok || delim != ']' {
		err = fmt.Errorf("expect JSON array close with ']'")
		return
	}

	return
}

//nolint
func handledelim(t json.Token, dec *json.Decoder) (res interface{}, err error) {
	if delim, ok := t.(json.Delim); ok {
		switch delim {
		case '{':
			r2 := NewRow().(*row)
			err = r2.parseobject(dec)
			if err != nil {
				return
			}
			return r2, nil
		case '[':
			var value []interface{}
			value, err = parsearray(dec)
			if err != nil {
				return
			}
			return value, nil
		default:
			return nil, fmt.Errorf("Unexpected delimiter: %q", delim)
		}
	}
	return t, nil
}

func (r *row) GetString(key string) string {
	result, _ := cast.ToString(r.GetOrNil(key))

	return result.(string)
}

func (r *row) GetInt(key string) int {
	result, _ := cast.ToInt(r.GetOrNil(key))

	return result.(int)
}

func (r *row) GetInt64(key string) int64 {
	result, _ := cast.ToInt64(r.GetOrNil(key))

	return result.(int64)
}

func (r *row) GetInt32(key string) int32 {
	result, _ := cast.ToInt32(r.GetOrNil(key))

	return result.(int32)
}

func (r *row) GetInt16(key string) int16 {
	result, _ := cast.ToInt16(r.GetOrNil(key))

	return result.(int16)
}

func (r *row) GetInt8(key string) int8 {
	result, _ := cast.ToInt8(r.GetOrNil(key))

	return result.(int8)
}

func (r *row) GetUint(key string) uint {
	result, _ := cast.ToUint(r.GetOrNil(key))

	return result.(uint)
}

func (r *row) GetUint64(key string) uint64 {
	result, _ := cast.ToUint64(r.GetOrNil(key))

	return result.(uint64)
}

func (r *row) GetUint32(key string) uint32 {
	result, _ := cast.ToUint32(r.GetOrNil(key))

	return result.(uint32)
}

func (r *row) GetUint16(key string) uint16 {
	result, _ := cast.ToUint16(r.GetOrNil(key))

	return result.(uint16)
}

func (r *row) GetUint8(key string) uint8 {
	result, _ := cast.ToUint8(r.GetOrNil(key))

	return result.(uint8)
}

func (r *row) GetFloat64(key string) float64 {
	result, _ := cast.ToFloat64(r.GetOrNil(key))

	return result.(float64)
}

func (r *row) GetFloat32(key string) float32 {
	result, _ := cast.ToFloat32(r.GetOrNil(key))

	return result.(float32)
}

func (r *row) GetBool(key string) bool {
	result, _ := cast.ToBool(r.GetOrNil(key))

	return result.(bool)
}

func (r *row) GetBytes(key string) []byte {
	result, _ := cast.ToBinary(r.GetOrNil(key))

	return result.([]byte)
}

func (r *row) GetTime(key string) time.Time {
	result, _ := cast.ToTime(r.GetOrNil(key))

	return result.(time.Time)
}
