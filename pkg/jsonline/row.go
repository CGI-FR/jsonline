package jsonline

import (
	"container/list"
	"encoding/json"
	"fmt"
)

// Row of data.
type Row interface {
	Set(key string, val Value) Row
	SetAtIndex(index int, val Value) Row

	ImportAtKey(key string, val interface{}) Row
	ImportAtIndex(index int, val interface{}) Row

	Get(key string) Value

	Clone() Row

	Value
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

func (r *row) GetFormat() format {
	return Auto
}

func (r *row) Raw() interface{} {
	return r
}

func (r *row) Export() interface{} {
	result := map[string]interface{}{}

	iter := r.Iter()

	for k, v, ok := iter(); ok; k, v, ok = iter() {
		result[k] = v.Export()
	}

	return result
}

func (r *row) Import(v interface{}) Value {
	switch values := v.(type) {
	case []interface{}:
		for i, val := range values {
			r.ImportAtIndex(i, val)
		}
	case map[string]interface{}:
		for key, val := range values {
			r.ImportAtKey(key, val)
		}
	}

	return r
}

func (r *row) ImportAtKey(key string, val interface{}) Row {
	if _, ok := r.m[key]; !ok {
		r.keys[key] = r.l.PushBack(key)
	}

	if value, exist := r.m[key]; exist {
		r.m[key] = value.Import(val)
	} else {
		r.m[key] = NewValueAuto(val)
	}

	return r
}

func (r *row) ImportAtIndex(index int, val interface{}) Row {
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

func (r *row) Set(key string, val Value) Row {
	if _, ok := r.m[key]; !ok {
		r.keys[key] = r.l.PushBack(key)
	}

	r.m[key] = val

	return r
}

func (r *row) SetAtIndex(index int, val Value) Row {
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

func (r *row) Get(key string) Value {
	return r.m[key]
}

func (r *row) Clone() Row {
	result := NewRow()

	for e := r.l.Front(); e != nil; e = e.Next() {
		k, _ := e.Value.(string)
		result.Set(k, r.m[k])
	}

	return result
}

func (r *row) Iter() func() (string, Value, bool) {
	e := r.l.Front()

	return func() (string, Value, bool) {
		if e != nil {
			key, _ := e.Value.(string)
			e = e.Next()

			return key, r.m[key], true
		}

		return "", NewValueAuto(nil), false
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
