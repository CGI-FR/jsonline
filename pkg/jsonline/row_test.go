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
	"testing"
	"time"

	"github.com/adrienaury/go-template/pkg/jsonline"
	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	r1 := jsonline.NewRow()
	r1.Set("string", jsonline.NewValueString("value"))
	r1.Set("numeric", jsonline.NewValueNumeric(42.5))
	r1.Set("boolean", jsonline.NewValueBoolean(true))
	r1.Set("binary", jsonline.NewValueBinary("value"))
	r1.Set("datetime", jsonline.NewValueDateTime(time.Date(2021, time.September, 24, 21, 21, 0, 0, time.UTC)))
	r1.Set("time", jsonline.NewValueTime(time.Now()))
	r1.Set("timestamp", jsonline.NewValueTimestamp(time.Now()))

	r2 := jsonline.NewRow()

	err := r2.UnmarshalJSON([]byte(r1.String()))
	assert.NoError(t, err)

	assert.Equal(t, r1.String(), r2.String())
}

//nolint:funlen
func TestExportImport(t *testing.T) {
	r1 := jsonline.NewRow()
	r1.Set("string", jsonline.NewValueString("value"))
	r1.Set("numeric", jsonline.NewValueNumeric(42.5))
	r1.Set("boolean", jsonline.NewValueBoolean(true))
	r1.Set("binary", jsonline.NewValueBinary("value"))
	r1.Set("datetime", jsonline.NewValueDateTime(time.Date(2021, time.September, 24, 21, 21, 0, 0, time.UTC)))
	r1.Set("time", jsonline.NewValueTime(time.Date(2021, time.September, 24, 21, 21, 0, 0, time.UTC)))
	r1.Set("timestamp", jsonline.NewValueTimestamp(time.Date(2021, time.September, 24, 21, 21, 0, 0, time.UTC)))

	res, err := r1.Export()
	assert.NoError(t, err)

	assert.IsType(t, map[string]interface{}{}, res)

	m, _ := res.(map[string]interface{})

	assert.Equal(t, "value", m["string"])
	assert.Equal(t, float64(42.5), m["numeric"])
	assert.Equal(t, true, m["boolean"])
	assert.Equal(t, "dmFsdWU=", m["binary"])
	assert.Equal(t, "2021-09-24T21:21:00Z", m["datetime"])
	assert.Equal(t, "21:21:00Z", m["time"])
	assert.Equal(t, int64(1632518460), m["timestamp"])

	r2 := jsonline.NewRow()
	err = r2.Import(m)
	assert.NoError(t, err)

	res, err = r2.Export()
	assert.NoError(t, err)

	assert.IsType(t, map[string]interface{}{}, res)

	m, _ = res.(map[string]interface{})

	assert.Equal(t, "value", m["string"])
	assert.Equal(t, float64(42.5), m["numeric"])
	assert.Equal(t, true, m["boolean"])
	assert.Equal(t, "dmFsdWU=", m["binary"])
	assert.Equal(t, "2021-09-24T21:21:00Z", m["datetime"])
	assert.Equal(t, "21:21:00Z", m["time"])
	assert.Equal(t, int64(1632518460), m["timestamp"])

	r3 := jsonline.NewRow()
	r3.Set("string", jsonline.NewValueNil(jsonline.String))
	r3.Set("numeric", jsonline.NewValueNil(jsonline.Numeric))
	r3.Set("boolean", jsonline.NewValueNil(jsonline.Boolean))
	r3.Set("binary", jsonline.NewValueNil(jsonline.Binary))
	r3.Set("datetime", jsonline.NewValueNil(jsonline.DateTime))
	r3.Set("time", jsonline.NewValueNil(jsonline.Time))
	r3.Set("timestamp", jsonline.NewValueNil(jsonline.Timestamp))

	err = r3.Import([]interface{}{"value", 42.5, true, "dmFsdWU=", "2021-09-24T21:21:00Z", "21:21:00Z", 1632518460})
	assert.NoError(t, err)

	assert.Equal(t, r1.String(), r3.String())
}
