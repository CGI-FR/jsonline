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

//nolint:cyclop,varnamelen
package cast

import (
	"fmt"
	"time"
)

func ToDate(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil:
		return val, nil

	case time.Time:
		return val.Format("2006-01-02"), nil

	case string:
		_, err := time.Parse("2006-01-02", val)
		if err != nil {
			i64, err := ToInt64(val)
			if err != nil {
				return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToDate, i, i)
			}

			return ToDate(i64)
		}

		return val, nil

	case []byte:
		t, err := ToDate(string(val))
		if err != nil {
			i64, err := ToInt64(val)
			if err != nil {
				return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToDate, i, i)
			}

			return ToDate(i64)
		}

		return t, nil

	case int64:
		return time.Unix(val, 0).Format("2006-01-02"), nil

	default:
		s, err := ToString(val)
		if err != nil {
			return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToTime, i, i)
		}

		return ToDate(s)
	}
}

func ToTime(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil, time.Time:
		return val, nil

	case string:
		t, err := time.Parse(TimeStringFormat, val)
		if err != nil {
			i64, err := ToInt64(val)
			if err != nil {
				return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToTime, i, i)
			}

			return ToTime(i64)
		}

		return t, nil

	case []byte:
		t, err := ToTime(string(val))
		if err != nil {
			i64, err := ToInt64(val)
			if err != nil {
				return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToTime, i, i)
			}

			return ToTime(i64)
		}

		return t, nil

	case int64:
		return time.Unix(val, 0), nil

	default:
		i64, err := ToInt64(val)
		if err != nil {
			return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToTime, i, i)
		}

		return ToTime(i64)
	}
}

func ToTimestamp(i interface{}) (interface{}, error) {
	switch val := i.(type) {
	case nil, int64:
		return val, nil
	case string:
		t, err := time.Parse(TimeStringFormat, val)
		if err == nil {
			return t.Unix(), nil
		}

		return nil, fmt.Errorf("%w: %#v (%T)", ErrUnableToCastToTime, i, i)
	case []byte:
		return ToTimestamp(string(val))
	case time.Time:
		return val.Unix(), nil
	default:
		return ToInt64(val)
	}
}
