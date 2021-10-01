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
func ToFloat64(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil:
		return nil, nil
	case int:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int16:
		return float64(val), nil
	case int8:
		return float64(val), nil
	case uint:
		return float64(val), nil
	case uint64:
		return float64(val), nil
	case uint32:
		return float64(val), nil
	case uint16:
		return float64(val), nil
	case uint8:
		return float64(val), nil
	case float64:
		return val, nil
	case float32:
		return float64(val), nil
	case bool:
		if val {
			return float64(1), nil
		}

		return float64(0), nil
	case string:
		v, err := strconv.ParseFloat(val, 64) //nolint:gomnd
		if err == nil {
			return v, nil
		}

		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToFloat64, i, i)
	case []byte:
		var f float64

		buf := bytes.NewReader(val)
		if err := binary.Read(buf, binary.LittleEndian, &f); err != nil {
			return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToFloat64, i, i)
		}

		return f, nil
	case json.Number:
		return ToFloat64(string(val))
	default:
		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToFloat64, i, i)
	}
}

//nolint:cyclop,funlen
func ToFloat32(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil:
		return nil, nil
	case int:
		return float32(val), nil
	case int64:
		return float32(val), nil
	case int32:
		return float32(val), nil
	case int16:
		return float32(val), nil
	case int8:
		return float32(val), nil
	case uint:
		return float32(val), nil
	case uint64:
		return float32(val), nil
	case uint32:
		return float32(val), nil
	case uint16:
		return float32(val), nil
	case uint8:
		return float32(val), nil
	case float64:
		return float32(val), nil
	case float32:
		return val, nil
	case bool:
		if val {
			return float32(1), nil
		}

		return float32(0), nil
	case string:
		v, err := strconv.ParseFloat(val, 32) //nolint:gomnd
		if err == nil {
			return float32(v), nil
		}

		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToFloat32, i, i)
	case []byte:
		var f float32

		buf := bytes.NewReader(val)
		if err := binary.Read(buf, binary.LittleEndian, &f); err != nil {
			return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToFloat32, i, i)
		}

		return f, nil
	case json.Number:
		return ToFloat32(string(val))
	default:
		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToFloat32, i, i)
	}
}
