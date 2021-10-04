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
	"unsafe"
)

const (
	sizeOfInt64   = 8
	sizeOfInt32   = 4
	sizeOfInt16   = 2
	sizeOfUint64  = 8
	sizeOfUint32  = 4
	sizeOfUint16  = 2
	sizeOfFloat64 = 8
	sizeOfFloat32 = 4
	sizeOfInt     = int(unsafe.Sizeof(int(0)))
	sizeOfUint    = int(unsafe.Sizeof(uint(0)))
)

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

func int32ToBytes(num int32) []byte {
	arr := make([]byte, sizeOfInt32)

	for i := 0; i < sizeOfInt32; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[i] = byt
	}

	return arr
}

func int32FromBytes(arr []byte) (interface{}, error) {
	var val int32

	if arr == nil || len(arr) != sizeOfInt32 {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToInt32, arr, arr)
	}

	for i := 0; i < sizeOfInt32; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}

	return val, nil
}

func int16ToBytes(num int16) []byte {
	arr := make([]byte, sizeOfInt16)

	for i := 0; i < sizeOfInt16; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[i] = byt
	}

	return arr
}

func int16FromBytes(arr []byte) (interface{}, error) {
	var val int16

	if arr == nil || len(arr) != sizeOfInt16 {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToInt16, arr, arr)
	}

	for i := 0; i < sizeOfInt16; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}

	return val, nil
}

func uintToBytes(num uint) []byte {
	arr := make([]byte, sizeOfUint)

	for i := 0; i < sizeOfUint; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[i] = byt
	}

	return arr
}

func uintFromBytes(arr []byte) (interface{}, error) {
	var val uint

	if arr == nil || len(arr) != sizeOfUint {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToUint, arr, arr)
	}

	for i := 0; i < sizeOfUint; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}

	return val, nil
}

func uint64ToBytes(num uint64) []byte {
	arr := make([]byte, sizeOfUint64)

	for i := 0; i < sizeOfUint64; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[i] = byt
	}

	return arr
}

func uint64FromBytes(arr []byte) (interface{}, error) {
	var val uint64

	if arr == nil || len(arr) != sizeOfUint64 {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToUint64, arr, arr)
	}

	for i := 0; i < sizeOfUint64; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}

	return val, nil
}

func uint32ToBytes(num uint32) []byte {
	arr := make([]byte, sizeOfUint32)

	for i := 0; i < sizeOfUint32; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[i] = byt
	}

	return arr
}

func uint32FromBytes(arr []byte) (interface{}, error) {
	var val uint32

	if arr == nil || len(arr) != sizeOfUint32 {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToUint32, arr, arr)
	}

	for i := 0; i < sizeOfUint32; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}

	return val, nil
}

func uint16ToBytes(num uint16) []byte {
	arr := make([]byte, sizeOfUint16)

	for i := 0; i < sizeOfUint16; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[i] = byt
	}

	return arr
}

func uint16FromBytes(arr []byte) (interface{}, error) {
	var val uint16

	if arr == nil || len(arr) != sizeOfUint16 {
		return nil, fmt.Errorf("%w: %T(%v)", ErrUnableToCastToUint16, arr, arr)
	}

	for i := 0; i < sizeOfUint16; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}

	return val, nil
}
