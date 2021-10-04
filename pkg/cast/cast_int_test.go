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

//nolint:funlen,dupl,unconvert
package cast_test

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"

	"github.com/cgi-fr/jsonline/pkg/cast"
	"github.com/stretchr/testify/assert"
)

func TestCastToInt(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
		err      error
	}{
		{nil, nil, nil},
		// from int
		{int(math.MinInt), int(math.MinInt), nil},
		{int(0), int(0), nil},
		{int(math.MaxInt), int(math.MaxInt), nil},
		// from int64
		{int64(math.MinInt64), int(math.MinInt64), nil},
		{int64(0), int(0), nil},
		{int64(math.MaxInt64), int(math.MaxInt64), nil},
		// from int32
		{int32(math.MinInt32), int(math.MinInt32), nil},
		{int32(0), int(0), nil},
		{int32(math.MaxInt32), int(math.MaxInt32), nil},
		// from int16
		{int16(math.MinInt16), int(math.MinInt16), nil},
		{int16(0), int(0), nil},
		{int16(math.MaxInt16), int(math.MaxInt16), nil},
		// from int8
		{int8(math.MinInt8), int(math.MinInt8), nil},
		{int8(0), int(0), nil},
		{int8(math.MaxInt8), int(math.MaxInt8), nil},
		// from uint
		{uint(0), int(0), nil},
		{uint(math.MaxUint), nil, cast.ErrUnableToCastToInt}, // expected error
		// from uint64
		{uint64(0), int(0), nil},
		{uint64(math.MaxUint64), nil, cast.ErrUnableToCastToInt}, // expected error
		// from uint32
		{uint32(0), int(0), nil},
		{uint32(math.MaxUint32), int(math.MaxUint32), nil},
		// from iunt16
		{uint16(0), int(0), nil},
		{uint16(math.MaxUint16), int(math.MaxUint16), nil},
		// from uint8
		{uint8(0), int(0), nil},
		{uint8(math.MaxUint8), int(math.MaxUint8), nil},
		// from float64
		{float64(-math.MaxFloat64), nil, cast.ErrUnableToCastToInt}, // expected error
		{float64(-math.SmallestNonzeroFloat64), int(0), nil},
		{float64(0.0), int(0), nil},
		{float64(math.SmallestNonzeroFloat64), int(0), nil},
		{float64(math.MaxFloat64), nil, cast.ErrUnableToCastToInt}, // expected error
		// from float32
		{float32(-math.MaxFloat32), nil, cast.ErrUnableToCastToInt}, // expected error
		{float32(-math.SmallestNonzeroFloat32), int(0), nil},
		{float32(0.0), int(0), nil},
		{float32(math.SmallestNonzeroFloat32), int(0), nil},
		{float32(math.MaxFloat32), nil, cast.ErrUnableToCastToInt}, // expected error
		// from bool
		{bool(true), int(1), nil},
		{bool(false), int(0), nil},
		// from string
		{string("-9223372036854775809"), nil, cast.ErrUnableToCastToInt}, // expected error
		{string("-9223372036854775808"), int(math.MinInt), nil},
		{string("-1"), int(-1), nil},
		{string("0"), int(0), nil},
		{string("1"), int(1), nil},
		{string("9223372036854775807"), int(math.MaxInt), nil},
		{string("9223372036854775808"), nil, cast.ErrUnableToCastToInt},      // expected error
		{string("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt}, // expected error ?
		{string("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt}, // expected error ?
		{string("-1.4"), nil, cast.ErrUnableToCastToInt},                     // expected error ?
		{string("0.0"), nil, cast.ErrUnableToCastToInt},                      // expected error ?
		{string("1.4"), nil, cast.ErrUnableToCastToInt},                      // expected error ?
		{string("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt},  // expected error ?
		{string("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt},  // expected error
		{string("hello"), nil, cast.ErrUnableToCastToInt},                    // expected error
		// from []byte
		{[]byte("-9223372036854775809"), nil, cast.ErrUnableToCastToInt},     // expected error
		{[]byte("-9223372036854775808"), nil, cast.ErrUnableToCastToInt},     // expected error
		{[]byte("-1"), nil, cast.ErrUnableToCastToInt},                       // expected error
		{[]byte("0"), nil, cast.ErrUnableToCastToInt},                        // expected error
		{[]byte("1"), nil, cast.ErrUnableToCastToInt},                        // expected error
		{[]byte("9223372036854775807"), nil, cast.ErrUnableToCastToInt},      // expected error
		{[]byte("9223372036854775808"), nil, cast.ErrUnableToCastToInt},      // expected error
		{[]byte("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt}, // expected error
		{[]byte("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt}, // expected error
		{[]byte("-1.4"), nil, cast.ErrUnableToCastToInt},                     // expected error
		{[]byte("0.0"), nil, cast.ErrUnableToCastToInt},                      // expected error
		{[]byte("1.4"), nil, cast.ErrUnableToCastToInt},                      // expected error
		{[]byte("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt},  // expected error
		{[]byte("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt},  // expected error
		{[]byte("hello"), nil, cast.ErrUnableToCastToInt},                    // expected error
		{[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80}, int(math.MinInt), nil},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, int(-1), nil},
		{[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, int(0), nil},
		{[]byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, int(1), nil},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}, int(math.MaxInt), nil},
		// from number
		{json.Number("9223372036854775807"), int(math.MaxInt), nil},
		{json.Number("9223372036854775808"), nil, cast.ErrUnableToCastToInt},      // expected error
		{json.Number("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt}, // expected error ?
		{json.Number("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt}, // expected error ?
		{json.Number("-1.4"), nil, cast.ErrUnableToCastToInt},                     // expected error ?
		{json.Number("0.0"), nil, cast.ErrUnableToCastToInt},                      // expected error ?
		{json.Number("1.4"), nil, cast.ErrUnableToCastToInt},                      // expected error ?
		{json.Number("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt},  // expected error ?
		{json.Number("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt},  // expected error ?
		{json.Number("hello"), nil, cast.ErrUnableToCastToInt},                    // expected error
		// from anything else
		{struct{ string }{""}, nil, cast.ErrUnableToCastToInt},  // expected error
		{[]int{1}, nil, cast.ErrUnableToCastToInt},              // expected error
		{map[string]int{"": 1}, nil, cast.ErrUnableToCastToInt}, // expected error
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td.value, td.value), func(t *testing.T) {
			result, err := cast.ToInt(td.value)
			assert.ErrorIs(t, err, td.err)
			assert.Equal(t, td.expected, result)
		})
	}
}

func TestCastToInt64(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
		err      error
	}{
		{nil, nil, nil},
		// from int
		{int(math.MinInt), int64(math.MinInt), nil},
		{int(0), int64(0), nil},
		{int(math.MaxInt), int64(math.MaxInt), nil},
		// from int64
		{int64(math.MinInt64), int64(math.MinInt64), nil},
		{int64(0), int64(0), nil},
		{int64(math.MaxInt64), int64(math.MaxInt64), nil},
		// from int32
		{int32(math.MinInt32), int64(math.MinInt32), nil},
		{int32(0), int64(0), nil},
		{int32(math.MaxInt32), int64(math.MaxInt32), nil},
		// from int16
		{int16(math.MinInt16), int64(math.MinInt16), nil},
		{int16(0), int64(0), nil},
		{int16(math.MaxInt16), int64(math.MaxInt16), nil},
		// from int8
		{int8(math.MinInt8), int64(math.MinInt8), nil},
		{int8(0), int64(0), nil},
		{int8(math.MaxInt8), int64(math.MaxInt8), nil},
		// from uint
		{uint(0), int64(0), nil},
		{uint(math.MaxUint), nil, cast.ErrUnableToCastToInt64}, // expected error
		// from uint64
		{uint64(0), int64(0), nil},
		{uint64(math.MaxUint64), nil, cast.ErrUnableToCastToInt64}, // expected error
		// from uint32
		{uint32(0), int64(0), nil},
		{uint32(math.MaxUint32), int64(math.MaxUint32), nil},
		// from iunt16
		{uint16(0), int64(0), nil},
		{uint16(math.MaxUint16), int64(math.MaxUint16), nil},
		// from uint8
		{uint8(0), int64(0), nil},
		{uint8(math.MaxUint8), int64(math.MaxUint8), nil},
		// from float64
		{float64(-math.MaxFloat64), nil, cast.ErrUnableToCastToInt64}, // expected error
		{float64(-math.SmallestNonzeroFloat64), int64(0), nil},
		{float64(0.0), int64(0), nil},
		{float64(math.SmallestNonzeroFloat64), int64(0), nil},
		{float64(math.MaxFloat64), nil, cast.ErrUnableToCastToInt64}, // expected error
		// from float32
		{float32(-math.MaxFloat32), nil, cast.ErrUnableToCastToInt64}, // expected error
		{float32(-math.SmallestNonzeroFloat32), int64(0), nil},
		{float32(0.0), int64(0), nil},
		{float32(math.SmallestNonzeroFloat32), int64(0), nil},
		{float32(math.MaxFloat32), nil, cast.ErrUnableToCastToInt64}, // expected error
		// from bool
		{bool(true), int64(1), nil},
		{bool(false), int64(0), nil},
		// from string
		{string("-9223372036854775809"), nil, cast.ErrUnableToCastToInt64}, // expected error
		{string("-9223372036854775808"), int64(math.MinInt64), nil},
		{string("-1"), int64(-1), nil},
		{string("0"), int64(0), nil},
		{string("1"), int64(1), nil},
		{string("9223372036854775807"), int64(math.MaxInt64), nil},
		{string("9223372036854775808"), nil, cast.ErrUnableToCastToInt64},      // expected error
		{string("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt64}, // expected error ?
		{string("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt64}, // expected error ?
		{string("-1.4"), nil, cast.ErrUnableToCastToInt64},                     // expected error ?
		{string("0.0"), nil, cast.ErrUnableToCastToInt64},                      // expected error ?
		{string("1.4"), nil, cast.ErrUnableToCastToInt64},                      // expected error ?
		{string("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt64},  // expected error ?
		{string("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt64},  // expected error
		{string("hello"), nil, cast.ErrUnableToCastToInt64},                    // expected error
		// from []byte
		{[]byte("-9223372036854775809"), nil, cast.ErrUnableToCastToInt64},     // expected error
		{[]byte("-9223372036854775808"), nil, cast.ErrUnableToCastToInt64},     // expected error
		{[]byte("-1"), nil, cast.ErrUnableToCastToInt64},                       // expected error
		{[]byte("0"), nil, cast.ErrUnableToCastToInt64},                        // expected error
		{[]byte("1"), nil, cast.ErrUnableToCastToInt64},                        // expected error
		{[]byte("9223372036854775807"), nil, cast.ErrUnableToCastToInt64},      // expected error
		{[]byte("9223372036854775808"), nil, cast.ErrUnableToCastToInt64},      // expected error
		{[]byte("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt64}, // expected error
		{[]byte("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt64}, // expected error
		{[]byte("-1.4"), nil, cast.ErrUnableToCastToInt64},                     // expected error
		{[]byte("0.0"), nil, cast.ErrUnableToCastToInt64},                      // expected error
		{[]byte("1.4"), nil, cast.ErrUnableToCastToInt64},                      // expected error
		{[]byte("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt64},  // expected error
		{[]byte("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt64},  // expected error
		{[]byte("hello"), nil, cast.ErrUnableToCastToInt64},                    // expected error
		{[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80}, int64(math.MinInt64), nil},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, int64(-1), nil},
		{[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, int64(0), nil},
		{[]byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, int64(1), nil},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}, int64(math.MaxInt64), nil},
		// from number
		{json.Number("9223372036854775807"), int64(math.MaxInt64), nil},
		{json.Number("9223372036854775808"), nil, cast.ErrUnableToCastToInt64},      // expected error
		{json.Number("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt64}, // expected error ?
		{json.Number("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt64}, // expected error ?
		{json.Number("-1.4"), nil, cast.ErrUnableToCastToInt64},                     // expected error ?
		{json.Number("0.0"), nil, cast.ErrUnableToCastToInt64},                      // expected error ?
		{json.Number("1.4"), nil, cast.ErrUnableToCastToInt64},                      // expected error ?
		{json.Number("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt64},  // expected error ?
		{json.Number("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt64},  // expected error ?
		{json.Number("hello"), nil, cast.ErrUnableToCastToInt64},                    // expected error
		// from anything else
		{struct{ string }{""}, nil, cast.ErrUnableToCastToInt64},  // expected error
		{[]int{1}, nil, cast.ErrUnableToCastToInt64},              // expected error
		{map[string]int{"": 1}, nil, cast.ErrUnableToCastToInt64}, // expected error
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td.value, td.value), func(t *testing.T) {
			result, err := cast.ToInt64(td.value)
			assert.ErrorIs(t, err, td.err)
			assert.Equal(t, td.expected, result)
		})
	}
}

func TestCastToInt32(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
		err      error
	}{
		{nil, nil, nil},
		// from int
		{int(math.MinInt), nil, cast.ErrUnableToCastToInt32}, // expected error
		{int(0), int32(0), nil},
		{int(math.MaxInt), nil, cast.ErrUnableToCastToInt32}, // expected error
		// from int64
		{int64(math.MinInt64), nil, cast.ErrUnableToCastToInt32}, // expected error
		{int64(0), int32(0), nil},
		{int64(math.MaxInt64), nil, cast.ErrUnableToCastToInt32}, // expected error
		// from int32
		{int32(math.MinInt32), int32(math.MinInt32), nil},
		{int32(0), int32(0), nil},
		{int32(math.MaxInt32), int32(math.MaxInt32), nil},
		// from int16
		{int16(math.MinInt16), int32(math.MinInt16), nil},
		{int16(0), int32(0), nil},
		{int16(math.MaxInt16), int32(math.MaxInt16), nil},
		// from int8
		{int8(math.MinInt8), int32(math.MinInt8), nil},
		{int8(0), int32(0), nil},
		{int8(math.MaxInt8), int32(math.MaxInt8), nil},
		// from uint
		{uint(0), int32(0), nil},
		{uint(math.MaxUint), nil, cast.ErrUnableToCastToInt32}, // expected error
		// from uint64
		{uint64(0), int32(0), nil},
		{uint64(math.MaxUint64), nil, cast.ErrUnableToCastToInt32}, // expected error
		// from uint32
		{uint32(0), int32(0), nil},
		{uint32(math.MaxUint32), nil, cast.ErrUnableToCastToInt32}, // expected error
		// from iunt16
		{uint16(0), int32(0), nil},
		{uint16(math.MaxUint16), int32(math.MaxUint16), nil},
		// from uint8
		{uint8(0), int32(0), nil},
		{uint8(math.MaxUint8), int32(math.MaxUint8), nil},
		// from float64
		{float64(-math.MaxFloat64), nil, cast.ErrUnableToCastToInt32}, // expected error
		{float64(-math.SmallestNonzeroFloat64), int32(0), nil},
		{float64(0.0), int32(0), nil},
		{float64(math.SmallestNonzeroFloat64), int32(0), nil},
		{float64(math.MaxFloat64), nil, cast.ErrUnableToCastToInt32}, // expected error
		// from float32
		{float32(-math.MaxFloat32), nil, cast.ErrUnableToCastToInt32}, // expected error
		{float32(-math.SmallestNonzeroFloat32), int32(0), nil},
		{float32(0.0), int32(0), nil},
		{float32(math.SmallestNonzeroFloat32), int32(0), nil},
		{float32(math.MaxFloat32), nil, cast.ErrUnableToCastToInt32}, // expected error
		// from bool
		{bool(true), int32(1), nil},
		{bool(false), int32(0), nil},
		// from string
		{string("-9223372036854775809"), nil, cast.ErrUnableToCastToInt32}, // expected error
		{string("-9223372036854775808"), nil, cast.ErrUnableToCastToInt32}, // expected error
		{string("-1"), int32(-1), nil},
		{string("0"), int32(0), nil},
		{string("1"), int32(1), nil},
		{string(fmt.Sprintf("%v", math.MaxInt32)), int32(math.MaxInt32), nil},
		{string("9223372036854775808"), nil, cast.ErrUnableToCastToInt32},      // expected error
		{string("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt32}, // expected error ?
		{string("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt32}, // expected error ?
		{string("-1.4"), nil, cast.ErrUnableToCastToInt32},                     // expected error ?
		{string("0.0"), nil, cast.ErrUnableToCastToInt32},                      // expected error ?
		{string("1.4"), nil, cast.ErrUnableToCastToInt32},                      // expected error ?
		{string("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt32},  // expected error ?
		{string("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt32},  // expected error
		{string("hello"), nil, cast.ErrUnableToCastToInt32},                    // expected error
		// from []byte
		{[]byte("-9223372036854775809"), nil, cast.ErrUnableToCastToInt32},     // expected error
		{[]byte("-9223372036854775808"), nil, cast.ErrUnableToCastToInt32},     // expected error
		{[]byte("-1"), nil, cast.ErrUnableToCastToInt32},                       // expected error
		{[]byte("0"), nil, cast.ErrUnableToCastToInt32},                        // expected error
		{[]byte("1"), nil, cast.ErrUnableToCastToInt32},                        // expected error
		{[]byte("9223372036854775807"), nil, cast.ErrUnableToCastToInt32},      // expected error
		{[]byte("9223372036854775808"), nil, cast.ErrUnableToCastToInt32},      // expected error
		{[]byte("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt32}, // expected error
		{[]byte("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt32}, // expected error
		{[]byte("-1.4"), int32(875442477), nil},                                // technically it's on 4 bytes
		{[]byte("0.0"), nil, cast.ErrUnableToCastToInt32},                      // expected error
		{[]byte("1.4"), nil, cast.ErrUnableToCastToInt32},                      // expected error
		{[]byte("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt32},  // expected error
		{[]byte("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt32},  // expected error
		{[]byte("hello"), nil, cast.ErrUnableToCastToInt32},                    // expected error
		{[]byte{0x0, 0x0, 0x0, 0x80}, int32(math.MinInt32), nil},
		{[]byte{0xff, 0xff, 0xff, 0xff}, int32(-1), nil},
		{[]byte{0x0, 0x0, 0x0, 0x0}, int32(0), nil},
		{[]byte{0x1, 0x0, 0x0, 0x0}, int32(1), nil},
		{[]byte{0xff, 0xff, 0xff, 0x7f}, int32(math.MaxInt32), nil},
		// from number
		{json.Number(fmt.Sprintf("%v", math.MaxInt32)), int32(math.MaxInt32), nil},
		{json.Number("9223372036854775808"), nil, cast.ErrUnableToCastToInt32},      // expected error
		{json.Number("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt32}, // expected error ?
		{json.Number("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt32}, // expected error ?
		{json.Number("-1.4"), nil, cast.ErrUnableToCastToInt32},                     // expected error ?
		{json.Number("0.0"), nil, cast.ErrUnableToCastToInt32},                      // expected error ?
		{json.Number("1.4"), nil, cast.ErrUnableToCastToInt32},                      // expected error ?
		{json.Number("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt32},  // expected error ?
		{json.Number("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt32},  // expected error ?
		{json.Number("hello"), nil, cast.ErrUnableToCastToInt32},                    // expected error
		// from anything else
		{struct{ string }{""}, nil, cast.ErrUnableToCastToInt32},  // expected error
		{[]int{1}, nil, cast.ErrUnableToCastToInt32},              // expected error
		{map[string]int{"": 1}, nil, cast.ErrUnableToCastToInt32}, // expected error
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td.value, td.value), func(t *testing.T) {
			result, err := cast.ToInt32(td.value)
			assert.ErrorIs(t, err, td.err)
			assert.Equal(t, td.expected, result)
		})
	}
}

func TestCastToInt16(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
		err      error
	}{
		{nil, nil, nil},
		// from int
		{int(math.MinInt), nil, cast.ErrUnableToCastToInt16}, // expected error
		{int(0), int16(0), nil},
		{int(math.MaxInt), nil, cast.ErrUnableToCastToInt16}, // expected error
		// from int64
		{int64(math.MinInt64), nil, cast.ErrUnableToCastToInt16}, // expected error
		{int64(0), int16(0), nil},
		{int64(math.MaxInt64), nil, cast.ErrUnableToCastToInt16}, // expected error
		// from int32
		{int32(math.MinInt32), nil, cast.ErrUnableToCastToInt16}, // expected error
		{int32(0), int16(0), nil},
		{int32(math.MaxInt32), nil, cast.ErrUnableToCastToInt16}, // expected error
		// from int16
		{int16(math.MinInt16), int16(math.MinInt16), nil},
		{int16(0), int16(0), nil},
		{int16(math.MaxInt16), int16(math.MaxInt16), nil},
		// from int8
		{int8(math.MinInt8), int16(math.MinInt8), nil},
		{int8(0), int16(0), nil},
		{int8(math.MaxInt8), int16(math.MaxInt8), nil},
		// from uint
		{uint(0), int16(0), nil},
		{uint(math.MaxUint), nil, cast.ErrUnableToCastToInt16}, // expected error
		// from uint64
		{uint64(0), int16(0), nil},
		{uint64(math.MaxUint64), nil, cast.ErrUnableToCastToInt16}, // expected error
		// from uint32
		{uint32(0), int16(0), nil},
		{uint32(math.MaxUint32), nil, cast.ErrUnableToCastToInt16}, // expected error
		// from uint16
		{uint16(0), int16(0), nil},
		{uint16(math.MaxUint16), nil, cast.ErrUnableToCastToInt16}, // expected error
		// from uint8
		{uint8(0), int16(0), nil},
		{uint8(math.MaxUint8), int16(math.MaxUint8), nil},
		// from float64
		{float64(-math.MaxFloat64), nil, cast.ErrUnableToCastToInt16}, // expected error
		{float64(-math.SmallestNonzeroFloat64), int16(0), nil},
		{float64(0.0), int16(0), nil},
		{float64(math.SmallestNonzeroFloat64), int16(0), nil},
		{float64(math.MaxFloat64), nil, cast.ErrUnableToCastToInt16}, // expected error
		// from float32
		{float32(-math.MaxFloat32), nil, cast.ErrUnableToCastToInt16}, // expected error
		{float32(-math.SmallestNonzeroFloat32), int16(0), nil},
		{float32(0.0), int16(0), nil},
		{float32(math.SmallestNonzeroFloat32), int16(0), nil},
		{float32(math.MaxFloat32), nil, cast.ErrUnableToCastToInt16}, // expected error
		// from bool
		{bool(true), int16(1), nil},
		{bool(false), int16(0), nil},
		// from string
		{string("-9223372036854775809"), nil, cast.ErrUnableToCastToInt16}, // expected error
		{string("-9223372036854775808"), nil, cast.ErrUnableToCastToInt16}, // expected error
		{string("-1"), int16(-1), nil},
		{string("0"), int16(0), nil},
		{string("1"), int16(1), nil},
		{string(fmt.Sprintf("%v", math.MaxInt16)), int16(math.MaxInt16), nil},
		{string("9223372036854775808"), nil, cast.ErrUnableToCastToInt16},      // expected error
		{string("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt16}, // expected error ?
		{string("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt16}, // expected error ?
		{string("-1.4"), nil, cast.ErrUnableToCastToInt16},                     // expected error ?
		{string("0.0"), nil, cast.ErrUnableToCastToInt16},                      // expected error ?
		{string("1.4"), nil, cast.ErrUnableToCastToInt16},                      // expected error ?
		{string("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt16},  // expected error ?
		{string("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt16},  // expected error
		{string("hello"), nil, cast.ErrUnableToCastToInt16},                    // expected error
		// from []byte
		{[]byte("-9223372036854775809"), nil, cast.ErrUnableToCastToInt16},     // expected error
		{[]byte("-9223372036854775808"), nil, cast.ErrUnableToCastToInt16},     // expected error
		{[]byte("-1"), int16(12589), nil},                                      // technically it's on 2 bytes
		{[]byte("0"), nil, cast.ErrUnableToCastToInt16},                        // expected error
		{[]byte("1"), nil, cast.ErrUnableToCastToInt16},                        // expected error
		{[]byte("9223372036854775807"), nil, cast.ErrUnableToCastToInt16},      // expected error
		{[]byte("9223372036854775808"), nil, cast.ErrUnableToCastToInt16},      // expected error
		{[]byte("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt16}, // expected error
		{[]byte("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt16}, // expected error
		{[]byte("-1.4"), nil, cast.ErrUnableToCastToInt16},                     // expected error
		{[]byte("0.0"), nil, cast.ErrUnableToCastToInt16},                      // expected error
		{[]byte("1.4"), nil, cast.ErrUnableToCastToInt16},                      // expected error
		{[]byte("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt16},  // expected error
		{[]byte("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt16},  // expected error
		{[]byte("hello"), nil, cast.ErrUnableToCastToInt16},                    // expected error
		{[]byte{0x0, 0x80}, int16(math.MinInt16), nil},
		{[]byte{0xff, 0xff}, int16(-1), nil},
		{[]byte{0x0, 0x0}, int16(0), nil},
		{[]byte{0x1, 0x0}, int16(1), nil},
		{[]byte{0xff, 0x7f}, int16(math.MaxInt16), nil},
		// from number
		{json.Number(fmt.Sprintf("%v", math.MaxInt16)), int16(math.MaxInt16), nil},
		{json.Number("9223372036854775808"), nil, cast.ErrUnableToCastToInt16},      // expected error
		{json.Number("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt16}, // expected error ?
		{json.Number("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt16}, // expected error ?
		{json.Number("-1.4"), nil, cast.ErrUnableToCastToInt16},                     // expected error ?
		{json.Number("0.0"), nil, cast.ErrUnableToCastToInt16},                      // expected error ?
		{json.Number("1.4"), nil, cast.ErrUnableToCastToInt16},                      // expected error ?
		{json.Number("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt16},  // expected error ?
		{json.Number("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt16},  // expected error ?
		{json.Number("hello"), nil, cast.ErrUnableToCastToInt16},                    // expected error
		// from anything else
		{struct{ string }{""}, nil, cast.ErrUnableToCastToInt16},  // expected error
		{[]int{1}, nil, cast.ErrUnableToCastToInt16},              // expected error
		{map[string]int{"": 1}, nil, cast.ErrUnableToCastToInt16}, // expected error
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td.value, td.value), func(t *testing.T) {
			result, err := cast.ToInt16(td.value)
			assert.ErrorIs(t, err, td.err)
			assert.Equal(t, td.expected, result)
		})
	}
}

func TestCastToInt8(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
		err      error
	}{
		{nil, nil, nil},
		// from int
		{int(math.MinInt), nil, cast.ErrUnableToCastToInt8}, // expected error
		{int(0), int8(0), nil},
		{int(math.MaxInt), nil, cast.ErrUnableToCastToInt8}, // expected error
		// from int64
		{int64(math.MinInt64), nil, cast.ErrUnableToCastToInt8}, // expected error
		{int64(0), int8(0), nil},
		{int64(math.MaxInt64), nil, cast.ErrUnableToCastToInt8}, // expected error
		// from int32
		{int32(math.MinInt32), nil, cast.ErrUnableToCastToInt8}, // expected error
		{int32(0), int8(0), nil},
		{int32(math.MaxInt32), nil, cast.ErrUnableToCastToInt8}, // expected error
		// from int16
		{int16(math.MinInt16), nil, cast.ErrUnableToCastToInt8}, // expected error
		{int16(0), int8(0), nil},
		{int16(math.MaxInt16), nil, cast.ErrUnableToCastToInt8}, // expected error
		// from int8
		{int8(math.MinInt8), int8(math.MinInt8), nil},
		{int8(0), int8(0), nil},
		{int8(math.MaxInt8), int8(math.MaxInt8), nil},
		// from uint
		{uint(0), int8(0), nil},
		{uint(math.MaxUint), nil, cast.ErrUnableToCastToInt8}, // expected error
		// from uint64
		{uint64(0), int8(0), nil},
		{uint64(math.MaxUint64), nil, cast.ErrUnableToCastToInt8}, // expected error
		// from uint32
		{uint32(0), int8(0), nil},
		{uint32(math.MaxUint32), nil, cast.ErrUnableToCastToInt8}, // expected error
		// from uint16
		{uint16(0), int8(0), nil},
		{uint16(math.MaxUint16), nil, cast.ErrUnableToCastToInt8}, // expected error
		// from uint8
		{uint8(0), int8(0), nil},
		{uint8(math.MaxUint8), nil, cast.ErrUnableToCastToInt8}, // expected error
		// from float64
		{float64(-math.MaxFloat64), nil, cast.ErrUnableToCastToInt8}, // expected error
		{float64(-math.SmallestNonzeroFloat64), int8(0), nil},
		{float64(0.0), int8(0), nil},
		{float64(math.SmallestNonzeroFloat64), int8(0), nil},
		{float64(math.MaxFloat64), nil, cast.ErrUnableToCastToInt8}, // expected error
		// from float32
		{float32(-math.MaxFloat32), nil, cast.ErrUnableToCastToInt8}, // expected error
		{float32(-math.SmallestNonzeroFloat32), int8(0), nil},
		{float32(0.0), int8(0), nil},
		{float32(math.SmallestNonzeroFloat32), int8(0), nil},
		{float32(math.MaxFloat32), nil, cast.ErrUnableToCastToInt8}, // expected error
		// from bool
		{bool(true), int8(1), nil},
		{bool(false), int8(0), nil},
		// from string
		{string("-9223372036854775809"), nil, cast.ErrUnableToCastToInt8}, // expected error
		{string("-9223372036854775808"), nil, cast.ErrUnableToCastToInt8}, // expected error
		{string("-1"), int8(-1), nil},
		{string("0"), int8(0), nil},
		{string("1"), int8(1), nil},
		{string(fmt.Sprintf("%v", math.MaxInt8)), int8(math.MaxInt8), nil},
		{string("9223372036854775808"), nil, cast.ErrUnableToCastToInt8},      // expected error
		{string("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt8}, // expected error ?
		{string("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt8}, // expected error ?
		{string("-1.4"), nil, cast.ErrUnableToCastToInt8},                     // expected error ?
		{string("0.0"), nil, cast.ErrUnableToCastToInt8},                      // expected error ?
		{string("1.4"), nil, cast.ErrUnableToCastToInt8},                      // expected error ?
		{string("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt8},  // expected error ?
		{string("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt8},  // expected error
		{string("hello"), nil, cast.ErrUnableToCastToInt8},                    // expected error
		// from []byte
		{[]byte("-9223372036854775809"), nil, cast.ErrUnableToCastToInt8},     // expected error
		{[]byte("-9223372036854775808"), nil, cast.ErrUnableToCastToInt8},     // expected error
		{[]byte("-1"), nil, cast.ErrUnableToCastToInt8},                       // expected error
		{[]byte("0"), int8(48), nil},                                          // technically it's 1 byte
		{[]byte("1"), int8(49), nil},                                          // technically it's 1 byte
		{[]byte("9223372036854775807"), nil, cast.ErrUnableToCastToInt8},      // expected error
		{[]byte("9223372036854775808"), nil, cast.ErrUnableToCastToInt8},      // expected error
		{[]byte("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt8}, // expected error
		{[]byte("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt8}, // expected error
		{[]byte("-1.4"), nil, cast.ErrUnableToCastToInt8},                     // expected error
		{[]byte("0.0"), nil, cast.ErrUnableToCastToInt8},                      // expected error
		{[]byte("1.4"), nil, cast.ErrUnableToCastToInt8},                      // expected error
		{[]byte("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt8},  // expected error
		{[]byte("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt8},  // expected error
		{[]byte("hello"), nil, cast.ErrUnableToCastToInt8},                    // expected error
		{[]byte{0x80}, int8(math.MinInt8), nil},
		{[]byte{0xff}, int8(-1), nil},
		{[]byte{0x0}, int8(0), nil},
		{[]byte{0x1}, int8(1), nil},
		{[]byte{0x7f}, int8(math.MaxInt8), nil},
		// from number
		{json.Number(fmt.Sprintf("%v", math.MaxInt8)), int8(math.MaxInt8), nil},
		{json.Number("9223372036854775808"), nil, cast.ErrUnableToCastToInt8},      // expected error
		{json.Number("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt8}, // expected error ?
		{json.Number("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt8}, // expected error ?
		{json.Number("-1.4"), nil, cast.ErrUnableToCastToInt8},                     // expected error ?
		{json.Number("0.0"), nil, cast.ErrUnableToCastToInt8},                      // expected error ?
		{json.Number("1.4"), nil, cast.ErrUnableToCastToInt8},                      // expected error ?
		{json.Number("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToInt8},  // expected error ?
		{json.Number("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToInt8},  // expected error ?
		{json.Number("hello"), nil, cast.ErrUnableToCastToInt8},                    // expected error
		// from anything else
		{struct{ string }{""}, nil, cast.ErrUnableToCastToInt8},  // expected error
		{[]int{1}, nil, cast.ErrUnableToCastToInt8},              // expected error
		{map[string]int{"": 1}, nil, cast.ErrUnableToCastToInt8}, // expected error
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td.value, td.value), func(t *testing.T) {
			result, err := cast.ToInt8(td.value)
			assert.ErrorIs(t, err, td.err)
			assert.Equal(t, td.expected, result)
		})
	}
}

func BenchmarkCastToIntZero(b *testing.B) { castToInt(b, 0) }

func BenchmarkCastToIntMaxInt64(b *testing.B) { castToInt(b, int64(math.MaxInt64)) }

func BenchmarkCastToIntMaxInt32(b *testing.B) { castToInt(b, int32(math.MaxInt32)) }

func BenchmarkCastToIntMaxInt16(b *testing.B) { castToInt(b, int16(math.MaxInt16)) }

func BenchmarkCastToIntMaxInt8(b *testing.B) { castToInt(b, int8(math.MaxInt8)) }

func BenchmarkCastToIntMinInt64(b *testing.B) { castToInt(b, int64(math.MinInt64)) }

func BenchmarkCastToIntMinInt32(b *testing.B) { castToInt(b, int32(math.MinInt32)) }

func BenchmarkCastToIntMinInt16(b *testing.B) { castToInt(b, int16(math.MinInt16)) }

func BenchmarkCastToIntMinInt8(b *testing.B) { castToInt(b, int8(math.MinInt8)) }

func BenchmarkCastToIntMaxFloat64(b *testing.B) { castToInt(b, float64(math.MaxFloat64)) }

func BenchmarkCastToIntMaxFloat32(b *testing.B) { castToInt(b, float32(math.MaxFloat32)) }

func BenchmarkCastToIntSmallestFloat64(b *testing.B) {
	castToInt(b, float64(math.SmallestNonzeroFloat64))
}

func BenchmarkCastToIntSmallestFloat32(b *testing.B) {
	castToInt(b, float32(math.SmallestNonzeroFloat32))
}

func BenchmarkCastToIntByte(b *testing.B) {
	castToInt(b, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
}

func BenchmarkCastToIntByteOverflow(b *testing.B) {
	castToInt(b, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
}

func castToInt(b *testing.B, i interface{}) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		cast.ToInt(i) //nolint:errcheck
	}
}

func BenchmarkCastToInt8Zero(b *testing.B) { castToInt8(b, 0) }

func BenchmarkCastToInt8MaxInt64(b *testing.B) { castToInt8(b, int64(math.MaxInt64)) }

func BenchmarkCastToInt8MaxInt32(b *testing.B) { castToInt8(b, int32(math.MaxInt32)) }

func BenchmarkCastToInt8MaxInt16(b *testing.B) { castToInt8(b, int16(math.MaxInt16)) }

func BenchmarkCastToInt8MaxInt8(b *testing.B) { castToInt8(b, int8(math.MaxInt8)) }

func BenchmarkCastToInt8MinInt64(b *testing.B) { castToInt8(b, int64(math.MinInt64)) }

func BenchmarkCastToInt8MinInt32(b *testing.B) { castToInt8(b, int32(math.MinInt32)) }

func BenchmarkCastToInt8MinInt16(b *testing.B) { castToInt8(b, int16(math.MinInt16)) }

func BenchmarkCastToInt8MinInt8(b *testing.B) { castToInt8(b, int8(math.MinInt8)) }

func BenchmarkCastToInt8MaxFloat64(b *testing.B) { castToInt8(b, float64(math.MaxFloat64)) }

func BenchmarkCastToInt8MaxFloat32(b *testing.B) { castToInt8(b, float32(math.MaxFloat32)) }

func BenchmarkCastToInt8SmallestFloat64(b *testing.B) {
	castToInt8(b, float64(math.SmallestNonzeroFloat64))
}

func BenchmarkCastToInt8SmallestFloat32(b *testing.B) {
	castToInt8(b, float32(math.SmallestNonzeroFloat32))
}

func BenchmarkCastToInt8Byte(b *testing.B) {
	castToInt8(b, []byte{0xff})
}

func BenchmarkCastToInt8ByteOverflow(b *testing.B) {
	castToInt8(b, []byte{0xff, 0xff})
}

func castToInt8(b *testing.B, i interface{}) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		cast.ToInt8(i) //nolint:errcheck
	}
}
