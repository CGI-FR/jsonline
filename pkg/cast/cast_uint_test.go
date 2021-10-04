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

func TestCastToUint(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
		err      error
	}{
		{nil, nil, nil},
		// from int
		{int(math.MinInt), nil, cast.ErrUnableToCastToUint}, // expected eror
		{int(0), uint(0), nil},
		{int(math.MaxInt), uint(math.MaxInt), nil},
		// from int64
		{int64(math.MinInt64), nil, cast.ErrUnableToCastToUint}, // expected eror
		{int64(0), uint(0), nil},
		{int64(math.MaxInt64), uint(math.MaxInt64), nil},
		// from int32
		{int32(math.MinInt32), nil, cast.ErrUnableToCastToUint}, // expected eror
		{int32(0), uint(0), nil},
		{int32(math.MaxInt32), uint(math.MaxInt32), nil},
		// from int16
		{int16(math.MinInt16), nil, cast.ErrUnableToCastToUint}, // expected eror
		{int16(0), uint(0), nil},
		{int16(math.MaxInt16), uint(math.MaxInt16), nil},
		// from int8
		{int8(math.MinInt8), nil, cast.ErrUnableToCastToUint}, // expected eror
		{int8(0), uint(0), nil},
		{int8(math.MaxInt8), uint(math.MaxInt8), nil},
		// from uint
		{uint(0), uint(0), nil},
		{uint(math.MaxUint), uint(math.MaxUint), nil},
		// from uint64
		{uint64(0), uint(0), nil},
		{uint64(math.MaxUint64), uint(math.MaxUint64), nil},
		// from uint32
		{uint32(0), uint(0), nil},
		{uint32(math.MaxUint32), uint(math.MaxUint32), nil},
		// from iunt16
		{uint16(0), uint(0), nil},
		{uint16(math.MaxUint16), uint(math.MaxUint16), nil},
		// from uint8
		{uint8(0), uint(0), nil},
		{uint8(math.MaxUint8), uint(math.MaxUint8), nil},
		// from float64
		{float64(-math.MaxFloat64), nil, cast.ErrUnableToCastToUint},             // expected eror
		{float64(-math.SmallestNonzeroFloat64), nil, cast.ErrUnableToCastToUint}, // expected eror
		{float64(0.0), uint(0), nil},
		{float64(math.SmallestNonzeroFloat64), uint(0), nil},
		{float64(math.MaxFloat64), nil, cast.ErrUnableToCastToUint}, // expected eror
		// from float32
		{float32(-math.MaxFloat32), nil, cast.ErrUnableToCastToUint},             // expected eror
		{float32(-math.SmallestNonzeroFloat32), nil, cast.ErrUnableToCastToUint}, // expected eror
		{float32(0.0), uint(0), nil},
		{float32(math.SmallestNonzeroFloat32), uint(0), nil},
		{float32(math.MaxFloat32), nil, cast.ErrUnableToCastToUint}, // expected eror
		// from bool
		{bool(true), uint(1), nil},
		{bool(false), uint(0), nil},
		// from string
		{string("-9223372036854775809"), nil, cast.ErrUnableToCastToUint}, // expected eror
		{string("-9223372036854775808"), nil, cast.ErrUnableToCastToUint}, // expected eror
		{string("-1"), nil, cast.ErrUnableToCastToUint},                   // expected eror
		{string("0"), uint(0), nil},
		{string("1"), uint(1), nil},
		{string("18446744073709551615"), uint(math.MaxUint), nil},
		{string("18446744073709551616"), nil, cast.ErrUnableToCastToUint},     // expected eror
		{string("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint}, // expected eror
		{string("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint}, // expected eror
		{string("-1.4"), nil, cast.ErrUnableToCastToUint},                     // expected eror
		{string("0.0"), nil, cast.ErrUnableToCastToUint},                      // expected eror ?
		{string("1.4"), nil, cast.ErrUnableToCastToUint},                      // expected eror ?
		{string("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint},  // expected eror ?
		{string("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint},  // expected eror
		{string("hello"), nil, cast.ErrUnableToCastToUint},                    // expected eror
		// from []byte
		{[]byte("-9223372036854775809"), nil, cast.ErrUnableToCastToUint},     // expected eror
		{[]byte("-9223372036854775808"), nil, cast.ErrUnableToCastToUint},     // expected eror
		{[]byte("-1"), nil, cast.ErrUnableToCastToUint},                       // expected eror
		{[]byte("0"), nil, cast.ErrUnableToCastToUint},                        // expected eror
		{[]byte("1"), nil, cast.ErrUnableToCastToUint},                        // expected eror
		{[]byte("9223372036854775807"), nil, cast.ErrUnableToCastToUint},      // expected eror
		{[]byte("9223372036854775808"), nil, cast.ErrUnableToCastToUint},      // expected eror
		{[]byte("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint}, // expected eror
		{[]byte("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint}, // expected eror
		{[]byte("-1.4"), nil, cast.ErrUnableToCastToUint},                     // expected eror
		{[]byte("0.0"), nil, cast.ErrUnableToCastToUint},                      // expected eror
		{[]byte("1.4"), nil, cast.ErrUnableToCastToUint},                      // expected eror
		{[]byte("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint},  // expected eror
		{[]byte("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint},  // expected eror
		{[]byte("hello"), nil, cast.ErrUnableToCastToUint},                    // expected eror
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, uint(math.MaxUint), nil},
		{[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, uint(0), nil},
		{[]byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, uint(1), nil},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}, uint(math.MaxInt), nil},
		// from number
		{json.Number("18446744073709551615"), uint(math.MaxUint), nil},
		{json.Number("18446744073709551616"), nil, cast.ErrUnableToCastToUint},     // expected eror
		{json.Number("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint}, // expected eror
		{json.Number("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint}, // expected eror
		{json.Number("-1.4"), nil, cast.ErrUnableToCastToUint},                     // expected eror
		{json.Number("0.0"), nil, cast.ErrUnableToCastToUint},                      // expected eror ?
		{json.Number("1.4"), nil, cast.ErrUnableToCastToUint},                      // expected eror ?
		{json.Number("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint},  // expected eror ?
		{json.Number("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint},  // expected eror ?
		{json.Number("hello"), nil, cast.ErrUnableToCastToUint},                    // expected eror
		// from anything else
		{struct{ string }{""}, nil, cast.ErrUnableToCastToUint},  // expected eror
		{[]int{1}, nil, cast.ErrUnableToCastToUint},              // expected eror
		{map[string]int{"": 1}, nil, cast.ErrUnableToCastToUint}, // expected eror
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td.value, td.value), func(t *testing.T) {
			result, err := cast.ToUint(td.value)
			assert.ErrorIs(t, err, td.err)
			assert.Equal(t, td.expected, result)
		})
	}
}

func TestCastToUint64(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
		err      error
	}{
		{nil, nil, nil},
		// from int
		{int(math.MinInt), nil, cast.ErrUnableToCastToUint64}, // expected eror
		{int(0), uint64(0), nil},
		{int(math.MaxInt), uint64(math.MaxInt), nil},
		// from int64
		{int64(math.MinInt64), nil, cast.ErrUnableToCastToUint64}, // expected eror
		{int64(0), uint64(0), nil},
		{int64(math.MaxInt64), uint64(math.MaxInt64), nil},
		// from int32
		{int32(math.MinInt32), nil, cast.ErrUnableToCastToUint64}, // expected eror
		{int32(0), uint64(0), nil},
		{int32(math.MaxInt32), uint64(math.MaxInt32), nil},
		// from int16
		{int16(math.MinInt16), nil, cast.ErrUnableToCastToUint64}, // expected eror
		{int16(0), uint64(0), nil},
		{int16(math.MaxInt16), uint64(math.MaxInt16), nil},
		// from int8
		{int8(math.MinInt8), nil, cast.ErrUnableToCastToUint64}, // expected eror
		{int8(0), uint64(0), nil},
		{int8(math.MaxInt8), uint64(math.MaxInt8), nil},
		// from uint
		{uint(0), uint64(0), nil},
		{uint(math.MaxUint), uint64(math.MaxUint), nil},
		// from uint64
		{uint64(0), uint64(0), nil},
		{uint64(math.MaxUint64), uint64(math.MaxUint64), nil},
		// from uint32
		{uint32(0), uint64(0), nil},
		{uint32(math.MaxUint32), uint64(math.MaxUint32), nil},
		// from iunt16
		{uint16(0), uint64(0), nil},
		{uint16(math.MaxUint16), uint64(math.MaxUint16), nil},
		// from uint8
		{uint8(0), uint64(0), nil},
		{uint8(math.MaxUint8), uint64(math.MaxUint8), nil},
		// from float64
		{float64(-math.MaxFloat64), nil, cast.ErrUnableToCastToUint64},             // expected eror
		{float64(-math.SmallestNonzeroFloat64), nil, cast.ErrUnableToCastToUint64}, // expected eror
		{float64(0.0), uint64(0), nil},
		{float64(math.SmallestNonzeroFloat64), uint64(0), nil},
		{float64(math.MaxFloat64), nil, cast.ErrUnableToCastToUint64}, // expected eror
		// from float32
		{float32(-math.MaxFloat32), nil, cast.ErrUnableToCastToUint64},             // expected eror
		{float32(-math.SmallestNonzeroFloat32), nil, cast.ErrUnableToCastToUint64}, // expected eror
		{float32(0.0), uint64(0), nil},
		{float32(math.SmallestNonzeroFloat32), uint64(0), nil},
		{float32(math.MaxFloat32), nil, cast.ErrUnableToCastToUint64}, // expected eror
		// from bool
		{bool(true), uint64(1), nil},
		{bool(false), uint64(0), nil},
		// from string
		{string("-9223372036854775809"), nil, cast.ErrUnableToCastToUint64}, // expected eror
		{string("-9223372036854775808"), nil, cast.ErrUnableToCastToUint64}, // expected eror
		{string("-1"), nil, cast.ErrUnableToCastToUint64},                   // expected eror
		{string("0"), uint64(0), nil},
		{string("1"), uint64(1), nil},
		{string("18446744073709551615"), uint64(math.MaxUint64), nil},
		{string("18446744073709551616"), nil, cast.ErrUnableToCastToUint64},     // expected eror
		{string("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint64}, // expected eror
		{string("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint64}, // expected eror
		{string("-1.4"), nil, cast.ErrUnableToCastToUint64},                     // expected eror
		{string("0.0"), nil, cast.ErrUnableToCastToUint64},                      // expected eror ?
		{string("1.4"), nil, cast.ErrUnableToCastToUint64},                      // expected eror ?
		{string("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint64},  // expected eror ?
		{string("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint64},  // expected eror
		{string("hello"), nil, cast.ErrUnableToCastToUint64},                    // expected eror
		// from []byte
		{[]byte("-9223372036854775809"), nil, cast.ErrUnableToCastToUint64},     // expected eror
		{[]byte("-9223372036854775808"), nil, cast.ErrUnableToCastToUint64},     // expected eror
		{[]byte("-1"), nil, cast.ErrUnableToCastToUint64},                       // expected eror
		{[]byte("0"), nil, cast.ErrUnableToCastToUint64},                        // expected eror
		{[]byte("1"), nil, cast.ErrUnableToCastToUint64},                        // expected eror
		{[]byte("9223372036854775807"), nil, cast.ErrUnableToCastToUint64},      // expected eror
		{[]byte("9223372036854775808"), nil, cast.ErrUnableToCastToUint64},      // expected eror
		{[]byte("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint64}, // expected eror
		{[]byte("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint64}, // expected eror
		{[]byte("-1.4"), nil, cast.ErrUnableToCastToUint64},                     // expected eror
		{[]byte("0.0"), nil, cast.ErrUnableToCastToUint64},                      // expected eror
		{[]byte("1.4"), nil, cast.ErrUnableToCastToUint64},                      // expected eror
		{[]byte("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint64},  // expected eror
		{[]byte("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint64},  // expected eror
		{[]byte("hello"), nil, cast.ErrUnableToCastToUint64},                    // expected eror
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, uint64(math.MaxUint64), nil},
		{[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, uint64(0), nil},
		{[]byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, uint64(1), nil},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}, uint64(math.MaxInt64), nil},
		// from number
		{json.Number("18446744073709551615"), uint64(math.MaxUint64), nil},
		{json.Number("18446744073709551616"), nil, cast.ErrUnableToCastToUint64},     // expected eror
		{json.Number("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint64}, // expected eror
		{json.Number("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint64}, // expected eror
		{json.Number("-1.4"), nil, cast.ErrUnableToCastToUint64},                     // expected eror
		{json.Number("0.0"), nil, cast.ErrUnableToCastToUint64},                      // expected eror ?
		{json.Number("1.4"), nil, cast.ErrUnableToCastToUint64},                      // expected eror ?
		{json.Number("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint64},  // expected eror ?
		{json.Number("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint64},  // expected eror ?
		{json.Number("hello"), nil, cast.ErrUnableToCastToUint64},                    // expected eror
		// from anything else
		{struct{ string }{""}, nil, cast.ErrUnableToCastToUint64},  // expected eror
		{[]int{1}, nil, cast.ErrUnableToCastToUint64},              // expected eror
		{map[string]int{"": 1}, nil, cast.ErrUnableToCastToUint64}, // expected eror
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td.value, td.value), func(t *testing.T) {
			result, err := cast.ToUint64(td.value)
			assert.ErrorIs(t, err, td.err)
			assert.Equal(t, td.expected, result)
		})
	}
}

func TestCastToUint32(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
		err      error
	}{
		{nil, nil, nil},
		// from int
		{int(math.MinInt), nil, cast.ErrUnableToCastToUint32}, // expected eror
		{int(0), uint32(0), nil},
		{int(math.MaxInt), nil, cast.ErrUnableToCastToUint32}, // expected eror
		// from int64
		{int64(math.MinInt64), nil, cast.ErrUnableToCastToUint32}, // expected eror
		{int64(0), uint32(0), nil},
		{int64(math.MaxInt64), nil, cast.ErrUnableToCastToUint32}, // expected eror
		// from int32
		{int32(math.MinInt32), nil, cast.ErrUnableToCastToUint32}, // expected eror
		{int32(0), uint32(0), nil},
		{int32(math.MaxInt32), uint32(math.MaxInt32), nil},
		// from int16
		{int16(math.MinInt16), nil, cast.ErrUnableToCastToUint32}, // expected eror
		{int16(0), uint32(0), nil},
		{int16(math.MaxInt16), uint32(math.MaxInt16), nil},
		// from int8
		{int8(math.MinInt8), nil, cast.ErrUnableToCastToUint32}, // expected eror
		{int8(0), uint32(0), nil},
		{int8(math.MaxInt8), uint32(math.MaxInt8), nil},
		// from uint
		{uint(0), uint32(0), nil},
		{uint(math.MaxUint), nil, cast.ErrUnableToCastToUint32}, // expected eror
		// from uint64
		{uint64(0), uint32(0), nil},
		{uint64(math.MaxUint64), nil, cast.ErrUnableToCastToUint32}, // expected eror
		// from uint32
		{uint32(0), uint32(0), nil},
		{uint32(math.MaxUint32), uint32(math.MaxUint32), nil},
		// from iunt16
		{uint16(0), uint32(0), nil},
		{uint16(math.MaxUint16), uint32(math.MaxUint16), nil},
		// from uint8
		{uint8(0), uint32(0), nil},
		{uint8(math.MaxUint8), uint32(math.MaxUint8), nil},
		// from float64
		{float64(-math.MaxFloat64), nil, cast.ErrUnableToCastToUint32},             // expected eror
		{float64(-math.SmallestNonzeroFloat64), nil, cast.ErrUnableToCastToUint32}, // expected eror
		{float64(0.0), uint32(0), nil},
		{float64(math.SmallestNonzeroFloat64), uint32(0), nil},
		{float64(math.MaxFloat64), nil, cast.ErrUnableToCastToUint32}, // expected eror
		// from float32
		{float32(-math.MaxFloat32), nil, cast.ErrUnableToCastToUint32},             // expected eror
		{float32(-math.SmallestNonzeroFloat32), nil, cast.ErrUnableToCastToUint32}, // expected eror
		{float32(0.0), uint32(0), nil},
		{float32(math.SmallestNonzeroFloat32), uint32(0), nil},
		{float32(math.MaxFloat32), nil, cast.ErrUnableToCastToUint32}, // expected eror
		// from bool
		{bool(true), uint32(1), nil},
		{bool(false), uint32(0), nil},
		// from string
		{string("-9223372036854775809"), nil, cast.ErrUnableToCastToUint32}, // expected eror
		{string("-9223372036854775808"), nil, cast.ErrUnableToCastToUint32}, // expected eror
		{string("-1"), nil, cast.ErrUnableToCastToUint32},                   // expected eror
		{string("0"), uint32(0), nil},
		{string("1"), uint32(1), nil},
		{string(fmt.Sprintf("%v", math.MaxInt32)), uint32(math.MaxInt32), nil},
		{string("9223372036854775808"), nil, cast.ErrUnableToCastToUint32},      // expected eror
		{string("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint32}, // expected eror ?
		{string("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint32}, // expected eror ?
		{string("-1.4"), nil, cast.ErrUnableToCastToUint32},                     // expected eror ?
		{string("0.0"), nil, cast.ErrUnableToCastToUint32},                      // expected eror ?
		{string("1.4"), nil, cast.ErrUnableToCastToUint32},                      // expected eror ?
		{string("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint32},  // expected eror ?
		{string("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint32},  // expected eror
		{string("hello"), nil, cast.ErrUnableToCastToUint32},                    // expected eror
		// from []byte
		{[]byte("-9223372036854775809"), nil, cast.ErrUnableToCastToUint32},     // expected eror
		{[]byte("-9223372036854775808"), nil, cast.ErrUnableToCastToUint32},     // expected eror
		{[]byte("-1"), nil, cast.ErrUnableToCastToUint32},                       // expected eror
		{[]byte("0"), nil, cast.ErrUnableToCastToUint32},                        // expected eror
		{[]byte("1"), nil, cast.ErrUnableToCastToUint32},                        // expected eror
		{[]byte("9223372036854775807"), nil, cast.ErrUnableToCastToUint32},      // expected eror
		{[]byte("9223372036854775808"), nil, cast.ErrUnableToCastToUint32},      // expected eror
		{[]byte("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint32}, // expected eror
		{[]byte("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint32}, // expected eror
		{[]byte("-1.4"), uint32(875442477), nil},                                // technically it's on 4 bytes
		{[]byte("0.0"), nil, cast.ErrUnableToCastToUint32},                      // expected eror
		{[]byte("1.4"), nil, cast.ErrUnableToCastToUint32},                      // expected eror
		{[]byte("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint32},  // expected eror
		{[]byte("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint32},  // expected eror
		{[]byte("hello"), nil, cast.ErrUnableToCastToUint32},                    // expected eror
		{[]byte{0xff, 0xff, 0xff, 0xff}, uint32(math.MaxUint32), nil},
		{[]byte{0x0, 0x0, 0x0, 0x0}, uint32(0), nil},
		{[]byte{0x1, 0x0, 0x0, 0x0}, uint32(1), nil},
		{[]byte{0xff, 0xff, 0xff, 0x7f}, uint32(math.MaxInt32), nil},
		// from number
		{json.Number(fmt.Sprintf("%v", math.MaxInt32)), uint32(math.MaxInt32), nil},
		{json.Number("9223372036854775808"), nil, cast.ErrUnableToCastToUint32},      // expected eror
		{json.Number("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint32}, // expected eror ?
		{json.Number("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint32}, // expected eror ?
		{json.Number("-1.4"), nil, cast.ErrUnableToCastToUint32},                     // expected eror ?
		{json.Number("0.0"), nil, cast.ErrUnableToCastToUint32},                      // expected eror ?
		{json.Number("1.4"), nil, cast.ErrUnableToCastToUint32},                      // expected eror ?
		{json.Number("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint32},  // expected eror ?
		{json.Number("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint32},  // expected eror ?
		{json.Number("hello"), nil, cast.ErrUnableToCastToUint32},                    // expected eror
		// from anything else
		{struct{ string }{""}, nil, cast.ErrUnableToCastToUint32},  // expected eror
		{[]int{1}, nil, cast.ErrUnableToCastToUint32},              // expected eror
		{map[string]int{"": 1}, nil, cast.ErrUnableToCastToUint32}, // expected eror
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td.value, td.value), func(t *testing.T) {
			result, err := cast.ToUint32(td.value)
			assert.ErrorIs(t, err, td.err)
			assert.Equal(t, td.expected, result)
		})
	}
}

func TestCastToUint16(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
		err      error
	}{
		{nil, nil, nil},
		// from int
		{int(math.MinInt), nil, cast.ErrUnableToCastToUint16}, // expected eror
		{int(0), uint16(0), nil},
		{int(math.MaxInt), nil, cast.ErrUnableToCastToUint16}, // expected eror
		// from int64
		{int64(math.MinInt64), nil, cast.ErrUnableToCastToUint16}, // expected eror
		{int64(0), uint16(0), nil},
		{int64(math.MaxInt64), nil, cast.ErrUnableToCastToUint16}, // expected eror
		// from int32
		{int32(math.MinInt32), nil, cast.ErrUnableToCastToUint16}, // expected eror
		{int32(0), uint16(0), nil},
		{int32(math.MaxInt32), nil, cast.ErrUnableToCastToUint16}, // expected eror
		// from int16
		{int16(math.MinInt16), nil, cast.ErrUnableToCastToUint16}, // expected eror
		{int16(0), uint16(0), nil},
		{int16(math.MaxInt16), uint16(math.MaxInt16), nil},
		// from int8
		{int8(math.MinInt8), nil, cast.ErrUnableToCastToUint16}, // expected eror
		{int8(0), uint16(0), nil},
		{int8(math.MaxInt8), uint16(math.MaxInt8), nil},
		// from uint
		{uint(0), uint16(0), nil},
		{uint(math.MaxUint), nil, cast.ErrUnableToCastToUint16}, // expected eror
		// from uint64
		{uint64(0), uint16(0), nil},
		{uint64(math.MaxUint64), nil, cast.ErrUnableToCastToUint16}, // expected eror
		// from uint32
		{uint32(0), uint16(0), nil},
		{uint32(math.MaxUint32), nil, cast.ErrUnableToCastToUint16}, // expected eror
		// from uint16
		{uint16(0), uint16(0), nil},
		{uint16(math.MaxUint16), uint16(math.MaxUint16), nil},
		// from uint8
		{uint8(0), uint16(0), nil},
		{uint8(math.MaxUint8), uint16(math.MaxUint8), nil},
		// from float64
		{float64(-math.MaxFloat64), nil, cast.ErrUnableToCastToUint16},             // expected eror
		{float64(-math.SmallestNonzeroFloat64), nil, cast.ErrUnableToCastToUint16}, // expected eror
		{float64(0.0), uint16(0), nil},
		{float64(math.SmallestNonzeroFloat64), uint16(0), nil},
		{float64(math.MaxFloat64), nil, cast.ErrUnableToCastToUint16}, // expected eror
		// from float32
		{float32(-math.MaxFloat32), nil, cast.ErrUnableToCastToUint16},             // expected eror
		{float32(-math.SmallestNonzeroFloat32), nil, cast.ErrUnableToCastToUint16}, // expected eror
		{float32(0.0), uint16(0), nil},
		{float32(math.SmallestNonzeroFloat32), uint16(0), nil},
		{float32(math.MaxFloat32), nil, cast.ErrUnableToCastToUint16}, // expected eror
		// from bool
		{bool(true), uint16(1), nil},
		{bool(false), uint16(0), nil},
		// from string
		{string("-9223372036854775809"), nil, cast.ErrUnableToCastToUint16}, // expected eror
		{string("-9223372036854775808"), nil, cast.ErrUnableToCastToUint16}, // expected eror
		{string("-1"), nil, cast.ErrUnableToCastToUint16},                   // expected eror
		{string("0"), uint16(0), nil},
		{string("1"), uint16(1), nil},
		{string(fmt.Sprintf("%v", math.MaxInt16)), uint16(math.MaxInt16), nil},
		{string("9223372036854775808"), nil, cast.ErrUnableToCastToUint16},      // expected eror
		{string("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint16}, // expected eror
		{string("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint16}, // expected eror
		{string("-1.4"), nil, cast.ErrUnableToCastToUint16},                     // expected eror
		{string("0.0"), nil, cast.ErrUnableToCastToUint16},                      // expected eror ?
		{string("1.4"), nil, cast.ErrUnableToCastToUint16},                      // expected eror ?
		{string("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint16},  // expected eror ?
		{string("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint16},  // expected eror
		{string("hello"), nil, cast.ErrUnableToCastToUint16},                    // expected eror
		// from []byte
		{[]byte("-9223372036854775809"), nil, cast.ErrUnableToCastToUint16},     // expected eror
		{[]byte("-9223372036854775808"), nil, cast.ErrUnableToCastToUint16},     // expected eror
		{[]byte("-1"), uint16(12589), nil},                                      // technically it's on 2 bytes
		{[]byte("0"), nil, cast.ErrUnableToCastToUint16},                        // expected eror
		{[]byte("1"), nil, cast.ErrUnableToCastToUint16},                        // expected eror
		{[]byte("9223372036854775807"), nil, cast.ErrUnableToCastToUint16},      // expected eror
		{[]byte("9223372036854775808"), nil, cast.ErrUnableToCastToUint16},      // expected eror
		{[]byte("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint16}, // expected eror
		{[]byte("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint16}, // expected eror
		{[]byte("-1.4"), nil, cast.ErrUnableToCastToUint16},                     // expected eror
		{[]byte("0.0"), nil, cast.ErrUnableToCastToUint16},                      // expected eror
		{[]byte("1.4"), nil, cast.ErrUnableToCastToUint16},                      // expected eror
		{[]byte("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint16},  // expected eror
		{[]byte("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint16},  // expected eror
		{[]byte("hello"), nil, cast.ErrUnableToCastToUint16},                    // expected eror
		{[]byte{0xff, 0xff}, uint16(math.MaxUint16), nil},
		{[]byte{0x0, 0x0}, uint16(0), nil},
		{[]byte{0x1, 0x0}, uint16(1), nil},
		{[]byte{0xff, 0x7f}, uint16(math.MaxInt16), nil},
		// from number
		{json.Number(fmt.Sprintf("%v", math.MaxInt16)), uint16(math.MaxInt16), nil},
		{json.Number("9223372036854775808"), nil, cast.ErrUnableToCastToUint16},      // expected eror
		{json.Number("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint16}, // expected eror
		{json.Number("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint16}, // expected eror
		{json.Number("-1.4"), nil, cast.ErrUnableToCastToUint16},                     // expected eror
		{json.Number("0.0"), nil, cast.ErrUnableToCastToUint16},                      // expected eror ?
		{json.Number("1.4"), nil, cast.ErrUnableToCastToUint16},                      // expected eror ?
		{json.Number("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint16},  // expected eror ?
		{json.Number("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint16},  // expected eror ?
		{json.Number("hello"), nil, cast.ErrUnableToCastToUint16},                    // expected eror
		// from anything else
		{struct{ string }{""}, nil, cast.ErrUnableToCastToUint16},  // expected eror
		{[]int{1}, nil, cast.ErrUnableToCastToUint16},              // expected eror
		{map[string]int{"": 1}, nil, cast.ErrUnableToCastToUint16}, // expected eror
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td.value, td.value), func(t *testing.T) {
			result, err := cast.ToUint16(td.value)
			assert.ErrorIs(t, err, td.err)
			assert.Equal(t, td.expected, result)
		})
	}
}

func TestCastToUint8(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
		err      error
	}{
		{nil, nil, nil},
		// from int
		{int(math.MinInt), nil, cast.ErrUnableToCastToUint8}, // expected eror
		{int(0), uint8(0), nil},
		{int(math.MaxInt), nil, cast.ErrUnableToCastToUint8}, // expected eror
		// from int64
		{int64(math.MinInt64), nil, cast.ErrUnableToCastToUint8}, // expected eror
		{int64(0), uint8(0), nil},
		{int64(math.MaxInt64), nil, cast.ErrUnableToCastToUint8}, // expected eror
		// from int32
		{int32(math.MinInt32), nil, cast.ErrUnableToCastToUint8}, // expected eror
		{int32(0), uint8(0), nil},
		{int32(math.MaxInt32), nil, cast.ErrUnableToCastToUint8}, // expected eror
		// from int16
		{int16(math.MinInt16), nil, cast.ErrUnableToCastToUint8}, // expected eror
		{int16(0), uint8(0), nil},
		{int16(math.MaxInt16), nil, cast.ErrUnableToCastToUint8}, // expected eror
		// from int8
		{int8(math.MinInt8), nil, cast.ErrUnableToCastToUint8}, // expected eror
		{int8(0), uint8(0), nil},
		{int8(math.MaxInt8), uint8(math.MaxInt8), nil},
		// from uint
		{uint(0), uint8(0), nil},
		{uint(math.MaxUint), nil, cast.ErrUnableToCastToUint8}, // expected eror
		// from uint64
		{uint64(0), uint8(0), nil},
		{uint64(math.MaxUint64), nil, cast.ErrUnableToCastToUint8}, // expected eror
		// from uint32
		{uint32(0), uint8(0), nil},
		{uint32(math.MaxUint32), nil, cast.ErrUnableToCastToUint8}, // expected eror
		// from uint16
		{uint16(0), uint8(0), nil},
		{uint16(math.MaxUint16), nil, cast.ErrUnableToCastToUint8}, // expected eror
		// from uint8
		{uint8(0), uint8(0), nil},
		{uint8(math.MaxUint8), uint8(math.MaxUint8), nil},
		// from float64
		{float64(-math.MaxFloat64), nil, cast.ErrUnableToCastToUint8},             // expected eror
		{float64(-math.SmallestNonzeroFloat64), nil, cast.ErrUnableToCastToUint8}, // expected eror
		{float64(0.0), uint8(0), nil},
		{float64(math.SmallestNonzeroFloat64), uint8(0), nil},
		{float64(math.MaxFloat64), nil, cast.ErrUnableToCastToUint8}, // expected eror
		// from float32
		{float32(-math.MaxFloat32), nil, cast.ErrUnableToCastToUint8},             // expected eror
		{float32(-math.SmallestNonzeroFloat32), nil, cast.ErrUnableToCastToUint8}, // expected eror
		{float32(0.0), uint8(0), nil},
		{float32(math.SmallestNonzeroFloat32), uint8(0), nil},
		{float32(math.MaxFloat32), nil, cast.ErrUnableToCastToUint8}, // expected eror
		// from bool
		{bool(true), uint8(1), nil},
		{bool(false), uint8(0), nil},
		// from string
		{string("-9223372036854775809"), nil, cast.ErrUnableToCastToUint8}, // expected eror
		{string("-9223372036854775808"), nil, cast.ErrUnableToCastToUint8}, // expected eror
		{string("-1"), nil, cast.ErrUnableToCastToUint8},                   // expected eror
		{string("0"), uint8(0), nil},
		{string("1"), uint8(1), nil},
		{string(fmt.Sprintf("%v", math.MaxInt8)), uint8(math.MaxInt8), nil},
		{string("9223372036854775808"), nil, cast.ErrUnableToCastToUint8},      // expected eror
		{string("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint8}, // expected eror
		{string("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint8}, // expected eror
		{string("-1.4"), nil, cast.ErrUnableToCastToUint8},                     // expected eror
		{string("0.0"), nil, cast.ErrUnableToCastToUint8},                      // expected eror ?
		{string("1.4"), nil, cast.ErrUnableToCastToUint8},                      // expected eror ?
		{string("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint8},  // expected eror ?
		{string("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint8},  // expected eror
		{string("hello"), nil, cast.ErrUnableToCastToUint8},                    // expected eror
		// from []byte
		{[]byte("-9223372036854775809"), nil, cast.ErrUnableToCastToUint8},     // expected eror
		{[]byte("-9223372036854775808"), nil, cast.ErrUnableToCastToUint8},     // expected eror
		{[]byte("-1"), nil, cast.ErrUnableToCastToUint8},                       // expected eror
		{[]byte("0"), uint8(48), nil},                                          // technically it's 1 byte
		{[]byte("1"), uint8(49), nil},                                          // technically it's 1 byte
		{[]byte("9223372036854775807"), nil, cast.ErrUnableToCastToUint8},      // expected eror
		{[]byte("9223372036854775808"), nil, cast.ErrUnableToCastToUint8},      // expected eror
		{[]byte("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint8}, // expected eror
		{[]byte("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint8}, // expected eror
		{[]byte("-1.4"), nil, cast.ErrUnableToCastToUint8},                     // expected eror
		{[]byte("0.0"), nil, cast.ErrUnableToCastToUint8},                      // expected eror
		{[]byte("1.4"), nil, cast.ErrUnableToCastToUint8},                      // expected eror
		{[]byte("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint8},  // expected eror
		{[]byte("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint8},  // expected eror
		{[]byte("hello"), nil, cast.ErrUnableToCastToUint8},                    // expected eror
		{[]byte{0xff}, uint8(math.MaxUint8), nil},
		{[]byte{0x0}, uint8(0), nil},
		{[]byte{0x1}, uint8(1), nil},
		{[]byte{0x7f}, uint8(math.MaxInt8), nil},
		// from number
		{json.Number(fmt.Sprintf("%v", math.MaxInt8)), uint8(math.MaxInt8), nil},
		{json.Number("9223372036854775808"), nil, cast.ErrUnableToCastToUint8},      // expected eror
		{json.Number("-1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint8}, // expected eror
		{json.Number("-1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint8}, // expected eror
		{json.Number("-1.4"), nil, cast.ErrUnableToCastToUint8},                     // expected eror
		{json.Number("0.0"), nil, cast.ErrUnableToCastToUint8},                      // expected eror ?
		{json.Number("1.4"), nil, cast.ErrUnableToCastToUint8},                      // expected eror ?
		{json.Number("1.7976931348623157e+308"), nil, cast.ErrUnableToCastToUint8},  // expected eror ?
		{json.Number("1.7976931348623158e+308"), nil, cast.ErrUnableToCastToUint8},  // expected eror ?
		{json.Number("hello"), nil, cast.ErrUnableToCastToUint8},                    // expected eror
		// from anything else
		{struct{ string }{""}, nil, cast.ErrUnableToCastToUint8},  // expected eror
		{[]int{1}, nil, cast.ErrUnableToCastToUint8},              // expected eror
		{map[string]int{"": 1}, nil, cast.ErrUnableToCastToUint8}, // expected eror
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td.value, td.value), func(t *testing.T) {
			result, err := cast.ToUint8(td.value)
			assert.ErrorIs(t, err, td.err)
			assert.Equal(t, td.expected, result)
		})
	}
}

func BenchmarkCastToUintZero(b *testing.B) { castToUint(b, 0) }

func BenchmarkCastToUintMaxInt64(b *testing.B) { castToUint(b, uint64(math.MaxInt64)) }

func BenchmarkCastToUintMaxInt32(b *testing.B) { castToUint(b, uint32(math.MaxInt32)) }

func BenchmarkCastToUintMaxInt16(b *testing.B) { castToUint(b, uint16(math.MaxInt16)) }

func BenchmarkCastToUintMaxInt8(b *testing.B) { castToUint(b, uint8(math.MaxInt8)) }

func BenchmarkCastToUintMaxFloat64(b *testing.B) { castToUint(b, float64(math.MaxFloat64)) }

func BenchmarkCastToUintMaxFloat32(b *testing.B) { castToUint(b, float32(math.MaxFloat32)) }

func BenchmarkCastToUintSmallestFloat64(b *testing.B) {
	castToUint(b, float64(math.SmallestNonzeroFloat64))
}

func BenchmarkCastToUintSmallestFloat32(b *testing.B) {
	castToUint(b, float32(math.SmallestNonzeroFloat32))
}

func BenchmarkCastToUintByte(b *testing.B) {
	castToUint(b, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
}

func BenchmarkCastToUintByteOverflow(b *testing.B) {
	castToUint(b, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
}

func castToUint(b *testing.B, i interface{}) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		cast.ToUint(i) //nolint:errcheck
	}
}

func BenchmarkCastToUint8Zero(b *testing.B) { castToUint8(b, 0) }

func BenchmarkCastToUint8MaxInt64(b *testing.B) { castToUint8(b, uint64(math.MaxInt64)) }

func BenchmarkCastToUint8MaxInt32(b *testing.B) { castToUint8(b, uint32(math.MaxInt32)) }

func BenchmarkCastToUint8MaxInt16(b *testing.B) { castToUint8(b, uint16(math.MaxInt16)) }

func BenchmarkCastToUint8MaxInt8(b *testing.B) { castToUint8(b, uint8(math.MaxInt8)) }

func BenchmarkCastToUint8MaxFloat64(b *testing.B) { castToUint8(b, float64(math.MaxFloat64)) }

func BenchmarkCastToUint8MaxFloat32(b *testing.B) { castToUint8(b, float32(math.MaxFloat32)) }

func BenchmarkCastToUint8SmallestFloat64(b *testing.B) {
	castToUint8(b, float64(math.SmallestNonzeroFloat64))
}

func BenchmarkCastToUint8SmallestFloat32(b *testing.B) {
	castToUint8(b, float32(math.SmallestNonzeroFloat32))
}

func BenchmarkCastToUint8Byte(b *testing.B) {
	castToUint8(b, []byte{0xff})
}

func BenchmarkCastToUint8ByteOverflow(b *testing.B) {
	castToUint8(b, []byte{0xff, 0xff})
}

func castToUint8(b *testing.B, i interface{}) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		cast.ToUint8(i) //nolint:errcheck
	}
}
