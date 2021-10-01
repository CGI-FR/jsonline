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

func importFromString(val string, targetType interface{}) (interface{}, error) {
	switch targetType.(type) {
	case nil:
		return val, nil
	case string:
		return val, nil
	case int:
		return strconv.ParseInt(val, conversionBase, conversionSize) //nolint
	default:
		return nil, fmt.Errorf("%w", ErrUnsupportedImportType)
	}
}

func importFromNumeric(val interface{}, targetType interface{}) (interface{}, error) {
	switch targetType.(type) {
	case nil:
		return val, nil
	case string:
		return val, nil
	default:
		return nil, fmt.Errorf("%w", ErrUnsupportedImportType)
	}
}

func importFromBoolean(val interface{}, targetType interface{}) (interface{}, error) {
	switch targetType.(type) {
	case nil:
		return val, nil
	case string:
		return val, nil
	default:
		return nil, fmt.Errorf("%w", ErrUnsupportedImportType)
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
		return nil, fmt.Errorf("%w", ErrUnsupportedImportType)
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
		return nil, fmt.Errorf("%w", ErrUnsupportedImportType)
	}
}

func importFromTime(val string, targetType interface{}) (interface{}, error) {
	switch targetType.(type) {
	case nil, time.Time:
		res, err := time.Parse("15:04:05Z07:00", val)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return res, nil
	default:
		return nil, fmt.Errorf("%w", ErrUnsupportedImportType)
	}
}

func importFromTimestamp(val interface{}, targetType interface{}) (interface{}, error) {
	switch targetType.(type) {
	case nil, time.Time:
		return time.Unix(convertNumericToInt64(val), 0), nil
	case string:
		return val, nil
	default:
		return nil, fmt.Errorf("%w", ErrUnsupportedImportType)
	}
}

//nolint:cyclop
func convertNumericToInt64(val interface{}) int64 {
	if v, ok := val.(int64); ok {
		return v
	}

	if v, ok := val.(int32); ok {
		return int64(v)
	}

	if v, ok := val.(int16); ok {
		return int64(v)
	}

	if v, ok := val.(int8); ok {
		return int64(v)
	}

	if v, ok := val.(int); ok {
		return int64(v)
	}

	if v, ok := val.(uint64); ok {
		return int64(v)
	}

	if v, ok := val.(uint32); ok {
		return int64(v)
	}

	if v, ok := val.(uint16); ok {
		return int64(v)
	}

	if v, ok := val.(uint8); ok {
		return int64(v)
	}

	if v, ok := val.(uint); ok {
		return int64(v)
	}

	return 0
}
