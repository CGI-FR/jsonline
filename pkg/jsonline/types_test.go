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

func TestValueMarshalString(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
	}{
		{nil, "null"},
		{1, "\"1\""},
		{1.2, "\"1.2\""},
		{"string", "\"string\""},
		{true, "\"true\""},
		{[]string{"a", "b"}, "\"[a b]\""},
		{[]byte("hello"), "\"hello\""},
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValue(td.value)
			value.SetExportFormat(jsonline.String)

			assert.Equal(t, td.expected, value.String())
		})
	}
}

func TestValueExportNumeric(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
	}{
		{nil, nil},
		{1, 1},
		{1.2, 1.2},
		{"1.2", 1.2},
		{true, 1},
		{false, 0},
		// {[]string{"1", "2.4"}, []float64{1, 2.4}},
		// {[]string{"1", "2"}, []int{1, 2}},
		{[]byte("1.5"), 1.5},
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValue(td.value)
			value.SetExportFormat(jsonline.Numeric)

			assert.Equal(t, td.expected, value.Export())
		})
	}
}

func TestValueMarshalNumeric(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
	}{
		{nil, "null"},
		{1, "1"},
		{1.2, "1.2"},
		{"1.2", "1.2"},
		{true, "1"},
		{false, "0"},
		// {[]string{"1", "2.4"}, []float64{1, 2.4}},
		// {[]string{"1", "2"}, []int{1, 2}},
		{[]byte("1.5"), "1.5"},
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValue(td.value)
			value.SetExportFormat(jsonline.Numeric)

			assert.Equal(t, td.expected, value.String())
		})
	}
}

func TestValueExportBoolean(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
	}{
		{nil, nil},
		{1, true},
		{"1", true},
		{"true", true},
		{"false", false},
		{0, false},
		{"0", false},
		{1.2, true},
		{0.0, false},
		{"1.2", true},
		{"0.0", false},
		{true, true},
		{false, false},
		{[]byte("true"), true},
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValue(td.value)
			value.SetExportFormat(jsonline.Boolean)

			assert.Equal(t, td.expected, value.Export())
		})
	}
}

func TestValueMarshalBoolean(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
	}{
		{nil, "null"},
		{1, "1"},
		{1.2, "1.2"},
		{"1.2", "1.2"},
		{true, "1"},
		{false, "0"},
		// {[]string{"1", "2.4"}, []float64{1, 2.4}},
		// {[]string{"1", "2"}, []int{1, 2}},
		{[]byte("1.5"), "1.5"},
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValue(td.value)
			value.SetExportFormat(jsonline.Numeric)

			assert.Equal(t, td.expected, value.String())
		})
	}
}
