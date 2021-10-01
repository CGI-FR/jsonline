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
	"encoding/json"
	"fmt"
	"time"
)

//nolint:cyclop
func To(targetType interface{}, val interface{}) (interface{}, error) {
	switch targetType.(type) {
	case int:
		return ToInt(val)
	case int64:
		return ToInt64(val)
	case int32:
		return ToInt32(val)
	case int16:
		return ToInt16(val)
	case int8:
		return ToInt8(val)
	case uint:
		return ToUint(val)
	case uint64:
		return ToUint64(val)
	case uint32:
		return ToUint32(val)
	case uint16:
		return ToUint16(val)
	case uint8:
		return ToUint8(val)
	case float64:
		return ToFloat64(val)
	case float32:
		return ToFloat32(val)
	case bool:
		return ToBool(val)
	case string:
		return ToString(val)
	case []byte:
		return ToBinary(val)
	case time.Time:
		return ToTime(val)
	case json.Number:
		n, err := ToString(val)
		if err != nil {
			return nil, err
		}

		return json.Number(n.(string)), err
	default:
		return nil, fmt.Errorf("%w: %#v to %T", ErrUnableToCast, val, targetType)
	}
}