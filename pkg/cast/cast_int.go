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

//nolint:cyclop,funlen,gomnd
package cast

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
)

func ToInt(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil:
		return nil, nil
	case int:
		return val, nil
	case int64:
		return int(val), nil
	case int32:
		return int(val), nil
	case int16:
		return int(val), nil
	case int8:
		return int(val), nil
	case uint:
		return int(val), nil
	case uint64:
		return int(val), nil
	case uint32:
		return int(val), nil
	case uint16:
		return int(val), nil
	case uint8:
		return int(val), nil
	case float64:
		return int(val), nil
	case float32:
		return int(val), nil
	case bool:
		if val {
			return int(1), nil
		}

		return int(0), nil
	case string:
		v, err := strconv.ParseInt(val, 0, 0)
		if err == nil {
			return int(v), nil
		}

		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToInt, i, i)
	case []byte:
		var v int

		buf := bytes.NewReader(val)
		if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
			return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToInt, i, i)
		}

		return v, nil
	case json.Number:
		return ToInt(string(val))
	default:
		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToInt, i, i)
	}
}

func ToInt64(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil:
		return nil, nil
	case int:
		return int64(val), nil
	case int64:
		return val, nil
	case int32:
		return int64(val), nil
	case int16:
		return int64(val), nil
	case int8:
		return int64(val), nil
	case uint:
		return int64(val), nil
	case uint64:
		return int64(val), nil
	case uint32:
		return int64(val), nil
	case uint16:
		return int64(val), nil
	case uint8:
		return int64(val), nil
	case float64:
		return int64(val), nil
	case float32:
		return int64(val), nil
	case bool:
		if val {
			return int64(1), nil
		}

		return int64(0), nil
	case string:
		v, err := strconv.ParseInt(val, 0, 64)
		if err == nil {
			return v, nil
		}

		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToInt64, i, i)
	case []byte:
		var v int64

		buf := bytes.NewReader(val)
		if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
			return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToInt64, i, i)
		}

		return v, nil
	case json.Number:
		return ToInt64(string(val))
	default:
		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToInt64, i, i)
	}
}

func ToInt32(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil:
		return nil, nil
	case int:
		return int32(val), nil
	case int64:
		return int32(val), nil
	case int32:
		return val, nil
	case int16:
		return int32(val), nil
	case int8:
		return int32(val), nil
	case uint:
		return int32(val), nil
	case uint64:
		return int32(val), nil
	case uint32:
		return int32(val), nil
	case uint16:
		return int32(val), nil
	case uint8:
		return int32(val), nil
	case float64:
		return int32(val), nil
	case float32:
		return int32(val), nil
	case bool:
		if val {
			return int32(1), nil
		}

		return int32(0), nil
	case string:
		v, err := strconv.ParseInt(val, 0, 32)
		if err == nil {
			return int32(v), nil
		}

		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToInt32, i, i)
	case []byte:
		var v int32

		buf := bytes.NewReader(val)
		if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
			return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToInt32, i, i)
		}

		return v, nil
	case json.Number:
		return ToInt32(string(val))
	default:
		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToInt32, i, i)
	}
}

func ToInt16(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil:
		return nil, nil
	case int:
		return int16(val), nil
	case int64:
		return int16(val), nil
	case int32:
		return int16(val), nil
	case int16:
		return val, nil
	case int8:
		return int16(val), nil
	case uint:
		return int16(val), nil
	case uint64:
		return int16(val), nil
	case uint32:
		return int16(val), nil
	case uint16:
		return int16(val), nil
	case uint8:
		return int16(val), nil
	case float64:
		return int16(val), nil
	case float32:
		return int16(val), nil
	case bool:
		if val {
			return int16(1), nil
		}

		return int16(0), nil
	case string:
		v, err := strconv.ParseInt(val, 0, 16)
		if err == nil {
			return int16(v), nil
		}

		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToInt16, i, i)
	case []byte:
		var v int16

		buf := bytes.NewReader(val)
		if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
			return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToInt16, i, i)
		}

		return v, nil
	case json.Number:
		return ToInt16(string(val))
	default:
		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToInt16, i, i)
	}
}

func ToInt8(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil:
		return nil, nil
	case int:
		return int8(val), nil
	case int64:
		return int8(val), nil
	case int32:
		return int8(val), nil
	case int16:
		return int8(val), nil
	case int8:
		return val, nil
	case uint:
		return int8(val), nil
	case uint64:
		return int8(val), nil
	case uint32:
		return int8(val), nil
	case uint16:
		return int8(val), nil
	case uint8:
		return int8(val), nil
	case float64:
		return int8(val), nil
	case float32:
		return int8(val), nil
	case bool:
		if val {
			return int8(1), nil
		}

		return int8(0), nil
	case string:
		v, err := strconv.ParseInt(val, 0, 8)
		if err == nil {
			return int8(v), nil
		}

		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToInt8, i, i)
	case []byte:
		var v int8

		buf := bytes.NewReader(val)
		if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
			return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToInt8, i, i)
		}

		return v, nil
	case json.Number:
		return ToInt8(string(val))
	default:
		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToInt8, i, i)
	}
}
