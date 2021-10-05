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
//
// Linking this library statically or dynamically with other modules is
// making a combined work based on this library.  Thus, the terms and
// conditions of the GNU General Public License cover the whole
// combination.
//
// As a special exception, the copyright holders of this library give you
// permission to link this library with independent modules to produce an
// executable, regardless of the license terms of these independent
// modules, and to copy and distribute the resulting executable under
// terms of your choice, provided that you also meet, for each linked
// independent module, the terms and conditions of the license of that
// module.  An independent module is a module which is not derived from
// or based on this library.  If you modify this library, you may extend
// this exception to your version of the library, but you are not
// obligated to do so.  If you do not wish to do so, delete this
// exception statement from your version.

package cast

import (
	"encoding/binary"
	"fmt"
	"math"
	"unsafe"
)

const (
	sizeOfInt64   = 8
	sizeOfInt32   = 4
	sizeOfInt16   = 2
	sizeOfInt8    = 1
	sizeOfUint64  = 8
	sizeOfUint32  = 4
	sizeOfUint16  = 2
	sizeOfUint8   = 1
	sizeOfFloat64 = 8
	sizeOfFloat32 = 4
	sizeOfBool    = 1
	sizeOfInt     = int(unsafe.Sizeof(int(0)))
	sizeOfUint    = int(unsafe.Sizeof(uint(0)))
)

func float64ToBytes(val float64) []byte {
	bytes := make([]byte, sizeOfFloat64)
	binary.LittleEndian.PutUint64(bytes, math.Float64bits(val))

	return bytes
}

func float64FromBytes(bytes []byte) (interface{}, error) {
	if bytes == nil || len(bytes) != sizeOfFloat64 {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToFloat64, bytes, bytes)
	}

	return math.Float64frombits(binary.LittleEndian.Uint64(bytes)), nil
}

func float32ToBytes(val float32) []byte {
	bytes := make([]byte, sizeOfFloat32)
	binary.LittleEndian.PutUint32(bytes, math.Float32bits(val))

	return bytes
}

func float32FromBytes(bytes []byte) (interface{}, error) {
	if bytes == nil || len(bytes) != sizeOfFloat32 {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToFloat32, bytes, bytes)
	}

	return math.Float32frombits(binary.LittleEndian.Uint32(bytes)), nil
}

func intToBytes(val int) []byte {
	bytes := make([]byte, sizeOfInt)
	if sizeOfInt == sizeOfInt64 {
		binary.LittleEndian.PutUint64(bytes, uint64(val))
	} else {
		binary.LittleEndian.PutUint32(bytes, uint32(val))
	}

	return bytes
}

func intFromBytes(bytes []byte) (interface{}, error) {
	var val int

	if bytes == nil || len(bytes) != sizeOfInt {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToInt, bytes, bytes)
	}

	if sizeOfInt == sizeOfInt64 {
		val = int(binary.LittleEndian.Uint64(bytes))
	} else {
		val = int(binary.LittleEndian.Uint32(bytes))
	}

	return val, nil
}

func int64ToBytes(val int64) []byte {
	bytes := make([]byte, sizeOfInt64)

	binary.LittleEndian.PutUint64(bytes, uint64(val))

	return bytes
}

func int64FromBytes(bytes []byte) (interface{}, error) {
	if bytes == nil || len(bytes) != sizeOfInt64 {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToInt64, bytes, bytes)
	}

	return int64(binary.LittleEndian.Uint64(bytes)), nil
}

func int32ToBytes(val int32) []byte {
	bytes := make([]byte, sizeOfInt32)

	binary.LittleEndian.PutUint32(bytes, uint32(val))

	return bytes
}

func int32FromBytes(bytes []byte) (interface{}, error) {
	if bytes == nil || len(bytes) != sizeOfInt32 {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToInt32, bytes, bytes)
	}

	return int32(binary.LittleEndian.Uint32(bytes)), nil
}

func int16ToBytes(val int16) []byte {
	bytes := make([]byte, sizeOfInt16)

	binary.LittleEndian.PutUint16(bytes, uint16(val))

	return bytes
}

func int16FromBytes(bytes []byte) (interface{}, error) {
	if bytes == nil || len(bytes) != sizeOfInt16 {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToInt16, bytes, bytes)
	}

	return int16(binary.LittleEndian.Uint16(bytes)), nil
}

func int8ToBytes(val int8) []byte {
	bytes := make([]byte, sizeOfInt8)
	bytes[0] = byte(val)

	return bytes
}

func int8FromBytes(bytes []byte) (interface{}, error) {
	if len(bytes) != sizeOfInt8 {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToInt8, bytes, bytes)
	}

	return int8(bytes[0]), nil
}

func uintToBytes(val uint) []byte {
	bytes := make([]byte, sizeOfUint)
	if sizeOfUint == sizeOfUint64 {
		binary.LittleEndian.PutUint64(bytes, uint64(val))
	} else {
		binary.LittleEndian.PutUint32(bytes, uint32(val))
	}

	return bytes
}

func uintFromBytes(bytes []byte) (interface{}, error) {
	var val uint

	if bytes == nil || len(bytes) != sizeOfUint {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToUint, bytes, bytes)
	}

	if sizeOfUint == sizeOfUint64 {
		val = uint(binary.LittleEndian.Uint64(bytes))
	} else {
		val = uint(binary.LittleEndian.Uint32(bytes))
	}

	return val, nil
}

func uint64ToBytes(val uint64) []byte {
	bytes := make([]byte, sizeOfUint64)

	binary.LittleEndian.PutUint64(bytes, val)

	return bytes
}

func uint64FromBytes(bytes []byte) (interface{}, error) {
	if bytes == nil || len(bytes) != sizeOfUint64 {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToUint64, bytes, bytes)
	}

	return binary.LittleEndian.Uint64(bytes), nil
}

func uint32ToBytes(val uint32) []byte {
	bytes := make([]byte, sizeOfUint32)

	binary.LittleEndian.PutUint32(bytes, val)

	return bytes
}

func uint32FromBytes(bytes []byte) (interface{}, error) {
	if bytes == nil || len(bytes) != sizeOfUint32 {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToUint32, bytes, bytes)
	}

	return binary.LittleEndian.Uint32(bytes), nil
}

func uint16ToBytes(val uint16) []byte {
	bytes := make([]byte, sizeOfUint16)

	binary.LittleEndian.PutUint16(bytes, val)

	return bytes
}

func uint16FromBytes(bytes []byte) (interface{}, error) {
	if bytes == nil || len(bytes) != sizeOfUint16 {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToUint16, bytes, bytes)
	}

	return binary.LittleEndian.Uint16(bytes), nil
}

func uint8ToBytes(val uint8) []byte {
	bytes := make([]byte, sizeOfUint8)
	bytes[0] = val

	return bytes
}

func uint8FromBytes(bytes []byte) (interface{}, error) {
	if len(bytes) != sizeOfUint8 {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToUint8, bytes, bytes)
	}

	return bytes[0], nil
}

func boolToBytes(val bool) []byte {
	bytes := make([]byte, sizeOfBool)
	if val {
		bytes[0] = 1
	}

	return bytes
}

func boolFromBytes(bytes []byte) (interface{}, error) {
	if len(bytes) != sizeOfBool {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToBool, bytes, bytes)
	}

	return bytes[0] != 0, nil
}
