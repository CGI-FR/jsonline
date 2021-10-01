package jsonline_test

import (
	"encoding/json"
	"testing"

	"github.com/adrienaury/go-template/pkg/jsonline"
	"github.com/stretchr/testify/assert"
)

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
		WithMappedString("title", []byte{}).
		WithMappedNumeric("release_date", []byte{})

	row, err := template.CreateRow(data)
	assert.NoError(t, err)
	assert.Equal(t, data, row.Raw())

	export, err := row.Export()
	assert.NoError(t, err)
	assert.Equal(t, expectedExport, export)

	b, err := row.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, expectedMarshal, string(b))

	result := template.CreateRowEmpty()
	assert.Equal(t, expectedRawEmpty, result.Raw())
	assert.Equal(t, expectedExportEmpty, result.String())

	err = result.UnmarshalJSON(b)
	assert.NoError(t, err)
	assert.Equal(t, data, result.Raw())
}