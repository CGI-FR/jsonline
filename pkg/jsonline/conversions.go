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
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
)

const (
	conversionSize = 64
)

func toString(val interface{}) interface{} {
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

func toNumeric(val interface{}) interface{} {
	switch typedValue := val.(type) {
	case nil:
		return nil
	case int64, int, int16, int8, byte, rune, float64, float32, uint, uint16, uint32, uint64:
		return typedValue
	case bool:
		if typedValue {
			return 1
		}

		return 0
	case string:
		r, err := strconv.ParseFloat(typedValue, conversionSize)
		if err == nil {
			return r
		}

		return fmt.Sprintf("ERROR: %v", err)
	case []byte:
		return toNumeric(fmt.Sprintf("%v", string(typedValue)))
	default:
		return toNumeric(fmt.Sprintf("%v", val))
	}
}

//nolint:funlen,cyclop
func toBoolean(val interface{}) interface{} {
	switch typedValue := val.(type) {
	case nil:
		return nil
	case int:
		return typedValue != 0
	case int8:
		return typedValue != 0
	case int16:
		return typedValue != 0
	case int32: // rune=int32
		return typedValue != 0
	case int64:
		return typedValue != 0
	case uint:
		return typedValue != 0
	case uint8: // byte=uint8
		return typedValue != 0
	case uint16:
		return typedValue != 0
	case uint32:
		return typedValue != 0
	case uint64:
		return typedValue != 0
	case float64:
		return typedValue != 0.0
	case float32:
		return typedValue != 0.0
	case complex64:
		return typedValue != 0i+0
	case complex128:
		return typedValue != 0i+0
	case bool:
		return typedValue
	case string:
		r, err := strconv.ParseBool(typedValue)
		if err == nil {
			return r
		}

		f64, err := strconv.ParseFloat(typedValue, conversionSize)
		if err == nil {
			return toBoolean(f64)
		}

		return fmt.Sprintf("ERROR: %v", err)
	case []byte:
		return toBoolean(fmt.Sprintf("%v", string(typedValue)))
	default:
		return toBoolean(fmt.Sprintf("%v", val))
	}
}

func toBinary(val interface{}) interface{} {
	switch typedValue := val.(type) {
	case nil:
		return nil
	case []byte:
		return base64.StdEncoding.EncodeToString(typedValue)
	case string:
		return base64.StdEncoding.EncodeToString([]byte(typedValue))
	default:
		return base64.StdEncoding.EncodeToString([]byte(toString(typedValue).(string)))
	}
}

//nolint:cyclop
func toDateTime(val interface{}) interface{} {
	switch typedValue := val.(type) {
	case nil:
		return nil
	case int64:
		return toDateTime(time.Unix(typedValue, 0))
	case int32:
		return toDateTime(time.Unix(int64(typedValue), 0))
	case int16:
		return toDateTime(time.Unix(int64(typedValue), 0))
	case int8:
		return toDateTime(time.Unix(int64(typedValue), 0))
	case int:
		return toDateTime(time.Unix(int64(typedValue), 0))
	case uint64:
		return toDateTime(time.Unix(int64(typedValue), 0))
	case uint32:
		return toDateTime(time.Unix(int64(typedValue), 0))
	case uint16:
		return toDateTime(time.Unix(int64(typedValue), 0))
	case uint8:
		return toDateTime(time.Unix(int64(typedValue), 0))
	case uint:
		return toDateTime(time.Unix(int64(typedValue), 0))
	case float32:
		return toDateTime(time.Unix(int64(typedValue), 0))
	case float64:
		return toDateTime(time.Unix(int64(typedValue), 0))
	case time.Time:
		return typedValue.Format(time.RFC3339)
	case string:
		t, err := time.Parse(time.RFC3339, typedValue)
		if err != nil {
			return fmt.Sprintf("ERROR: %v", err)
		}

		return toDateTime(t)
	default:
		return toDateTime(toString(val))
	}
}

//nolint:cyclop
func toTime(val interface{}) interface{} {
	switch typedValue := val.(type) {
	case nil:
		return nil
	case int64:
		return toTime(time.Unix(typedValue, 0))
	case int32:
		return toTime(time.Unix(int64(typedValue), 0))
	case int16:
		return toTime(time.Unix(int64(typedValue), 0))
	case int8:
		return toTime(time.Unix(int64(typedValue), 0))
	case int:
		return toTime(time.Unix(int64(typedValue), 0))
	case uint64:
		return toTime(time.Unix(int64(typedValue), 0))
	case uint32:
		return toTime(time.Unix(int64(typedValue), 0))
	case uint16:
		return toTime(time.Unix(int64(typedValue), 0))
	case uint8:
		return toTime(time.Unix(int64(typedValue), 0))
	case uint:
		return toTime(time.Unix(int64(typedValue), 0))
	case float32:
		return toTime(time.Unix(int64(typedValue), 0))
	case float64:
		return toTime(time.Unix(int64(typedValue), 0))
	case time.Time:
		return typedValue.Format("15:04:05Z07:00")
	case string:
		t, err := time.Parse("15:04:05Z07:00", typedValue)
		if err != nil {
			return fmt.Sprintf("ERROR: %v", err)
		}

		return toTime(t)
	default:
		return toTime(toString(val))
	}
}

func toTimeStamp(val interface{}) interface{} {
	switch typedValue := val.(type) {
	case nil:
		return nil
	case time.Time:
		return typedValue.Unix()
	case int64:
		return typedValue
	case string:
		t, err := time.Parse(time.RFC3339, typedValue)
		if err != nil {
			return fmt.Sprintf("ERROR: %v", err)
		}

		return toDateTime(t)
	default:
		return toDateTime(toString(val))
	}
}
