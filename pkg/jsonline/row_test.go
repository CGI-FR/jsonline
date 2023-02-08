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

package jsonline_test

import (
	"encoding/json"
	"fmt"
	"strings"
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

func TestGetTypedValue(t *testing.T) {
	r1 := jsonline.NewRow()
	r1.SetValue("string", jsonline.NewValueString("value"))
	r1.SetValue("numeric", jsonline.NewValueNumeric(42.5))
	r1.SetValue("boolean", jsonline.NewValueBoolean(true))
	r1.SetValue("binary", jsonline.NewValueBinary("value"))
	r1.SetValue("datetime", jsonline.NewValueDateTime(time.Date(2021, time.September, 24, 21, 21, 0, 0, time.UTC)))

	assert.Equal(t, "value", r1.GetString("string"))
	assert.Equal(t, 42, r1.GetInt("numeric"))
	assert.Equal(t, int64(42), r1.GetInt64("numeric"))
	assert.Equal(t, int32(42), r1.GetInt32("numeric"))
	assert.Equal(t, int16(42), r1.GetInt16("numeric"))
	assert.Equal(t, int8(42), r1.GetInt8("numeric"))
	assert.Equal(t, uint(42), r1.GetUint("numeric"))
	assert.Equal(t, uint64(42), r1.GetUint64("numeric"))
	assert.Equal(t, uint32(42), r1.GetUint32("numeric"))
	assert.Equal(t, uint16(42), r1.GetUint16("numeric"))
	assert.Equal(t, uint8(42), r1.GetUint8("numeric"))
	assert.Equal(t, float64(42.5), r1.GetFloat64("numeric"))
	assert.Equal(t, float32(42.5), r1.GetFloat32("numeric"))
	assert.Equal(t, true, r1.GetBool("boolean"))
	assert.Equal(t, []byte("value"), r1.GetBytes("binary"))
	assert.Equal(t, time.Date(2021, time.September, 24, 21, 21, 0, 0, time.UTC), r1.GetTime("datetime"))
}

func TestDate(t *testing.T) {
	r1 := jsonline.NewRow()
	r1.SetValue("fromtime", jsonline.NewValue(time.Date(2021, time.September, 24, 21, 21, 0, 0, time.UTC), jsonline.Date, nil)) //nolint:lll
	r1.SetValue("fromstring", jsonline.NewValue("2021-09-24", jsonline.Date, nil))

	res, err := r1.Export()
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"fromstring": "2021-09-24", "fromtime": "2021-09-24"}, res)
}

func TestPath(t *testing.T) {
	r2 := jsonline.NewRow()
	r2.SetValue("sub1", jsonline.NewValueAuto(12))
	r2.SetValue("sub2", jsonline.NewValueAuto("hello"))

	r1 := jsonline.NewRow()
	r1.SetValue("root", r2)

	res1, exists1 := r1.GetAtPath("root.sub1")
	assert.True(t, exists1)
	assert.Equal(t, 12, res1)

	res2, exists2 := r1.GetAtPath("root.sub2")
	assert.True(t, exists2)
	assert.Equal(t, "hello", res2)

	res3, exists3 := r1.GetAtPath("root.sub3")
	assert.False(t, exists3)
	assert.Nil(t, res3)

	err1 := r1.ImportAtPath("root.sub1", 11)
	assert.NoError(t, err1)

	err2 := r1.ImportAtPath("root.sub2", "world")
	assert.NoError(t, err2)

	fmt.Println(r1.DebugString())
}

func TestPathArrays(t *testing.T) {
	row1str := `
	{
		"lignes": [
			{"nom": "ligne1", "details": [
				{"nom": "detail 1.1", "tag": 0},
				{"nom": "detail 1.2", "tag": 1},
				{"nom": "detail 1.3", "tag": 2}
			]},
			{"nom": "ligne2", "details": [
				{"nom": "detail 2.1", "tag": 3},
				{"nom": "detail 2.2", "tag": 4}
			]},
			{"nom": "ligne3", "details": []},
			{"nom": "ligne4"},
			{"nom": "ligne5", "details": [
				{"nom": "detail 5.1", "tag": 5},
				{"tag": -1},
				{"nom": null, "tag": 6}
			]}
		]
	}
	`

	row1, err1 := jsonline.NewImporter(strings.NewReader(strings.ReplaceAll(row1str, "\n", ""))).ReadOne()
	assert.NoError(t, err1)

	values1, exists1 := row1.FindValuesAtPath("lignes.details.nom")
	assert.True(t, exists1)

	expected1 := []interface{}{
		"detail 1.1",
		"detail 1.2",
		"detail 1.3",
		"detail 2.1",
		"detail 2.2",
		"detail 5.1",
		nil,
	}

	actual1 := []interface{}{}

	for _, value1 := range values1 {
		actual1 = append(actual1, value1.Raw())
	}

	assert.Equal(t, expected1, actual1)
}

func TestPathArrays2(t *testing.T) {
	row1str := `
	{
		"designation": "Acme Inc.",
		"adresseSiègeSocial": {
		  "NumeroRue": "123",
		  "NomRue": "Main St.",
		  "CodePostal": "12345",
		  "Ville": "Anytown"
		},
		"employe": [
		  {
			"Nom": "Jane",
			"Prenom": "Doe",
			"dateNaissance": "01-01-1970",
			"LieuNaissance": "Anytown",
			"NombreEnfants": 1,
			"enfant": [
			  {
				"Prenom": "John",
				"DateNaissance": "01-01-2000",
				"LieuNaissance": "Anytown"
			  },
			  {
				"Prenom": "Jack",
				"DateNaissance": "01-01-2002",
				"LieuNaissance": "Anytown"
			  }
			]
		  }
		]
	  }
	`

	row1, err1 := jsonline.NewImporter(strings.NewReader(strings.ReplaceAll(row1str, "\n", ""))).ReadOne()
	assert.NoError(t, err1)

	values1, exists1 := row1.FindValuesAtPath("employe.enfant")
	assert.True(t, exists1)

	result1, err2 := json.Marshal(values1)
	assert.NoError(t, err2)

	expected1 := `[[{"Prenom":"John","DateNaissance":"01-01-2000","LieuNaissance":"Anytown"},{"Prenom":"Jack","DateNaissance":"01-01-2002","LieuNaissance":"Anytown"}]]` //nolint: lll

	assert.Equal(t, expected1, string(result1))
}

func TestPathArrays3(t *testing.T) {
	row1str := `
	{
		"designation": "Acme Inc.",
		"adresseSiègeSocial": {
		  "NumeroRue": "123",
		  "NomRue": "Main St.",
		  "CodePostal": "12345",
		  "Ville": "Anytown"
		},
		"employe": [
		  {
			"Nom": "Jane",
			"Prenom": "Doe",
			"dateNaissance": "01-01-1970",
			"LieuNaissance": "Anytown",
			"NombreEnfants": 1,
			"enfant": [
			  {
				"Prenom": "John",
				"DateNaissance": "01-01-2000",
				"LieuNaissance": "Anytown"
			  },
			  {
				"Prenom": "Jack",
				"DateNaissance": "01-01-2002",
				"LieuNaissance": "Anytown"
			  }
			]
		  }
		]
	  }
	`

	row1, err1 := jsonline.NewImporter(strings.NewReader(strings.ReplaceAll(row1str, "\n", ""))).ReadOne()
	assert.NoError(t, err1)

	values1, exists1 := row1.FindValuesAtPath("employe.enfant.*")
	assert.True(t, exists1)

	result1, err2 := json.Marshal(values1)
	assert.NoError(t, err2)

	expected1 := `[{"Prenom":"John","DateNaissance":"01-01-2000","LieuNaissance":"Anytown"},{"Prenom":"Jack","DateNaissance":"01-01-2002","LieuNaissance":"Anytown"}]` //nolint: lll

	assert.Equal(t, expected1, string(result1))
}
