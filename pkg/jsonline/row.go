package jsonline

import (
	"container/list"
	"encoding/json"
	"fmt"
)

// Row of data.
type Row interface {
	Set(key string, val Value) Row

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

func (r *row) Export() interface{} {
	return r
}

func (r *row) Import(v interface{}) Value {
	switch values := v.(type) {
	case []interface{}:
		//
	case map[string]interface{}:
		for key, val := range values {
			r.Set(key, NewValueAuto(val))
		}
	}

	return r
}

func (r *row) Set(key string, val Value) Row {
	if _, ok := r.m[key]; !ok {
		r.keys[key] = r.l.PushBack(key)
	}

	r.m[key] = val

	return r
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
