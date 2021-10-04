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

//nolint:funlen
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

func TestCastToBool(t *testing.T) {
	testdatas := []struct {
		value    interface{}
		expected interface{}
		err      error
	}{
		{nil, nil, nil},
		// from int
		{int(math.MinInt), bool(true), nil},
		{int(0), bool(false), nil},
		{int(math.MaxInt), bool(true), nil},
		// from int64
		{int64(math.MinInt64), bool(true), nil},
		{int64(0), bool(false), nil},
		{int64(math.MaxInt64), bool(true), nil},
		// from int32
		{int32(math.MinInt32), bool(true), nil},
		{int32(0), bool(false), nil},
		{int32(math.MaxInt32), bool(true), nil},
		// from int16
		{int16(math.MinInt16), bool(true), nil},
		{int16(0), bool(false), nil},
		{int16(math.MaxInt16), bool(true), nil},
		// from int8
		{int8(math.MinInt8), bool(true), nil},
		{int8(0), bool(false), nil},
		{int8(math.MaxInt8), bool(true), nil},
		// from uint
		{uint(0), bool(false), nil},
		{uint(math.MaxUint), bool(true), nil},
		// from uint64
		{uint64(0), bool(false), nil},
		{uint64(math.MaxUint64), bool(true), nil},
		// from uint32
		{uint32(0), bool(false), nil},
		{uint32(math.MaxUint32), bool(true), nil},
		// from iunt16
		{uint16(0), bool(false), nil},
		{uint16(math.MaxUint16), bool(true), nil},
		// from uint8
		{uint8(0), bool(false), nil},
		{uint8(math.MaxUint8), bool(true), nil},
		// from float64
		{float64(-math.MaxFloat64), bool(true), nil},
		{float64(-math.SmallestNonzeroFloat64), bool(true), nil},
		{float64(0.0), bool(false), nil},
		{float64(math.SmallestNonzeroFloat64), bool(true), nil},
		{float64(math.MaxFloat64), bool(true), nil},
		// from float32
		{float32(-math.MaxFloat32), bool(true), nil},
		{float32(-math.SmallestNonzeroFloat32), bool(true), nil},
		{float32(0.0), bool(false), nil},
		{float32(math.SmallestNonzeroFloat32), bool(true), nil},
		{float32(math.MaxFloat32), bool(true), nil},
		// from bool
		{bool(true), bool(true), nil},
		{bool(false), bool(false), nil},
		// from string
		{string("-9223372036854775809"), bool(true), nil},
		{string("-9223372036854775808"), bool(true), nil},
		{string("-1"), bool(true), nil},
		{bool(false), bool(false), nil},
		{string("1"), bool(true), nil},
		{string("9223372036854775807"), bool(true), nil},
		{string("9223372036854775808"), bool(true), nil},
		{string("-1.7976931348623158e+308"), bool(true), nil},
		{string("-1.7976931348623157e+308"), bool(true), nil},
		{string("-1.4"), bool(true), nil},
		{string("0.0"), bool(false), nil},
		{string("1.4"), bool(true), nil},
		{string("1.7976931348623157e+308"), bool(true), nil},
		{string("1.7976931348623158e+308"), bool(true), nil},
		{string("hello"), nil, cast.ErrUnableToCastToBool}, // expected error
		{string("true"), bool(true), nil},
		{string("false"), bool(false), nil},
		// from []byte
		{[]byte("-9223372036854775809"), bool(true), nil},
		{[]byte("-9223372036854775808"), bool(true), nil},
		{[]byte("-1"), bool(true), nil},
		{[]byte("0"), bool(false), nil},
		{[]byte("1"), bool(true), nil},
		{[]byte("9223372036854775807"), bool(true), nil},
		{[]byte("9223372036854775808"), bool(true), nil},
		{[]byte("-1.7976931348623158e+308"), bool(true), nil},
		{[]byte("-1.7976931348623157e+308"), bool(true), nil},
		{[]byte("-1.4"), bool(true), nil},
		{[]byte("0.0"), bool(false), nil},
		{[]byte("1.4"), bool(true), nil},
		{[]byte("1.7976931348623157e+308"), bool(true), nil},
		{[]byte("1.7976931348623158e+308"), bool(true), nil},
		{[]byte("hello"), nil, cast.ErrUnableToCastToBool},                                        // expected error
		{[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80}, nil, cast.ErrUnableToCastToBool},        // expected error
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, nil, cast.ErrUnableToCastToBool}, // expected error
		{[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, nil, cast.ErrUnableToCastToBool},         // expected error
		{[]byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, nil, cast.ErrUnableToCastToBool},         // expected error
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}, nil, cast.ErrUnableToCastToBool}, // expected error
		// from number
		{json.Number("9223372036854775807"), bool(true), nil},
		{json.Number("9223372036854775808"), bool(true), nil},
		{json.Number("-1.7976931348623158e+308"), bool(true), nil},
		{json.Number("-1.7976931348623157e+308"), bool(true), nil},
		{json.Number("-1.4"), bool(true), nil},
		{json.Number("0.0"), bool(false), nil},
		{json.Number("1.4"), bool(true), nil},
		{json.Number("1.7976931348623157e+308"), bool(true), nil},
		{json.Number("1.7976931348623158e+308"), bool(true), nil},
		{json.Number("hello"), nil, cast.ErrUnableToCastToBool}, // expected error
		// from time
		{time.Date(2021, time.October, 4, 13, 3, 56, 0, time.UTC), bool(true), nil},
		{time.Unix(0, 0), bool(false), nil},
		// from anything else
		{struct{ string }{""}, nil, cast.ErrUnableToCastToBool},  // expected error
		{[]int{1}, nil, cast.ErrUnableToCastToBool},              // expected error
		{map[string]int{"": 1}, nil, cast.ErrUnableToCastToBool}, // expected error
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td.value, td.value), func(t *testing.T) {
			result, err := cast.ToBool(td.value)
			assert.ErrorIs(t, err, td.err)
			assert.Equal(t, td.expected, result)
		})
	}
}

func BenchmarkCastToBoolTrue(b *testing.B) { castToBool(b, "true") }

func BenchmarkCastToBoolFalse(b *testing.B) { castToBool(b, "false") }

func castToBool(b *testing.B, i interface{}) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		cast.ToBool(i) //nolint:errcheck
	}
}
