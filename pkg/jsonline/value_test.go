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

//nolint:dupl
package jsonline_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cgi-fr/jsonline/pkg/jsonline"
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
		{int32(1), "1"}, // int32 is an alias of rune
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
		// complex numbers => NOT SUPPORTED YET
		// {complex64(1.2i + 5), "(5+1.2i)"},
		// {complex128(-1.0i + 8), "(8-1i)"},
		// booleans
		{true, "true"},
		{false, "false"},
		// strings
		{"string", "string"},
		{'r', "114"}, // rune is an alias of int32
		// binary
		{byte(36), "36"}, // byte is an alias for uint8
		{[]byte("hello"), "hello"},
		// composite => NOT SUPPORTED YET
		// {stringArray, "[a b]"},
		// {struct {
		// 	a string
		// 	b string
		// }{"a", "b"}, "{a b}"},
		// references (slices, maps) => NOT SUPPORTED YET
		// {[]string{"a", "b"}, "[a b]"},
		// {map[string]string{"k1": "a", "k2": "b"}, "map[k1:a k2:b]"},
		// interfaces
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValueString(td.value)
			result, err := value.Export()
			assert.NoError(t, err)
			assert.Equal(t, td.expected, result)
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
		{int32(1), `"1"`}, // int32 is an alias of rune
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
		// complex numbers => NOT SUPPORTED YET
		// {complex64(1.2i + 5), `"(5+1.2i)"`},
		// {complex128(-1.0i + 8), `"(8-1i)"`},
		// booleans
		{true, `"true"`},
		{false, `"false"`},
		// strings
		{"string", `"string"`},
		{'r', `"114"`}, // rune is an alias of int32
		// binary
		{byte(36), `"36"`}, // byte is an alias for uint8
		{[]byte("hello"), `"hello"`},
		// composite => NOT SUPPORTED YET
		// {stringArray, `"[a b]"`},
		// {struct {
		// 	a string
		// 	b string
		// }{"a", "b"}, `"{a b}"`},
		// references (slices, maps) => NOT SUPPORTED YET
		// {[]string{"a", "b"}, `"[a b]"`},
		// {map[string]string{"k1": "a", "k2": "b"}, `"map[k1:a k2:b]"`},
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
		{int(-2), json.Number("-2")}, // either 32 or 64 bits
		{int8(-1), json.Number("-1")},
		{int16(0), json.Number("0")},
		{int32(1), json.Number("1")}, // int32 is an alias of rune
		{int64(2), json.Number("2")},
		// unsigned integers
		{uint(0), json.Number("0")}, // either 32 or 64 bits
		{uint8(1), json.Number("1")},
		{uint16(2), json.Number("2")},
		{uint32(3), json.Number("3")},
		{uint64(4), json.Number("4")},
		// floats
		{float32(1.2), json.Number("1.2")},
		{float64(-1.2), json.Number("-1.2")},
		// complex numbers
		// {complex64(1.2i + 5), "(5+1.2i)"}, => NOT SUPPORTED
		// {complex128(-1.0i + 8), "(8-1i)"}, => NOT SUPPORTED
		// booleans
		{true, json.Number("1")},
		{false, json.Number("0")},
		// strings
		{"1.5", json.Number("1.5")},
		{'r', json.Number("114")},
		// binary
		{byte(36), json.Number("36")}, // byte is an alias for uint8
		{[]byte("1.5"), json.Number("1.5")},
		// composite => NOT SUPPORTED
		// references (slices, maps) => NOT SUPPORTED
		// interfaces => NOT SUPPORTED
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValueNumeric(td.value)
			result, err := value.Export()
			assert.NoError(t, err)
			assert.Equal(t, td.expected, result)
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
		// complex numbers => NOT SUPPORTED YET
		// {complex64(1.2i + 5), true},
		// {complex64(1.2i), true},
		// {complex64(0), false},
		// {complex128(-1.0i + 8), true},
		// {complex128(-1.0i), true},
		// {complex128(0), false},
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
		// {[]byte("true"), true}, => NOT SUPPORTED
		// {[]byte("false"), false}, => NOT SUPPORTED
		// composite => NOT SUPPORTED
		// references (slices, maps) => NOT SUPPORTED
		// interfaces => NOT SUPPORTED
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValueBoolean(td.value)
			result, err := value.Export()
			assert.NoError(t, err)
			assert.Equal(t, td.expected, result)
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
		// complex numbers => NOT SUPPORTED YET
		// {complex64(1.2i + 5), `true`},
		// {complex64(1.2i), `true`},
		// {complex64(0), `false`},
		// {complex128(-1.0i + 8), `true`},
		// {complex128(-1.0i), `true`},
		// {complex128(0), `false`},
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
		// {[]byte("true"), `true`},
		// {[]byte("false"), `false`},
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

func TestValueExportBinary(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
	}{
		{nil, nil},
		// signed integers
		{int(-2), "/v////////8="},
		{int8(-1), "/w=="},
		{int16(0), "AAA="},
		{int32(1), "AQAAAA=="}, // int32 is an alias of rune
		{int64(2), "AgAAAAAAAAA="},
		// unsigned integers
		{uint(0), "AAAAAAAAAAA="},
		{uint8(1), "AQ=="},
		{uint16(2), "AgA="},
		{uint32(3), "AwAAAA=="},
		{uint64(4), "BAAAAAAAAAA="},
		// floats
		{float32(1.2), "mpmZPw=="},
		{float64(-1.2), "MzMzMzMz878="},
		// complex numbers => NOT SUPPORTED YET
		// {complex64(1.2i + 5), "KDUrMS4yaSk="},
		// {complex128(-1.0i + 8), "KDgtMWkp"},
		// booleans
		{true, "AQ=="},
		{false, "AA=="},
		// strings
		{"string", "c3RyaW5n"},
		{'r', "cgAAAA=="},
		// binary
		{byte(36), "JA=="}, // byte is an alias for uint8
		{[]byte("hello"), "aGVsbG8="},
		// composite
		// {stringArray, "W2EgYl0="},
		// {struct {
		// 	a string
		// 	b string
		// }{"a", "b"}, "e2EgYn0="},
		// references (slices, maps)
		// {[]string{"a", "b"}, "W2EgYl0="},
		// {map[string]string{"k1": "a", "k2": "b"}, "bWFwW2sxOmEgazI6Yl0="},
		// interfaces => NOT SUPPORTED
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%#v)", td.value, td.value), func(t *testing.T) {
			value := jsonline.NewValueBinary(td.value)
			result, err := value.Export()
			assert.NoError(t, err)
			assert.Equal(t, td.expected, result)
		})
	}
}

func TestValueImportBinary(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
	}{
		{nil, nil},
		{"LTI=", []byte{0x2d, 0x32}},
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValueBinary(nil)
			err := value.Import(td.value)
			assert.NoError(t, err)
			assert.Equal(t, td.expected, value.Raw())
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
		{int(-2), `"/v////////8="`},
		{int8(-1), `"/w=="`},
		{int16(0), `"AAA="`},
		{int32(1), `"AQAAAA=="`}, // int32 is an alias of rune
		{int64(2), `"AgAAAAAAAAA="`},
		// unsigned integers
		{uint(0), `"AAAAAAAAAAA="`},
		{uint8(1), `"AQ=="`},
		{uint16(2), `"AgA="`},
		{uint32(3), `"AwAAAA=="`},
		{uint64(4), `"BAAAAAAAAAA="`},
		// floats
		{float32(1.2), `"mpmZPw=="`},
		{float64(-1.2), `"MzMzMzMz878="`},
		// complex numbers => NOT SUPPORTED YET
		// {complex64(1.2i + 5), `"KDUrMS4yaSk="`},
		// {complex128(-1.0i + 8), `"KDgtMWkp"`},
		// booleans
		{true, `"AQ=="`},
		{false, `"AA=="`},
		// strings
		{"string", `"c3RyaW5n"`},
		{'r', `"cgAAAA=="`},
		// binary
		{byte(36), `"JA=="`}, // byte is an alias for uint8
		{[]byte("hello"), `"aGVsbG8="`},
		// composite => NOT SUPPORTED YET
		// {stringArray, `"W2EgYl0="`},
		// {struct {
		// 	a string
		// 	b string
		// }{"a", "b"}, `"e2EgYn0="`},
		// references (slices, maps) => NOT SUPPORTED YET
		// {[]string{"a", "b"}, `"W2EgYl0="`},
		// {map[string]string{"k1": "a", "k2": "b"}, `"bWFwW2sxOmEgazI6Yl0="`},
		// interfaces => NOT SUPPORTED
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValueBinary(td.value)
			assert.Equal(t, td.expected, value.String())
		})
	}
}

func TestValueExportDateTime(t *testing.T) {
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
			result, err := value.Export()
			assert.NoError(t, err)
			assert.Equal(t, td.expected, result)
		})
	}
}

func TestValueImportDateTime(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
	}{
		{nil, nil},
		{"2021-09-28T11:59:49+02:00", time.Date(2021, time.September, 28, 11, 59, 49, 0, time.Local)},
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValueDateTime(nil)
			err := value.Import(td.value)
			assert.NoError(t, err)
			assert.Equal(t, td.expected, value.Raw())
		})
	}
}

func TestValueExportTimestamp(t *testing.T) {
	tz, _ := time.LoadLocation("Asia/Shanghai")
	testdatas := []struct {
		value    interface{}
		expected interface{}
	}{
		{nil, nil},
		// signed integers
		{int(1632823189), int64(1632823189)},
		{int8(127), int64(127)},
		{int16(32767), int64(32767)},
		{int32(1632823189), int64(1632823189)},
		{int64(1632823189), int64(1632823189)},
		// unsigned integers
		{uint(1632823189), int64(1632823189)},
		{uint8(255), int64(255)},
		{uint16(65535), int64(65535)},
		{uint32(1632823189), int64(1632823189)},
		{uint64(1632823189), int64(1632823189)},
		// floats
		{float32(1632823189.2), int64(1632823168)}, // rounding error
		{float64(-1632823189.2), int64(-1632823189)},
		// complex numbers
		// {complex64(1.2i + 5), "(5+1.2i)"}, => UNSUPPORTED
		// {complex128(-1.0i + 8), "(8-1i)"}, => UNSUPPORTED
		// booleans
		// {true, "true"}, => UNSUPPORTED
		// {false, "false"}, => UNSUPPORTED
		// strings
		{"2006-01-02T15:04:05Z", int64(1136214245)},
		{"2006-01-02T15:04:05+07:00", int64(1136189045)},
		// {'r', "r"}, = int32
		// binary
		// {byte(36), "36"}, // = uint8
		{[]byte("2006-01-02T15:04:05+07:00"), int64(1136189045)},
		// composite => UNSUPPORTED
		// references (structs)
		{time.Date(2021, time.September, 24, 21, 21, 0, 0, time.UTC), int64(1632518460)},
		{time.Date(2021, time.December, 25, 0, 0, 0, 0, tz), int64(1640361600)},
		// interfaces
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValueTimestamp(td.value)
			result, err := value.Export()
			assert.NoError(t, err)
			assert.Equal(t, td.expected, result)
		})
	}
}

func TestValueImportTimestamp(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
	}{
		{nil, nil},
		{1136189045, int64(1136189045)},
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%#v", td.value), func(t *testing.T) {
			value := jsonline.NewValueTimestamp(nil)
			err := value.Import(td.value)
			assert.NoError(t, err)
			assert.Equal(t, td.expected, value.Raw())
		})
	}
}

func TestValueFormatAuto(t *testing.T) {
	testdatas := []interface{}{
		nil,
		// signed integers
		int(-2),
		int8(-1),
		int16(0),
		int32(1), // int32 is an alias of
		int64(2),
		// unsigned int
		uint(0),
		uint8(1),
		uint16(2),
		uint32(3),
		uint64(4),
		// floats
		float32(1.2),
		float64(-1.2),
		// complex numbers
		complex64(1.2i + 5),
		complex128(-1.0i + 8),
		// booleans
		true,
		false,
		// strings
		"string",
		'r',
		// binary
		byte(36),
		[]byte("hello"),
		// composite
		stringArray,
		// references
		[]string{"a", "b"},
		map[string]string{"k1": "a", "k2": "b"},
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("Auto %#v", td), func(t *testing.T) {
			value := jsonline.NewValueAuto(td)
			result, err := value.Export()
			assert.NoError(t, err)
			assert.Equal(t, td, result)
		})
		t.Run(fmt.Sprintf("Hidden %#v", td), func(t *testing.T) {
			value := jsonline.NewValueHidden(td)
			result, err := value.Export()
			assert.NoError(t, err)
			assert.Equal(t, td, result)
		})
	}
}
