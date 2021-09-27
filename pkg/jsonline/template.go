package jsonline

import (
	"fmt"
	"io"

	"github.com/rs/zerolog/log"
)

type Template interface {
	WithString(string) Template
	WithNumeric(string) Template
	WithBoolean(string) Template
	WithBinary(string) Template
	WithDateTime(string) Template
	WithTime(string) Template
	WithTimestamp(string) Template
	WithAuto(string) Template
	WithHidden(string) Template
	WithRow(string, Template) Template

	Create(interface{}) Row
	CreateEmpty() Row

	GetExporter(io.Writer) Exporter
}

type template struct {
	empty Row
}

func NewTemplate() Template {
	return &template{
		empty: NewRow(),
	}
}

func (t *template) WithString(name string) Template {
	t.empty.Set(name, NewValueString(nil))

	return t
}

func (t *template) WithNumeric(name string) Template {
	t.empty.Set(name, NewValueNumeric(nil))

	return t
}

func (t *template) WithBoolean(name string) Template {
	t.empty.Set(name, NewValueBoolean(nil))

	return t
}

func (t *template) WithBinary(name string) Template {
	t.empty.Set(name, NewValueBinary(nil))

	return t
}

func (t *template) WithDateTime(name string) Template {
	t.empty.Set(name, NewValueDateTime(nil))

	return t
}

func (t *template) WithTime(name string) Template {
	t.empty.Set(name, NewValueTime(nil))

	return t
}

func (t *template) WithTimestamp(name string) Template {
	t.empty.Set(name, NewValueTimestamp(nil))

	return t
}

func (t *template) WithAuto(name string) Template {
	t.empty.Set(name, NewValueAuto(nil))

	return t
}

func (t *template) WithHidden(name string) Template {
	t.empty.Set(name, NewValueHidden(nil))

	return t
}

func (t *template) WithRow(name string, rowt Template) Template {
	t.empty.Set(name, rowt.CreateEmpty())

	return t
}

func (t *template) Create(v interface{}) Row {
	result := CloneRow(t.empty)

	switch values := v.(type) {
	case []interface{}:
		log.Debug().Msg("create new row from slice")

		for i, val := range values {
			result.ImportAtIndex(i, val)
		}
	case map[string]interface{}:
		log.Debug().Msg("create new row from map")

		for key, val := range values {
			result.ImportAtKey(key, val)
		}

	default:
		log.Warn().Str("type", fmt.Sprintf("%T", values)).Msg("can't create row from this type")
	}

	return result
}

func (t *template) CreateEmpty() Row {
	return CloneRow(t.empty)
}

func (t *template) GetExporter(w io.Writer) Exporter {
	return NewExporter(w).WithTemplate(t)
}
