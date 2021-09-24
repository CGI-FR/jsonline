package jsonline_test

import (
	"fmt"
	"testing"

	"github.com/adrienaury/go-template/pkg/jsonline"
	"github.com/stretchr/testify/assert"
)

func TestValueExportString(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
	}{
		{nil, nil},
		{1, "1"},
		{1.2, "1.2"},
		{"string", "string"},
		{true, "true"},
		{[]string{"a", "b"}, "[a b]"},
		{[]byte("hello"), "hello"},
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValue(td.value)
			value.SetExportFormat(jsonline.String)

			assert.Equal(t, td.expected, value.Export())
		})
	}
}
