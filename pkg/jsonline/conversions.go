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

package jsonline

import (
	"fmt"
	"strconv"
)

const (
	conversionBase = 10
	conversionSize = 64
)

func convertToString(val interface{}) interface{} {
	switch typedValue := val.(type) {
	case nil:
		return nil
	case rune:
		return string(typedValue)
	case []byte:
		return string(typedValue)
	default:
		return fmt.Sprintf("%v", val)
	}
}

func convertToNumeric(val interface{}) (interface{}, error) {
	switch typedValue := val.(type) {
	case nil:
		return nil, nil
	case int64, int, int16, int8, byte, rune, float64, float32, uint, uint16, uint32, uint64:
		return typedValue, nil
	case bool:
		if typedValue {
			return 1, nil
		}

		return 0, nil
	case string:
		r, err := strconv.ParseFloat(typedValue, conversionSize)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return r, nil
	case []byte:
		return convertToNumeric(fmt.Sprintf("%v", string(typedValue)))
	default:
		return convertToNumeric(fmt.Sprintf("%v", val))
	}
}

//nolint:funlen,cyclop
func convertToBoolean(val interface{}) (interface{}, error) {
	switch typedValue := val.(type) {
	case nil:
		return nil, nil
	case int:
		return typedValue != 0, nil
	case int8:
		return typedValue != 0, nil
	case int16:
		return typedValue != 0, nil
	case int32: // rune=int32
		return typedValue != 0, nil
	case int64:
		return typedValue != 0, nil
	case uint:
		return typedValue != 0, nil
	case uint8: // byte=uint8
		return typedValue != 0, nil
	case uint16:
		return typedValue != 0, nil
	case uint32:
		return typedValue != 0, nil
	case uint64:
		return typedValue != 0, nil
	case float64:
		return typedValue != 0.0, nil
	case float32:
		return typedValue != 0.0, nil
	case complex64:
		return typedValue != 0i+0, nil
	case complex128:
		return typedValue != 0i+0, nil
	case bool:
		return typedValue, nil
	case string:
		r, err := strconv.ParseBool(typedValue)
		if err == nil {
			return r, nil
		}

		f64, err := strconv.ParseFloat(typedValue, conversionSize)
		if err == nil {
			return convertToBoolean(f64)
		}

		return nil, fmt.Errorf("%w", err)
	case []byte:
		return convertToBoolean(fmt.Sprintf("%v", string(typedValue)))
	default:
		return convertToBoolean(fmt.Sprintf("%v", val))
	}
}
