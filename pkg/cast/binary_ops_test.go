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

package cast //nolint:testpackage

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnsafeInt(t *testing.T) {
	testdatas := []int{math.MinInt, -1, 0, 1, math.MaxInt}
	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td, td), func(t *testing.T) {
			bytes := intToBytes(td)
			result, err := intFromBytes(bytes)
			assert.NoError(t, err)
			assert.Equal(t, td, result)
		})
	}
}

func TestUnsafeInt64(t *testing.T) {
	testdatas := []int64{math.MinInt64, -1, 0, 1, math.MaxInt64}
	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td, td), func(t *testing.T) {
			bytes := int64ToBytes(td)
			result, err := int64FromBytes(bytes)
			assert.NoError(t, err)
			assert.Equal(t, td, result)
		})
	}
}

func TestUnsafeInt32(t *testing.T) {
	testdatas := []int32{math.MinInt32, -1, 0, 1, math.MaxInt32}
	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td, td), func(t *testing.T) {
			bytes := int32ToBytes(td)
			result, err := int32FromBytes(bytes)
			assert.NoError(t, err)
			assert.Equal(t, td, result)
		})
	}
}

func TestUnsafeInt16(t *testing.T) {
	testdatas := []int16{math.MinInt16, -1, 0, 1, math.MaxInt16}
	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td, td), func(t *testing.T) {
			bytes := int16ToBytes(td)
			result, err := int16FromBytes(bytes)
			assert.NoError(t, err)
			assert.Equal(t, td, result)
		})
	}
}

func TestUnsafeUint(t *testing.T) {
	testdatas := []uint{0, 1, math.MaxInt, math.MaxUint}
	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td, td), func(t *testing.T) {
			bytes := uintToBytes(td)
			result, err := uintFromBytes(bytes)
			assert.NoError(t, err)
			assert.Equal(t, td, result)
		})
	}
}

func TestUnsafeUint64(t *testing.T) {
	testdatas := []uint64{0, 1, math.MaxInt64, math.MaxUint64}
	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td, td), func(t *testing.T) {
			bytes := uint64ToBytes(td)
			result, err := uint64FromBytes(bytes)
			assert.NoError(t, err)
			assert.Equal(t, td, result)
		})
	}
}

func TestUnsafeUint32(t *testing.T) {
	testdatas := []uint32{0, 1, math.MaxInt32, math.MaxUint32}
	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td, td), func(t *testing.T) {
			bytes := uint32ToBytes(td)
			result, err := uint32FromBytes(bytes)
			assert.NoError(t, err)
			assert.Equal(t, td, result)
		})
	}
}

func TestUnsafeUint16(t *testing.T) {
	testdatas := []uint16{0, 1, math.MaxInt16, math.MaxUint16}
	for _, td := range testdatas {
		t.Run(fmt.Sprintf("%T(%v)", td, td), func(t *testing.T) {
			bytes := uint16ToBytes(td)
			result, err := uint16FromBytes(bytes)
			assert.NoError(t, err)
			assert.Equal(t, td, result)
		})
	}
}
