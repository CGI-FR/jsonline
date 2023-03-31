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

//nolint:varnamelen
package jsonline_test

import (
	"os"
	"testing"

	"github.com/cgi-fr/jsonline/pkg/jsonline"
	"github.com/stretchr/testify/assert"
)

func init() { //nolint:gochecknoinits
	os.Setenv("TZ", "Europe/Paris")
}

func TestTemplateCreateEmpty(t *testing.T) {
	template := jsonline.NewTemplate().
		WithString("string").
		WithNumeric("numeric").
		WithBoolean("boolean").
		WithBinary("binary").
		WithDateTime("datetime").
		WithTimestamp("timestamp").
		WithHidden("hidden").
		WithAuto("auto").
		WithRow("row", jsonline.NewTemplate())

	row := template.CreateRowEmpty()

	assert.Equal(t,
		`{"string":null,"numeric":null,"boolean":null,"binary":null,"datetime":null,"timestamp":null,"auto":null,"row":{}}`,
		row.String())
}

func TestTemplateCreateFromSlice(t *testing.T) {
	template := jsonline.NewTemplate().
		WithString("string").
		WithNumeric("numeric").
		WithBoolean("boolean").
		WithBinary("binary").
		WithDateTime("datetime").
		WithTimestamp("timestamp").
		WithHidden("hidden").
		WithAuto("auto").
		WithRow("row", jsonline.NewTemplate())

	sub := map[string]interface{}{
		"first": 0,
	}

	row, err := template.CreateRow([]interface{}{"value", 0, true, "value", 1566844858, 1566844858, "hidden", "hello", sub, "extra1", "extra2"}) //nolint:lll
	assert.NoError(t, err)

	//nolint:lll
	assert.Equal(t,
		`{"string":"value","numeric":0,"boolean":true,"binary":"dmFsdWU=","datetime":"2019-08-26T20:40:58+02:00","timestamp":1566844858,"auto":"hello","row":{"first":0},"":"extra2"}`,
		row.String())
}

func TestTemplateCreateFromMap(t *testing.T) {
	template := jsonline.NewTemplate().
		WithString("string").
		WithNumeric("numeric").
		WithBoolean("boolean").
		WithBinary("binary").
		WithDateTime("datetime").
		WithTimestamp("timestamp").
		WithHidden("hidden").
		WithAuto("auto").
		WithRow("row", jsonline.NewTemplate())

	sub := map[string]interface{}{
		"first": 0,
	}

	row, err := template.CreateRow(map[string]interface{}{
		"datetime":  1566844858,
		"numeric":   0,
		"boolean":   true,
		"extra":     "extra",
		"timestamp": 1566844858,
		"binary":    "value",
		"hidden":    "hidden",
		"auto":      "hello",
		"row":       sub,
		"string":    "value",
	})
	assert.NoError(t, err)

	//nolint:lll
	assert.Equal(t,
		`{"string":"value","numeric":0,"boolean":true,"binary":"dmFsdWU=","datetime":"2019-08-26T20:40:58+02:00","timestamp":1566844858,"auto":"hello","row":{"first":0},"extra":"extra"}`,
		row.String())
}

func TestTemplateCreateFromRow(t *testing.T) {
	template := jsonline.NewTemplate().
		WithString("string").
		WithNumeric("numeric").
		WithBoolean("boolean").
		WithBinary("binary").
		WithDateTime("datetime").
		WithTimestamp("timestamp").
		WithHidden("hidden").
		WithAuto("auto").
		WithRow("row", jsonline.NewTemplate())

	sub := map[string]interface{}{
		"first": 0,
	}

	rsrc := jsonline.NewRow().
		SetValue("datetime", jsonline.NewValueDateTime(1566844858)).
		SetValue("numeric", jsonline.NewValueNumeric(0)).
		SetValue("boolean", jsonline.NewValueBoolean(true)).
		SetValue("extra", jsonline.NewValueAuto("extra")).
		SetValue("timestamp", jsonline.NewValueTimestamp(1566844858)).
		SetValue("binary", jsonline.NewValueBinary("value")).
		SetValue("hidden", jsonline.NewValueHidden("hidden")).
		SetValue("auto", jsonline.NewValueAuto("hello")).
		SetValue("row", jsonline.NewValueAuto(sub)).
		SetValue("string", jsonline.NewValueString("value"))

	row, err := template.CreateRow(rsrc)
	assert.NoError(t, err)

	//nolint:lll
	assert.Equal(t,
		`{"string":"value","numeric":0,"boolean":true,"binary":"dmFsdWU=","datetime":"2019-08-26T20:40:58+02:00","timestamp":1566844858,"auto":"hello","row":{"first":0},"extra":"extra"}`,
		row.String())
}

func TestTemplateCreateFromString(t *testing.T) {
	template := jsonline.NewTemplate().
		WithString("string").
		WithNumeric("numeric").
		WithBoolean("boolean").
		WithBinary("binary").
		WithDateTime("datetime").
		WithTimestamp("timestamp").
		WithHidden("hidden").
		WithAuto("auto").
		WithRow("row", jsonline.NewTemplate())

	//nolint:lll
	row, err := template.CreateRow(`{"string":"value","numeric":0,"boolean":true,"binary":"dmFsdWU=","datetime":"2019-08-26T20:40:58+02:00","timestamp":1566844858,"auto":"hello","row":{"first":0},"extra":"extra"}`)
	assert.NoError(t, err)

	//nolint:lll
	assert.Equal(t,
		`{"string":"value","numeric":0,"boolean":true,"binary":"dmFsdWU=","datetime":"2019-08-26T20:40:58+02:00","timestamp":1566844858,"auto":"hello","row":{"first":0},"extra":"extra"}`,
		row.String())
}

func TestTemplateCreateFromByteBuffer(t *testing.T) {
	template := jsonline.NewTemplate().
		WithString("string").
		WithNumeric("numeric").
		WithBoolean("boolean").
		WithBinary("binary").
		WithDateTime("datetime").
		WithTimestamp("timestamp").
		WithHidden("hidden").
		WithAuto("auto").
		WithRow("row", jsonline.NewTemplate())

	//nolint:lll
	row, err := template.CreateRow([]byte(`{"string":"value","numeric":0,"boolean":true,"binary":"dmFsdWU=","datetime":"2019-08-26T20:40:58+02:00","timestamp":1566844858,"auto":"hello","row":{"first":0},"extra":"extra"}`))
	assert.NoError(t, err)

	//nolint:lll
	assert.Equal(t,
		`{"string":"value","numeric":0,"boolean":true,"binary":"dmFsdWU=","datetime":"2019-08-26T20:40:58+02:00","timestamp":1566844858,"auto":"hello","row":{"first":0},"extra":"extra"}`,
		row.String())
}

func TestChaining(t *testing.T) {
	t1 := jsonline.NewTemplate().WithMappedBinary("input", int64(0))
	t2 := jsonline.NewTemplate().WithNumeric("input")

	row1, err := t1.CreateRow(`{"input":"AgAAAAAAAAA="}`)
	assert.NoError(t, err)
	assert.Equal(t, row1.DebugString(), "input={int64(2)}")

	row2 := t2.CreateRowEmpty()
	err = row2.Import(row1.Raw())
	assert.NoError(t, err)
	assert.Equal(t, row2.DebugString(), "input={json.Number(`2`)}")

	assert.Equal(t, row2.String(), `{"input":2}`)
}
