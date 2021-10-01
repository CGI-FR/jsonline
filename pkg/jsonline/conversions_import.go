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

	"github.com/adrienaury/go-template/pkg/cast"
)

func importFromString(val string, targetType interface{}) (interface{}, error) {
	switch targetType.(type) {
	case nil:
		return val, nil
	case string:
		return val, nil
	case []byte:
		return []byte(val), nil
	case int:
		return strconv.ParseInt(val, conversionBase, conversionSize) //nolint
	default:
		return nil, fmt.Errorf("%w: %T", ErrUnsupportedImportType, targetType)
	}
}

//nolint:cyclop,wrapcheck
func importFromNumeric(val interface{}, targetType interface{}) (interface{}, error) {
	switch targetType.(type) {
	case int:
		return cast.ToInt(val)
	case int64:
		return cast.ToInt64(val)
	case int32:
		return cast.ToInt32(val)
	case int16:
		return cast.ToInt16(val)
	case int8:
		return cast.ToInt8(val)
	case uint:
		return cast.ToUint(val)
	case uint64:
		return cast.ToUint64(val)
	case uint32:
		return cast.ToUint32(val)
	case uint16:
		return cast.ToUint16(val)
	case uint8:
		return cast.ToUint8(val)
	case nil, float64:
		return cast.ToFloat64(val)
	case float32:
		return cast.ToFloat32(val)
	case bool:
		return cast.ToBool(val)
	case string:
		return cast.ToString(val)
	case []byte:
		return cast.ToBinary(val)
	case time.Time:
		return cast.ToTime(val)
	default:
		return nil, fmt.Errorf("%w: %T", ErrUnsupportedImportType, targetType)
	}
}

func importFromBoolean(val interface{}, targetType interface{}) (interface{}, error) {
	switch targetType.(type) {
	case nil:
		return val, nil
	case string:
		return val, nil
	default:
		return nil, fmt.Errorf("%w: %T", ErrUnsupportedImportType, targetType)
	}
}

func importFromBinary(val string, targetType interface{}) (interface{}, error) {
	switch targetType.(type) {
	case string:
		b, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return string(b), nil
	case nil, []byte:
		b, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return b, nil
	default:
		return nil, fmt.Errorf("%w: %T", ErrUnsupportedImportType, targetType)
	}
}

func importFromDateTime(val string, targetType interface{}) (interface{}, error) {
	switch targetType.(type) {
	case nil, time.Time:
		res, err := time.Parse(time.RFC3339, val)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return res, nil
	default:
		return nil, fmt.Errorf("%w: %T", ErrUnsupportedImportType, targetType)
	}
}

//nolint:cyclop,wrapcheck
func importFromTimestamp(val interface{}, targetType interface{}) (interface{}, error) {
	switch targetType.(type) {
	case int:
		return cast.ToInt(val)
	case nil, int64:
		return cast.ToInt64(val)
	case int32:
		return cast.ToInt32(val)
	case int16:
		return cast.ToInt16(val)
	case int8:
		return cast.ToInt8(val)
	case uint:
		return cast.ToUint(val)
	case uint64:
		return cast.ToUint64(val)
	case uint32:
		return cast.ToUint32(val)
	case uint16:
		return cast.ToUint16(val)
	case uint8:
		return cast.ToUint8(val)
	case float64:
		return cast.ToFloat64(val)
	case float32:
		return cast.ToFloat32(val)
	case bool:
		return cast.ToBool(val)
	case string:
		return cast.ToString(val)
	case []byte:
		return cast.ToBinary(val)
	case time.Time:
		i64, err := cast.ToInt64(val)
		if err != nil {
			return nil, err
		}

		return time.Unix(i64.(int64), 0), nil
	default:
		return nil, fmt.Errorf("%w: %T", ErrUnsupportedImportType, targetType)
	}
}
