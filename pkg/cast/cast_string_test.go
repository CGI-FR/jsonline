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

//nolint:funlen,lll
package cast_test

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/cgi-fr/jsonline/pkg/cast"
	"github.com/stretchr/testify/assert"
)

func TestCastToString(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
		err      error
	}{
		{nil, nil, nil},
		// from int
		{int(math.MinInt), string("-9223372036854775808"), nil},
		{int(0), string("0"), nil},
		{int(math.MaxInt), string("9223372036854775807"), nil},
		// from int64
		{int64(math.MinInt64), string("-9223372036854775808"), nil},
		{int64(0), string("0"), nil},
		{int64(math.MaxInt64), string("9223372036854775807"), nil},
		// from int32
		{int32(math.MinInt32), string("-2147483648"), nil},
		{int32(0), string("0"), nil},
		{int32(math.MaxInt32), string("2147483647"), nil},
		// from int16
		{int16(math.MinInt16), string("-32768"), nil},
		{int16(0), string("0"), nil},
		{int16(math.MaxInt16), string("32767"), nil},
		// from int8
		{int8(math.MinInt8), string("-128"), nil},
		{int8(0), string("0"), nil},
		{int8(math.MaxInt8), string("127"), nil},
		// from uint
		{uint(0), string("0"), nil},
		{uint(math.MaxUint), string("18446744073709551615"), nil},
		// from uint64
		{uint64(0), string("0"), nil},
		{uint64(math.MaxUint64), string("18446744073709551615"), nil},
		// from uint32
		{uint32(0), string("0"), nil},
		{uint32(math.MaxUint32), string("4294967295"), nil},
		// from iunt16
		{uint16(0), string("0"), nil},
		{uint16(math.MaxUint16), string("65535"), nil},
		// from uint8
		{uint8(0), string("0"), nil},
		{uint8(math.MaxUint8), string("255"), nil},
		// from float64
		{float64(-math.MaxFloat64), string("-179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"), nil},
		{float64(-math.SmallestNonzeroFloat64), string("-0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005"), nil},
		{float64(0.0), string("0"), nil},
		{float64(math.SmallestNonzeroFloat64), string("0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005"), nil},
		{float64(math.MaxFloat64), string("179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"), nil},
		// from float32
		{float32(-math.MaxFloat32), string("-340282350000000000000000000000000000000"), nil},
		{float32(-math.SmallestNonzeroFloat32), string("-0.000000000000000000000000000000000000000000001"), nil},
		{float32(0.0), string("0"), nil},
		{float32(math.SmallestNonzeroFloat32), string("0.000000000000000000000000000000000000000000001"), nil},
		{float32(math.MaxFloat32), string("340282350000000000000000000000000000000"), nil},
		// from bool
		{bool(true), string("true"), nil},
		{bool(false), string("false"), nil},
		// from string
		{string("-9223372036854775809"), string("-9223372036854775809"), nil},
		{string("-9223372036854775808"), string("-9223372036854775808"), nil},
		{string("-1"), string("-1"), nil},
		{string("0"), string("0"), nil},
		{string("1"), string("1"), nil},
		{string("9223372036854775807"), string("9223372036854775807"), nil},
		{string("9223372036854775808"), string("9223372036854775808"), nil},
		{string("-1.7976931348623158e+308"), string("-1.7976931348623158e+308"), nil},
		{string("-1.7976931348623157e+308"), string("-1.7976931348623157e+308"), nil},
		{string("-1.4"), string("-1.4"), nil},
		{string("0.0"), string("0.0"), nil},
		{string("1.4"), string("1.4"), nil},
		{string("1.7976931348623157e+308"), string("1.7976931348623157e+308"), nil},
		{string("1.7976931348623158e+308"), string("1.7976931348623158e+308"), nil},
		{string("hello"), string("hello"), nil},
		// from []byte
		{[]byte("-9223372036854775809"), string("-9223372036854775809"), nil},
		{[]byte("-9223372036854775808"), string("-9223372036854775808"), nil},
		{[]byte("-1"), string("-1"), nil},
		{[]byte("0"), string("0"), nil},
		{[]byte("1"), string("1"), nil},
		{[]byte("9223372036854775807"), string("9223372036854775807"), nil},
		{[]byte("9223372036854775808"), string("9223372036854775808"), nil},
		{[]byte("-1.7976931348623158e+308"), string("-1.7976931348623158e+308"), nil},
		{[]byte("-1.7976931348623157e+308"), string("-1.7976931348623157e+308"), nil},
		{[]byte("-1.4"), string("-1.4"), nil},
		{[]byte("0.0"), string("0.0"), nil},
		{[]byte("1.4"), string("1.4"), nil},
		{[]byte("1.7976931348623157e+308"), string("1.7976931348623157e+308"), nil},
		{[]byte("1.7976931348623158e+308"), string("1.7976931348623158e+308"), nil},
		{[]byte("hello"), string("hello"), nil},
		{[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80}, string("\x00\x00\x00\x00\x00\x00\x00\x80"), nil},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, string("\xff\xff\xff\xff\xff\xff\xff\xff"), nil},
		{[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, string("\x00\x00\x00\x00\x00\x00\x00\x00"), nil},
		{[]byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, string("\x01\x00\x00\x00\x00\x00\x00\x00"), nil},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}, string("\xff\xff\xff\xff\xff\xff\xff\u007f"), nil},
		// from number
		{json.Number("9223372036854775807"), string("9223372036854775807"), nil},
		{json.Number("9223372036854775808"), string("9223372036854775808"), nil},
		{json.Number("-1.7976931348623158e+308"), string("-1.7976931348623158e+308"), nil},
		{json.Number("-1.7976931348623157e+308"), string("-1.7976931348623157e+308"), nil},
		{json.Number("-1.4"), string("-1.4"), nil},
		{json.Number("0.0"), string("0.0"), nil},
		{json.Number("1.4"), string("1.4"), nil},
		{json.Number("1.7976931348623157e+308"), string("1.7976931348623157e+308"), nil},
		{json.Number("1.7976931348623158e+308"), string("1.7976931348623158e+308"), nil},
		{json.Number("hello"), string("hello"), nil},
		// from time
		{time.Date(2021, time.October, 4, 13, 3, 56, 0, time.UTC), string("2021-10-04T13:03:56Z"), nil},
		// from anything else
		{struct{ string }{""}, nil, cast.ErrUnableToCastToString},  // expected eror
		{[]int{1}, nil, cast.ErrUnableToCastToString},              // expected eror
		{map[string]int{"": 1}, nil, cast.ErrUnableToCastToString}, // expected eror
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td.value, td.value), func(t *testing.T) {
			result, err := cast.ToString(td.value)
			assert.ErrorIs(t, err, td.err)
			assert.Equal(t, td.expected, result)
		})
	}
}

func BenchmarkCastToStringEmpty(b *testing.B) { castToString(b, 0) }

func BenchmarkCastToStringMaxInt64(b *testing.B) { castToString(b, int64(math.MaxInt64)) }

func BenchmarkCastToStringMaxInt32(b *testing.B) { castToString(b, int32(math.MaxInt32)) }

func BenchmarkCastToStringMaxInt16(b *testing.B) { castToString(b, int16(math.MaxInt16)) }

func BenchmarkCastToStringMaxInt8(b *testing.B) { castToString(b, int8(math.MaxInt8)) }

func BenchmarkCastToStringMinInt64(b *testing.B) { castToString(b, int64(math.MinInt64)) }

func BenchmarkCastToStringMinInt32(b *testing.B) { castToString(b, int32(math.MinInt32)) }

func BenchmarkCastToStringMinInt16(b *testing.B) { castToString(b, int16(math.MinInt16)) }

func BenchmarkCastToStringMinInt8(b *testing.B) { castToString(b, int8(math.MinInt8)) }

func BenchmarkCastToStringMaxFloat64(b *testing.B) { castToString(b, float64(math.MaxFloat64)) }

func BenchmarkCastToStringMaxFloat32(b *testing.B) { castToString(b, float32(math.MaxFloat32)) }

func BenchmarkCastToStringSmallestFloat64(b *testing.B) {
	castToString(b, float64(math.SmallestNonzeroFloat64))
}

func BenchmarkCastToStringSmallestFloat32(b *testing.B) {
	castToString(b, float32(math.SmallestNonzeroFloat32))
}

func BenchmarkCastToStringByte(b *testing.B) {
	castToString(b, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
}

func BenchmarkCastToStringByteOverflow(b *testing.B) {
	castToString(b, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
}

func castToString(b *testing.B, i interface{}) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		cast.ToString(i) //nolint:errcheck
	}
}
