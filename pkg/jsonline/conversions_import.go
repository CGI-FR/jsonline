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
	"time"
)

func importToBinary(val interface{}) (interface{}, error) {
	if val == nil {
		return nil, nil
	}

	b, err := base64.StdEncoding.DecodeString(convertToString(val).(string))
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return b, nil
}

func importToDateTime(val interface{}) (interface{}, error) {
	switch typedValue := val.(type) {
	case nil:
		return nil, nil
	case string:
		res, err := time.Parse(time.RFC3339, typedValue)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return res, nil
	default:
		return importToDateTime(convertToString(val))
	}
}

func importToTime(val interface{}) (interface{}, error) {
	switch typedValue := val.(type) {
	case nil:
		return nil, nil
	case string:
		res, err := time.Parse("15:04:05Z07:00", typedValue)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return res, nil
	default:
		return importToTime(convertToString(val))
	}
}

//nolint:cyclop
func importToTimestamp(val interface{}) (interface{}, error) {
	switch typedValue := val.(type) {
	case nil:
		return nil, nil
	case int64:
		return typedValue, nil
	case int32:
		return int64(typedValue), nil
	case int16:
		return int64(typedValue), nil
	case int8:
		return int64(typedValue), nil
	case int:
		return int64(typedValue), nil
	case uint64:
		return int64(typedValue), nil
	case uint32:
		return int64(typedValue), nil
	case uint16:
		return int64(typedValue), nil
	case uint8:
		return int64(typedValue), nil
	case uint:
		return int64(typedValue), nil
	default:
		if res, err := convertToNumeric(typedValue); err == nil {
			return res, nil
		}

		t, err := importToDateTime(typedValue)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return t.(time.Time).Unix(), nil
	}
}
