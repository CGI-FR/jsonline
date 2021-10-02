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

	"github.com/cgi-fr/jsonline/pkg/cast"
)

//nolint:wrapcheck
func importFromString(val interface{}, targetType interface{}) (interface{}, error) {
	switch targetType.(type) {
	case nil:
		return cast.ToString(val)
	default:
		return cast.To(targetType, val)
	}
}

//nolint:wrapcheck
func importFromNumeric(val interface{}, targetType interface{}) (interface{}, error) {
	switch targetType.(type) {
	case nil:
		return cast.ToNumber(val)
	default:
		return cast.To(targetType, val)
	}
}

//nolint:wrapcheck
func importFromBoolean(val interface{}, targetType interface{}) (interface{}, error) {
	switch targetType.(type) {
	case nil:
		return cast.ToBool(val)
	default:
		return cast.To(targetType, val)
	}
}

//nolint:wrapcheck
func importFromBinary(val interface{}, targetType interface{}) (interface{}, error) {
	str, err := cast.ToString(val)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	b, err := base64.StdEncoding.DecodeString(str.(string))
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	switch targetType.(type) {
	case nil:
		return b, nil
	default:
		return cast.To(targetType, b)
	}
}

//nolint:wrapcheck
func importFromDateTime(val interface{}, targetType interface{}) (interface{}, error) {
	switch targetType.(type) {
	case nil:
		return cast.ToTime(val)
	default:
		return cast.To(targetType, val)
	}
}

//nolint:wrapcheck
func importFromTimestamp(val interface{}, targetType interface{}) (interface{}, error) {
	switch targetType.(type) {
	case nil:
		return cast.ToInt64(val)
	default:
		return cast.To(targetType, val)
	}
}
