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

//nolint:cyclop,funlen
package cast

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func ToBinary(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil, []byte:
		return val, nil
	case float64:
		return float64ToBytes(val), nil
	case float32:
		return float32ToBytes(val), nil
	case int:
		return intToBytes(val), nil
	case int64:
		return int64ToBytes(val), nil
	case int32:
		return int32ToBytes(val), nil
	case int16:
		return int16ToBytes(val), nil
	case int8:
		return int8ToBytes(val), nil
	case uint:
		return uintToBytes(val), nil
	case uint64:
		return uint64ToBytes(val), nil
	case uint32:
		return uint32ToBytes(val), nil
	case uint16:
		return uint16ToBytes(val), nil
	case uint8:
		return uint8ToBytes(val), nil
	case bool:
		return boolToBytes(val), nil
	case string:
		return []byte(val), nil
	case json.Number:
		return []byte(val), nil
	case time.Time:
		return ToBinary(val.Unix())
	default:
		v := reflect.ValueOf(val)
		switch v.Kind() { //nolint:exhaustive
		case reflect.Array:
			if v.Type().Elem().Kind() == reflect.Uint8 {
				b := make([]uint8, v.Len())
				if n := reflect.Copy(reflect.ValueOf(b), v); n != v.Len() {
					return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToBinary, i, i)
				}

				return b, nil
			}

			return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToBinary, i, i)
		default:
			return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToBinary, i, i)
		}
	}
}
