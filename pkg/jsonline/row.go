package jsonline

import "encoding/json"

// Row of data.
type Row interface {
	Get(key string) Value
	GetIfHas(key string) (Value, bool)
	Has(key string) bool
	Set(key string, val Value)
	Delete(key string) (Value, bool)
	Export() map[string]interface{}
	Iter() func() (string, Value, bool)
	Len() int
	Update(other Row) Row
	json.Marshaler
}
