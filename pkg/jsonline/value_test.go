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
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/adrienaury/go-template/pkg/jsonline"
	"github.com/stretchr/testify/assert"
)

//nolint:gochecknoglobals
var (
	stringArray [2]string
	intArray    [2]int
)

func TestMain(m *testing.M) {
	stringArray[0] = "a"
	stringArray[1] = "b"
	intArray[0] = 1
	intArray[1] = 2

	os.Exit(m.Run())
}

//nolint:dupl
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
		{int32(1), "\x01"}, // int32 is an alias of rune
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
		{'r', "r"},
		// binary
		{byte(36), "36"}, // byte is an alias for uint8
		{[]byte("hello"), "hello"},
		// composite
		{stringArray, "[a b]"},
		{struct {
			a string
			b string
		}{"a", "b"}, "{a b}"},
		// references (slices, maps)
		{[]string{"a", "b"}, "[a b]"},
		{map[string]string{"k1": "a", "k2": "b"}, "map[k1:a k2:b]"},
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
		{nil, `null`},
		// signed integers
		{int(-2), `"-2"`},
		{int8(-1), `"-1"`},
		{int16(0), `"0"`},
		{int32(1), `"\u0001"`}, // int32 is an alias of rune
		{int64(2), `"2"`},
		// unsigned integers
		{uint(0), `"0"`},
		{uint8(1), `"1"`},
		{uint16(2), `"2"`},
		{uint32(3), `"3"`},
		{uint64(4), `"4"`},
		// floats
		{float32(1.2), `"1.2"`},
		{float64(-1.2), `"-1.2"`},
		// complex numbers
		{complex64(1.2i + 5), `"(5+1.2i)"`},
		{complex128(-1.0i + 8), `"(8-1i)"`},
		// booleans
		{true, `"true"`},
		{false, `"false"`},
		// strings
		{"string", `"string"`},
		{'r', `"r"`},
		// binary
		{byte(36), `"36"`}, // byte is an alias for uint8
		{[]byte("hello"), `"hello"`},
		// composite
		{stringArray, `"[a b]"`},
		{struct {
			a string
			b string
		}{"a", "b"}, `"{a b}"`},
		// references (slices, maps)
		{[]string{"a", "b"}, `"[a b]"`},
		{map[string]string{"k1": "a", "k2": "b"}, `"map[k1:a k2:b]"`},
		// interfaces
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
		// signed integers
		{int(-2), -2}, // either 32 or 64 bits
		{int8(-1), int8(-1)},
		{int16(0), int16(0)},
		{int32(1), int32(1)}, // int32 is an alias of rune
		{int64(2), int64(2)},
		// unsigned integers
		{uint(0), uint(0)}, // either 32 or 64 bits
		{uint8(1), uint8(1)},
		{uint16(2), uint16(2)},
		{uint32(3), uint32(3)},
		{uint64(4), uint64(4)},
		// floats
		{float32(1.2), float32(1.2)},
		{float64(-1.2), float64(-1.2)},
		// complex numbers
		// {complex64(1.2i + 5), "(5+1.2i)"}, => NOT SUPPORTED
		// {complex128(-1.0i + 8), "(8-1i)"}, => NOT SUPPORTED
		// booleans
		{true, 1},
		{false, 0},
		// strings
		{"1.5", 1.5},
		{'r', int32(114)},
		// binary
		{byte(36), byte(36)}, // byte is an alias for uint8
		{[]byte("1.5"), 1.5},
		// composite => NOT SUPPORTED
		// references (slices, maps) => NOT SUPPORTED
		// interfaces => NOT SUPPORTED
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
		{nil, `null`},
		// signed integers
		{int(-2), `-2`}, // either 32 or 64 bits
		{int8(-1), `-1`},
		{int16(0), `0`},
		{int32(1), `1`}, // int32 is an alias of rune
		{int64(2), `2`},
		// unsigned integers
		{uint(0), `0`}, // either 32 or 64 bits
		{uint8(1), `1`},
		{uint16(2), `2`},
		{uint32(3), `3`},
		{uint64(4), `4`},
		// floats
		{float32(1.2), `1.2`},
		{float64(-1.2), `-1.2`},
		// complex numbers
		// {complex64(1.2i + 5), "(5+1.2i)"}, => NOT SUPPORTED
		// {complex128(-1.0i + 8), "(8-1i)"}, => NOT SUPPORTED
		// booleans
		{true, `1`},
		{false, `0`},
		// strings
		{"1.5", `1.5`},
		{'r', `114`},
		// binary
		{byte(36), `36`}, // byte is an alias for uint8
		{[]byte("1.5"), `1.5`},
		// composite => NOT SUPPORTED
		// references (slices, maps) => NOT SUPPORTED
		// interfaces => NOT SUPPORTED
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
		// signed integers
		{int(-2), true},
		{int8(-1), true},
		{int16(0), false},
		{int32(1), true}, // int32 is an alias of rune
		{int64(2), true},
		// unsigned integers
		{uint(0), false},
		{uint8(1), true},
		{uint16(2), true},
		{uint32(3), true},
		{uint64(4), true},
		// floats
		{float32(1.2), true},
		{float64(-1.2), true},
		{float32(0.0), false},
		{float64(0.0), false},
		// complex numbers
		{complex64(1.2i + 5), true},
		{complex64(1.2i), true},
		{complex64(0), false},
		{complex128(-1.0i + 8), true},
		{complex128(-1.0i), true},
		{complex128(0), false},
		// booleans
		{true, true},
		{false, false},
		// strings
		// {"string", true}, => error
		{"1", true},
		{"true", true},
		{"false", false},
		{"0", false},
		{"1.2", true},
		{"0.0", false},
		{'r', true},
		// binary
		{byte(36), true}, // byte is an alias for uint8
		{[]byte("true"), true},
		{[]byte("false"), false},
		// composite => NOT SUPPORTED
		// references (slices, maps) => NOT SUPPORTED
		// interfaces => NOT SUPPORTED
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
		// signed integers
		{int(-2), `true`},
		{int8(-1), `true`},
		{int16(0), `false`},
		{int32(1), `true`}, // int32 is an alias of rune
		{int64(2), `true`},
		// unsigned integers
		{uint(0), `false`},
		{uint8(1), `true`},
		{uint16(2), `true`},
		{uint32(3), `true`},
		{uint64(4), `true`},
		// floats
		{float32(1.2), `true`},
		{float64(-1.2), `true`},
		{float32(0.0), `false`},
		{float64(0.0), `false`},
		// complex numbers
		{complex64(1.2i + 5), `true`},
		{complex64(1.2i), `true`},
		{complex64(0), `false`},
		{complex128(-1.0i + 8), `true`},
		{complex128(-1.0i), `true`},
		{complex128(0), `false`},
		// booleans
		{true, `true`},
		{false, `false`},
		// strings
		// {"string", `true`}, => error
		{"1", `true`},
		{"true", `true`},
		{"false", `false`},
		{"0", `false`},
		{"1.2", `true`},
		{"0.0", `false`},
		{'r', `true`},
		// binary
		{byte(36), `true`}, // byte is an alias for uint8
		{[]byte("true"), `true`},
		{[]byte("false"), `false`},
		// composite => NOT SUPPORTED
		// references (slices, maps) => NOT SUPPORTED
		// interfaces => NOT SUPPORTED
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValueBoolean(td.value)
			assert.Equal(t, td.expected, value.String())
		})
	}
}

//nolint:dupl
func TestValueFormatBinary(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
	}{
		{nil, nil},
		// signed integers
		{int(-2), "LTI="},
		{int8(-1), "LTE="},
		{int16(0), "MA=="},
		{int32(1), "AQ=="}, // int32 is an alias of rune
		{int64(2), "Mg=="},
		// unsigned integers
		{uint(0), "MA=="},
		{uint8(1), "MQ=="},
		{uint16(2), "Mg=="},
		{uint32(3), "Mw=="},
		{uint64(4), "NA=="},
		// floats
		{float32(1.2), "MS4y"},
		{float64(-1.2), "LTEuMg=="},
		// complex numbers
		{complex64(1.2i + 5), "KDUrMS4yaSk="},
		{complex128(-1.0i + 8), "KDgtMWkp"},
		// booleans
		{true, "dHJ1ZQ=="},
		{false, "ZmFsc2U="},
		// strings
		{"string", "c3RyaW5n"},
		{'r', "cg=="},
		// binary
		{byte(36), "MzY="}, // byte is an alias for uint8
		{[]byte("hello"), "aGVsbG8="},
		// composite
		{stringArray, "W2EgYl0="},
		{struct {
			a string
			b string
		}{"a", "b"}, "e2EgYn0="},
		// references (slices, maps)
		{[]string{"a", "b"}, "W2EgYl0="},
		{map[string]string{"k1": "a", "k2": "b"}, "bWFwW2sxOmEgazI6Yl0="},
		// interfaces => NOT SUPPORTED
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
		// signed integers
		{int(-2), `"LTI="`},
		{int8(-1), `"LTE="`},
		{int16(0), `"MA=="`},
		{int32(1), `"AQ=="`}, // int32 is an alias of rune
		{int64(2), `"Mg=="`},
		// unsigned integers
		{uint(0), `"MA=="`},
		{uint8(1), `"MQ=="`},
		{uint16(2), `"Mg=="`},
		{uint32(3), `"Mw=="`},
		{uint64(4), `"NA=="`},
		// floats
		{float32(1.2), `"MS4y"`},
		{float64(-1.2), `"LTEuMg=="`},
		// complex numbers
		{complex64(1.2i + 5), `"KDUrMS4yaSk="`},
		{complex128(-1.0i + 8), `"KDgtMWkp"`},
		// booleans
		{true, `"dHJ1ZQ=="`},
		{false, `"ZmFsc2U="`},
		// strings
		{"string", `"c3RyaW5n"`},
		{'r', `"cg=="`},
		// binary
		{byte(36), `"MzY="`}, // byte is an alias for uint8
		{[]byte("hello"), `"aGVsbG8="`},
		// composite
		{stringArray, `"W2EgYl0="`},
		{struct {
			a string
			b string
		}{"a", "b"}, `"e2EgYn0="`},
		// references (slices, maps)
		{[]string{"a", "b"}, `"W2EgYl0="`},
		{map[string]string{"k1": "a", "k2": "b"}, `"bWFwW2sxOmEgazI6Yl0="`},
		// interfaces => NOT SUPPORTED
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
		// signed integers
		{int(1632823189), "2021-09-28T11:59:49+02:00"},
		{int8(127), "1970-01-01T01:02:07+01:00"},
		{int16(32767), "1970-01-01T10:06:07+01:00"},
		{int32(1632823189), "2021-09-28T11:59:49+02:00"},
		{int64(1632823189), "2021-09-28T11:59:49+02:00"},
		// unsigned integers
		{uint(1632823189), "2021-09-28T11:59:49+02:00"},
		{uint8(255), "1970-01-01T01:04:15+01:00"},
		{uint16(65535), "1970-01-01T19:12:15+01:00"},
		{uint32(1632823189), "2021-09-28T11:59:49+02:00"},
		{uint64(1632823189), "2021-09-28T11:59:49+02:00"},
		// floats
		{float32(1632823189.2), "2021-09-28T11:59:28+02:00"},
		{float64(-1632823189.2), "1918-04-05T15:00:11+01:00"},
		// complex numbers
		// {complex64(1.2i + 5), "(5+1.2i)"}, => UNSUPPORTED
		// {complex128(-1.0i + 8), "(8-1i)"}, => UNSUPPORTED
		// booleans
		// {true, "true"}, => UNSUPPORTED
		// {false, "false"}, => UNSUPPORTED
		// strings
		{"2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z"},
		{"2006-01-02T15:04:05+07:00", "2006-01-02T15:04:05+07:00"},
		// {'r', "r"}, = int32
		// binary
		// {byte(36), "36"}, // = uint8
		{[]byte("2006-01-02T15:04:05+07:00"), "2006-01-02T15:04:05+07:00"},
		// composite => UNSUPPORTED
		// references (structs)
		{time.Date(2021, time.September, 24, 21, 21, 0, 0, time.UTC), "2021-09-24T21:21:00Z"},
		{time.Date(2021, time.December, 25, 0, 0, 0, 0, tz), "2021-12-25T00:00:00+08:00"},
		// interfaces
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValueDateTime(td.value)
			assert.Equal(t, td.expected, value.Export())
		})
	}
}
