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
	"encoding/json"
	"fmt"
	"math"
	"unsafe"
)

//nolint:cyclop,funlen,gomnd
func ToBinary(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil, []byte:
		return val, nil
	case float64:
		bits := math.Float64bits(val)
		bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, bits)

		return bytes, nil
	case float32:
		bits := math.Float32bits(val)
		bytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(bytes, bits)

		return bytes, nil
	case int:
		size := unsafe.Sizeof(val)
		bytes := make([]byte, size)

		switch size {
		case 64:
			binary.LittleEndian.PutUint64(bytes, uint64(val))
		default: // case 32:
			binary.LittleEndian.PutUint32(bytes, uint32(val))
		}

		return bytes, nil
	case int64:
		bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, uint64(val))

		return bytes, nil
	case int32:
		bytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(bytes, uint32(val))

		return bytes, nil
	case int16:
		bytes := make([]byte, 2)
		binary.LittleEndian.PutUint16(bytes, uint16(val))

		return bytes, nil
	case int8:
		bytes := make([]byte, 1)
		bytes[0] = byte(val)

		return bytes, nil
	case uint:
		size := unsafe.Sizeof(val)
		bytes := make([]byte, size)

		switch size {
		case 64:
			binary.LittleEndian.PutUint64(bytes, uint64(val))
		default: // case 32:
			binary.LittleEndian.PutUint32(bytes, uint32(val))
		}

		return bytes, nil
	case uint64:
		bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, val)

		return bytes, nil
	case uint32:
		bytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(bytes, val)

		return bytes, nil
	case uint16:
		bytes := make([]byte, 2)
		binary.LittleEndian.PutUint16(bytes, val)

		return bytes, nil
	case uint8:
		bytes := make([]byte, 1)
		bytes[0] = val

		return bytes, nil
	case bool:
		b := byte(0)
		if val {
			b = 1
		}

		return []byte{b}, nil
	case string:
		return []byte(val), nil
	case json.Number:
		return ToBool(string(val))
	default:
		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToBinary, i, i)
	}
}
