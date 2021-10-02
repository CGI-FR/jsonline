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
	WithTimestamp(string) Template
	WithAuto(string) Template
	WithHidden(string) Template
	WithRow(string, Template) Template

	WithMappedString(string, interface{}) Template
	WithMappedNumeric(string, interface{}) Template
	WithMappedBoolean(string, interface{}) Template
	WithMappedBinary(string, interface{}) Template
	WithMappedDateTime(string, interface{}) Template
	WithMappedTimestamp(string, interface{}) Template
	WithMappedAuto(string, interface{}) Template

	CreateRow(interface{}) (Row, error)
	CreateRowEmpty() Row

	GetExporter(io.Writer) Exporter
	GetImporter(io.Reader) Importer
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
	t.empty.Set(name, rowt.CreateRowEmpty())

	return t
}

func (t *template) WithMappedString(name string, rawtype interface{}) Template {
	t.empty.Set(name, NewValue(nil, String, rawtype))

	return t
}

func (t *template) WithMappedNumeric(name string, rawtype interface{}) Template {
	t.empty.Set(name, NewValue(nil, Numeric, rawtype))

	return t
}

func (t *template) WithMappedBoolean(name string, rawtype interface{}) Template {
	t.empty.Set(name, NewValue(nil, Boolean, rawtype))

	return t
}

func (t *template) WithMappedBinary(name string, rawtype interface{}) Template {
	t.empty.Set(name, NewValue(nil, Binary, rawtype))

	return t
}

func (t *template) WithMappedDateTime(name string, rawtype interface{}) Template {
	t.empty.Set(name, NewValue(nil, DateTime, rawtype))

	return t
}

func (t *template) WithMappedTimestamp(name string, rawtype interface{}) Template {
	t.empty.Set(name, NewValue(nil, Timestamp, rawtype))

	return t
}

func (t *template) WithMappedAuto(name string, rawtype interface{}) Template {
	t.empty.Set(name, NewValue(nil, Auto, rawtype))

	return t
}

//nolint:cyclop,funlen
func (t *template) CreateRow(v interface{}) (Row, error) {
	result := CloneRow(t.empty)

	switch values := v.(type) {
	case []interface{}:
		log.Debug().Msg("create new row from slice")

		for i, val := range values {
			target := result.GetAtIndex(i)
			if target != nil {
				target = NewValue(val, target.GetFormat(), target.GetRawType())
			} else {
				target = NewValueAuto(val)
			}

			result.SetAtIndex(i, target)
		}
	case map[string]interface{}:
		log.Debug().Msg("create new row from map")

		for key, val := range values {
			target := result.Get(key)
			if target != nil {
				target = NewValue(val, target.GetFormat(), target.GetRawType())
			} else {
				target = NewValueAuto(val)
			}

			result.Set(key, target)
		}
	case Row:
		log.Debug().Msg("create new row from row")

		iter := values.Iter()

		for key, val, ok := iter(); ok; key, val, ok = iter() {
			target := result.Get(key)
			if target != nil {
				target = NewValue(val.Raw(), target.GetFormat(), target.GetRawType())
			} else {
				target = NewValueAuto(val.Raw())
			}

			result.Set(key, target)
		}

	case []byte:
		log.Debug().Msg("create new row from byte buffer")

		if err := result.UnmarshalJSON(values); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

	case string:
		log.Debug().Msg("create new row from JSON string")

		if err := result.UnmarshalJSON([]byte(values)); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

	default:
		return nil, fmt.Errorf("%w", ErrUnsupportedImportType)
	}

	return result, nil
}

func (t *template) CreateRowEmpty() Row {
	return CloneRow(t.empty)
}

func (t *template) GetExporter(w io.Writer) Exporter {
	return NewExporter(w).WithTemplate(t)
}

func (t *template) GetImporter(r io.Reader) Importer {
	return NewImporter(r).WithTemplate(t)
}
