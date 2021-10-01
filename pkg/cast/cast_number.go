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
	"strconv"
)

//nolint:cyclop,gomnd
func ToNumber(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil, json.Number:
		return val, nil
	case bool:
		if val {
			return json.Number("1"), nil
		}

		return json.Number("0"), nil
	case float64:
		return json.Number(strconv.FormatFloat(val, 'f', -1, 64)), nil
	case float32:
		return json.Number(strconv.FormatFloat(float64(val), 'f', -1, 32)), nil
	case int:
		return json.Number(strconv.Itoa(val)), nil
	case int64:
		return json.Number(strconv.FormatInt(val, 10)), nil
	case int32:
		return json.Number(strconv.Itoa(int(val))), nil
	case int16:
		return json.Number(strconv.FormatInt(int64(val), 10)), nil
	case int8:
		return json.Number(strconv.FormatInt(int64(val), 10)), nil
	case uint:
		return json.Number(strconv.FormatUint(uint64(val), 10)), nil
	case uint64:
		return json.Number(strconv.FormatUint(val, 10)), nil
	case uint32:
		return json.Number(strconv.FormatUint(uint64(val), 10)), nil
	case uint16:
		return json.Number(strconv.FormatUint(uint64(val), 10)), nil
	case uint8:
		return json.Number(strconv.FormatUint(uint64(val), 10)), nil
	case []byte:
		return json.Number(string(val)), nil
	case string:
		return json.Number(val), nil
	default:
		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToNumber, i, i)
	}
}
