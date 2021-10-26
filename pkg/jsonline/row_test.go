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

package jsonline_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/cgi-fr/jsonline/pkg/jsonline"
	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	r1 := jsonline.NewRow()
	r1.SetValue("string", jsonline.NewValueString("value"))
	r1.SetValue("numeric", jsonline.NewValueNumeric(42.5))
	r1.SetValue("boolean", jsonline.NewValueBoolean(true))
	r1.SetValue("binary", jsonline.NewValueBinary("value"))
	r1.SetValue("datetime", jsonline.NewValueDateTime(time.Date(2021, time.September, 24, 21, 21, 0, 0, time.UTC)))
	r1.SetValue("timestamp", jsonline.NewValueTimestamp(time.Now()))

	r2 := jsonline.NewRow()

	err := r2.UnmarshalJSON([]byte(r1.String()))
	assert.NoError(t, err)

	assert.Equal(t, r1.String(), r2.String())
}

func TestExportImport(t *testing.T) {
	r1 := jsonline.NewRow()
	r1.SetValue("string", jsonline.NewValueString("value"))
	r1.SetValue("numeric", jsonline.NewValueNumeric(42.5))
	r1.SetValue("boolean", jsonline.NewValueBoolean(true))
	r1.SetValue("binary", jsonline.NewValueBinary("value"))
	r1.SetValue("datetime", jsonline.NewValueDateTime(time.Date(2021, time.September, 24, 21, 21, 0, 0, time.UTC)))
	r1.SetValue("timestamp", jsonline.NewValueTimestamp(time.Date(2021, time.September, 24, 21, 21, 0, 0, time.UTC)))

	res, err := r1.Export()
	assert.NoError(t, err)

	assert.IsType(t, map[string]interface{}{}, res)

	m, _ := res.(map[string]interface{})

	assert.Equal(t, "value", m["string"])
	assert.Equal(t, json.Number("42.5"), m["numeric"])
	assert.Equal(t, true, m["boolean"])
	assert.Equal(t, "dmFsdWU=", m["binary"])
	assert.Equal(t, "2021-09-24T21:21:00Z", m["datetime"])
	assert.Equal(t, int64(1632518460), m["timestamp"])

	r2 := jsonline.NewRow()
	err = r2.Import(m)
	assert.NoError(t, err)

	res, err = r2.Export()
	assert.NoError(t, err)

	assert.IsType(t, map[string]interface{}{}, res)

	m, _ = res.(map[string]interface{})

	assert.Equal(t, "value", m["string"])
	assert.Equal(t, json.Number("42.5"), m["numeric"])
	assert.Equal(t, true, m["boolean"])
	assert.Equal(t, "dmFsdWU=", m["binary"])
	assert.Equal(t, "2021-09-24T21:21:00Z", m["datetime"])
	assert.Equal(t, int64(1632518460), m["timestamp"])

	r3 := jsonline.NewRow()
	r3.SetValue("string", jsonline.NewValueString(nil))
	r3.SetValue("numeric", jsonline.NewValueNumeric(nil))
	r3.SetValue("boolean", jsonline.NewValueBoolean(nil))
	r3.SetValue("binary", jsonline.NewValueBinary(nil))
	r3.SetValue("datetime", jsonline.NewValueDateTime(nil))
	r3.SetValue("timestamp", jsonline.NewValueTimestamp(nil))

	err = r3.Import([]interface{}{"value", 42.5, true, "dmFsdWU=", "2021-09-24T21:21:00Z", 1632518460})
	assert.NoError(t, err)

	assert.Equal(t, r1.String(), r3.String())
}

func TestGetOrNil(t *testing.T) {
	r1 := jsonline.NewRow()
	r1.SetValue("string", jsonline.NewValueString("value"))

	assert.Equal(t, nil, r1.GetOrNil("nokey"))
	assert.Equal(t, "value", r1.GetOrNil("string"))

	assert.Equal(t, nil, r1.GetAtIndexOrNil(2))
	assert.Equal(t, "value", r1.GetAtIndexOrNil(0))
}

func TestMap(t *testing.T) {
	r1 := jsonline.NewRow()
	r1.SetValue("int", jsonline.NewValueNumeric(-1))
	r1.SetValue("uint", jsonline.NewValueNumeric(uint(1)))
	r1.SetValue("string", jsonline.NewValueString("value"))
	r1.SetValue("numeric", jsonline.NewValueNumeric(42.5))
	r1.SetValue("boolean", jsonline.NewValueBoolean(true))
	r1.SetValue("binary", jsonline.NewValue("value", jsonline.Binary, []byte{}))

	result := struct {
		Int     int
		Uint    uint
		String  string
		Numeric float32
		Boolean bool
		Binary  []byte
	}{}

	r1.MapTo(&result)

	expected := struct {
		Int     int
		Uint    uint
		String  string
		Numeric float32
		Boolean bool
		Binary  []byte
	}{
		Int:     -1,
		Uint:    1,
		String:  "value",
		Numeric: 42.5,
		Boolean: true,
		Binary:  []byte("value"),
	}

	assert.Equal(t, expected, result)
}
