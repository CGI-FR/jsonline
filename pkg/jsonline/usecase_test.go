package jsonline_test

import (
	"encoding/json"
	"testing"

	"github.com/cgi-fr/jsonline/pkg/jsonline"
	"github.com/stretchr/testify/assert"
)

func TestUseCase(t *testing.T) {
	tmpl := jsonline.NewTemplate().
		WithBinary("from string").
		WithBinary("from byte array").
		WithBinary("from byte slice").
		WithBinary("from int")

	r, err := tmpl.CreateRow(map[string]interface{}{
		"from string":     "hello",
		"from byte array": [2]byte{1, 2},
		"from byte slice": []byte{1, 2},
		"from int":        1,
	})
	assert.NoError(t, err)
	assert.Equal(t, `{"from string":"aGVsbG8=","from byte array":"AQI=","from byte slice":"AQI=","from int":"AQAAAAAAAAA="}`, r.String()) //nolint:lll
}

func TestLinoUseCase(t *testing.T) {
	data := map[string]interface{}{
		"title":        []byte("The Matrix"),
		"release_date": []byte("1999"),
	}

	expectedExport := map[string]interface{}{"release_date": json.Number("1999"), "title": "The Matrix"}
	expectedMarshal := `{"title":"The Matrix","release_date":1999}`
	expectedRawEmpty := map[string]interface{}{"release_date": nil, "title": nil}
	expectedExportEmpty := `{"title":null,"release_date":null}`

	template := jsonline.NewTemplate().
		WithString("title").
		WithNumeric("release_date")

	row, err := template.CreateRow(data)
	assert.NoError(t, err)
	assert.Equal(t, data, row.Raw())

	export, err := row.Export()
	assert.NoError(t, err)
	assert.Equal(t, expectedExport, export)

	b, err := row.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, expectedMarshal, string(b))

	template = jsonline.NewTemplate().
		WithMappedAuto("title", []byte{}).
		WithMappedAuto("release_date", []byte{})

	result := template.CreateRowEmpty()
	assert.Equal(t, expectedRawEmpty, result.Raw())
	assert.Equal(t, expectedExportEmpty, result.String())

	err = result.UnmarshalJSON(b)
	assert.NoError(t, err)
	assert.Equal(t, data, result.Raw())
}

func BenchmarkLinoUseCase(b *testing.B) {
	data := map[string]interface{}{
		"title":        []byte("The Matrix"),
		"release_date": []byte("1999"),
	}

	for n := 0; n < b.N; n++ {
		// lino pull
		template := jsonline.NewTemplate().
			WithString("title").
			WithNumeric("release_date")

		row, _ := template.CreateRow(data)
		b, _ := row.MarshalJSON() // output data to the stdout

		// lino push
		template = jsonline.NewTemplate().
			WithMappedAuto("title", []byte{}).
			WithMappedAuto("release_date", []byte{})

		result := template.CreateRowEmpty()
		result.UnmarshalJSON(b) //nolint:errcheck
	}
}
