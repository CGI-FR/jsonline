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

//nolint:varnamelen
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

func TestCastBack(t *testing.T) {
	testdatas := []struct {
		value  interface{}
		target interface{}
	}{
		{float64(0), []byte{}},
		{float64(1.2), []byte{}},
		{float64(-1.2), []byte{}},
		{float64(math.SmallestNonzeroFloat64), []byte{}},
		{float64(math.MaxFloat64), []byte{}},
		{float32(0), []byte{}},
		{float32(1.2), []byte{}},
		{float32(-1.2), []byte{}},
		{float32(math.SmallestNonzeroFloat32), []byte{}},
		{float32(math.MaxFloat32), []byte{}},
		{int(0), []byte{}},
		{int(1), []byte{}},
		{int(-1), []byte{}},
		{int64(0), []byte{}},
		{int64(1), []byte{}},
		{int64(-1), []byte{}},
		{int32(0), []byte{}},
		{int32(1), []byte{}},
		{int32(-1), []byte{}},
		{int16(0), []byte{}},
		{int16(1), []byte{}},
		{int16(-1), []byte{}},
		{int8(0), []byte{}},
		{int8(1), []byte{}},
		{int8(-1), []byte{}},
		{uint(0), []byte{}},
		{uint(1), []byte{}},
		{uint64(0), []byte{}},
		{uint64(1), []byte{}},
		{uint32(0), []byte{}},
		{uint32(1), []byte{}},
		{uint16(0), []byte{}},
		{uint16(1), []byte{}},
		{uint8(0), []byte{}},
		{uint8(1), []byte{}},
		{bool(true), []byte{}},
		{bool(false), []byte{}},
		{"2021-10-04T13:03:56Z", time.Now()},
		// {time.Now(), string("")}, drop nanosecs
		{1, nil},
		{"true", bool(true)},
		{bool(true), string("")},
		{json.Number("53.56"), float64(0)},
	}

	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v) back from %T", td.value, td.value, td.target), func(t *testing.T) {
			result, err := cast.To(td.target, td.value)
			assert.NoError(t, err)
			result, err = cast.To(td.value, result)
			assert.NoError(t, err)
			assert.Equal(t, td.value, result)
		})
	}
}
