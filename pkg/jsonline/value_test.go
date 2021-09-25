package jsonline_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/adrienaury/go-template/pkg/jsonline"
	"github.com/stretchr/testify/assert"
)

func TestValueFormatString(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
	}{
		{nil, nil},
		// signed integers
		{int(-2), "-2"},
		{int8(-1), "-1"},
		{int16(0), "0"},
		{int32(1), "1"},
		{int64(2), "2"},
		// unsigned integers
		{uint(0), "0"},
		{uint8(1), "1"},
		{uint16(2), "2"},
		{uint32(3), "3"},
		{uint64(4), "4"},
		// floats
		{float32(1.2), "1.2"},
		{float64(-1.2), "-1.2"},
		// complex numbers
		{complex64(1.2i + 5), "(5+1.2i)"},
		{complex128(-1.0i + 8), "(8-1i)"},
		// booleans
		{true, "true"},
		{false, "false"},
		// strings
		{"string", "string"},
		{'r', "114"},
		// binary
		{byte(36), "36"},
		{[]byte("hello"), "hello"},
		// composite
		{[]string{"a", "b"}, "[a b]"},
		{struct {
			a string
			b string
		}{"a", "b"}, "{a b}"},
		// references (slices, maps, channels, pointers, functions)
		// interfaces
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValueString(td.value)
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
			value := jsonline.NewValueString(td.value)
			assert.Equal(t, td.expected, value.String())
		})
	}
}

func TestValueFormatNumeric(t *testing.T) {
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
			value := jsonline.NewValueNumeric(td.value)
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
			value := jsonline.NewValueNumeric(td.value)
			assert.Equal(t, td.expected, value.String())
		})
	}
}

func TestValueFormatBoolean(t *testing.T) {
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
			value := jsonline.NewValueBoolean(td.value)
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
		{1, "true"},
		{"1", "true"},
		{"true", "true"},
		{"false", "false"},
		{0, "false"},
		{"0", "false"},
		{1.2, "true"},
		{0.0, "false"},
		{"1.2", "true"},
		{"0.0", "false"},
		{true, "true"},
		{false, "false"},
		{[]byte("true"), "true"},
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValueBoolean(td.value)
			assert.Equal(t, td.expected, value.String())
		})
	}
}

func TestValueFormatBinary(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
	}{
		{nil, nil},
		{1, "MQ=="},
		{1.2, "MS4y"},
		{"string", "c3RyaW5n"},
		{true, "dHJ1ZQ=="},
		{[]byte("hello"), "aGVsbG8="},
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValueBinary(td.value)
			assert.Equal(t, td.expected, value.Export())
		})
	}
}

func TestValueMarshalBinary(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
	}{
		{nil, "null"},
		{1, "\"MQ==\""},
		{1.2, "\"MS4y\""},
		{"string", "\"c3RyaW5n\""},
		{true, "\"dHJ1ZQ==\""},
		{[]byte("hello"), "\"aGVsbG8=\""},
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValueBinary(td.value)
			assert.Equal(t, td.expected, value.String())
		})
	}
}

func TestValueFormatDateTime(t *testing.T) {
	tz, _ := time.LoadLocation("Asia/Shanghai")
	testdatas := []struct {
		value    interface{}
		expected interface{}
	}{
		{nil, nil},
		{"2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z"},
		{"2006-01-02T15:04:05+07:00", "2006-01-02T15:04:05+07:00"},
		{time.Date(2021, time.September, 24, 21, 21, 0, 0, time.UTC), "2021-09-24T21:21:00Z"},
		{time.Date(2021, time.December, 25, 0, 0, 0, 0, tz), "2021-12-25T00:00:00+08:00"},
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValueDateTime(td.value)
			assert.Equal(t, td.expected, value.Export())
		})
	}
}
