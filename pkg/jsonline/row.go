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
	"bytes"
	"container/list"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Row interface {
	Value

	Has(key string) bool
	Get(key string) (interface{}, bool)
	GetAtIndex(index int) (interface{}, bool)
	Set(key string, val interface{}) error
	SetAtIndex(index int, val interface{}) error
	Len() int
	Iter() func() (string, interface{}, bool)

	GetValue(key string) (Value, bool)
	GetValueAtIndex(index int) (Value, bool)
	SetValue(key string, val Value) Row
	SetValueAtIndex(index int, val Value) Row
	IterValues() func() (string, Value, bool)
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
			if err := r.SetAtIndex(i, val); err != nil {
				return err
			}
		}
	case map[string]interface{}:
		for key, val := range values {
			if err := r.Set(key, val); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("%w", ErrUnsupportedImportType)
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

func (r *row) Set(key string, val interface{}) error {
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

func (r *row) SetAtIndex(index int, val interface{}) error {
	var key string

	for cur := r.l.Front(); cur != nil; cur = cur.Next() {
		if index == 0 {
			key, _ = cur.Value.(string)

			break
		}
		index--
	}

	return r.Set(key, val)
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
