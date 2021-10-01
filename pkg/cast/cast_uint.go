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
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
)

//nolint:cyclop,funlen
func ToUint(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil:
		return nil, nil
	case int:
		return uint(val), nil
	case int64:
		return uint(val), nil
	case int32:
		return uint(val), nil
	case int16:
		return uint(val), nil
	case int8:
		return uint(val), nil
	case uint:
		return val, nil
	case uint64:
		return uint(val), nil
	case uint32:
		return uint(val), nil
	case uint16:
		return uint(val), nil
	case uint8:
		return uint(val), nil
	case float64:
		return uint(val), nil
	case float32:
		return uint(val), nil
	case bool:
		if val {
			return uint(1), nil
		}

		return uint(0), nil
	case string:
		v, err := strconv.ParseUint(val, 0, 0)
		if err == nil {
			return uint(v), nil
		}

		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToUint, i, i)
	case []byte:
		var v uint

		buf := bytes.NewReader(val)
		if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
			return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToUint, i, i)
		}

		return v, nil
	case json.Number:
		return ToUint(string(val))
	default:
		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToUint, i, i)
	}
}

//nolint:cyclop,funlen
func ToUint64(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil:
		return nil, nil
	case int:
		return uint64(val), nil
	case int64:
		return uint64(val), nil
	case int32:
		return uint64(val), nil
	case int16:
		return uint64(val), nil
	case int8:
		return uint64(val), nil
	case uint:
		return uint64(val), nil
	case uint64:
		return val, nil
	case uint32:
		return uint64(val), nil
	case uint16:
		return uint64(val), nil
	case uint8:
		return uint64(val), nil
	case float64:
		return uint64(val), nil
	case float32:
		return uint64(val), nil
	case bool:
		if val {
			return uint64(1), nil
		}

		return uint64(0), nil
	case string:
		v, err := strconv.ParseUint(val, 0, 64) //nolint:gomnd
		if err == nil {
			return v, nil
		}

		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToUint64, i, i)
	case []byte:
		var v uint64

		buf := bytes.NewReader(val)
		if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
			return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToUint64, i, i)
		}

		return v, nil
	case json.Number:
		return ToUint64(string(val))
	default:
		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToUint64, i, i)
	}
}

//nolint:cyclop,funlen
func ToUint32(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil:
		return nil, nil
	case int:
		return uint32(val), nil
	case int64:
		return uint32(val), nil
	case int32:
		return uint32(val), nil
	case int16:
		return uint32(val), nil
	case int8:
		return uint32(val), nil
	case uint:
		return uint32(val), nil
	case uint64:
		return uint32(val), nil
	case uint32:
		return val, nil
	case uint16:
		return uint32(val), nil
	case uint8:
		return uint32(val), nil
	case float64:
		return uint32(val), nil
	case float32:
		return uint32(val), nil
	case bool:
		if val {
			return uint32(1), nil
		}

		return uint32(0), nil
	case string:
		v, err := strconv.ParseUint(val, 0, 32) //nolint:gomnd
		if err == nil {
			return uint32(v), nil
		}

		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToUint32, i, i)
	case []byte:
		var v uint32

		buf := bytes.NewReader(val)
		if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
			return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToUint32, i, i)
		}

		return v, nil
	case json.Number:
		return ToUint32(string(val))
	default:
		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToUint32, i, i)
	}
}

//nolint:cyclop,funlen
func ToUint16(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil:
		return nil, nil
	case int:
		return uint16(val), nil
	case int64:
		return uint16(val), nil
	case int32:
		return uint16(val), nil
	case int16:
		return uint16(val), nil
	case int8:
		return uint16(val), nil
	case uint:
		return uint16(val), nil
	case uint64:
		return uint16(val), nil
	case uint32:
		return uint16(val), nil
	case uint16:
		return val, nil
	case uint8:
		return uint16(val), nil
	case float64:
		return uint16(val), nil
	case float32:
		return uint16(val), nil
	case bool:
		if val {
			return uint16(1), nil
		}

		return uint16(0), nil
	case string:
		v, err := strconv.ParseUint(val, 0, 16) //nolint:gomnd
		if err == nil {
			return uint16(v), nil
		}

		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToUint16, i, i)
	case []byte:
		var v uint16

		buf := bytes.NewReader(val)
		if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
			return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToUint16, i, i)
		}

		return v, nil
	case json.Number:
		return ToUint16(string(val))
	default:
		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToUint16, i, i)
	}
}

//nolint:cyclop,funlen
func ToUint8(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil:
		return nil, nil
	case int:
		return uint8(val), nil
	case int64:
		return uint8(val), nil
	case int32:
		return uint8(val), nil
	case int16:
		return uint8(val), nil
	case int8:
		return uint8(val), nil
	case uint:
		return uint8(val), nil
	case uint64:
		return uint8(val), nil
	case uint32:
		return uint8(val), nil
	case uint16:
		return uint8(val), nil
	case uint8:
		return val, nil
	case float64:
		return uint8(val), nil
	case float32:
		return uint8(val), nil
	case bool:
		if val {
			return uint8(1), nil
		}

		return uint8(0), nil
	case string:
		v, err := strconv.ParseUint(val, 0, 8) //nolint:gomnd
		if err == nil {
			return uint8(v), nil
		}

		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToUint8, i, i)
	case []byte:
		var v uint8

		buf := bytes.NewReader(val)
		if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
			return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToUint8, i, i)
		}

		return v, nil
	case json.Number:
		return ToUint8(string(val))
	default:
		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToUint8, i, i)
	}
}
