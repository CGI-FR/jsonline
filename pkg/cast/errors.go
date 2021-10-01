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
	"errors"
)

var (
	ErrUnableToCastToInt     = errors.New("unable to cast value to int")
	ErrUnableToCastToInt64   = errors.New("unable to cast value to int64")
	ErrUnableToCastToInt32   = errors.New("unable to cast value to int32")
	ErrUnableToCastToInt16   = errors.New("unable to cast value to int16")
	ErrUnableToCastToInt8    = errors.New("unable to cast value to int8")
	ErrUnableToCastToUint    = errors.New("unable to cast value to uint")
	ErrUnableToCastToUint64  = errors.New("unable to cast value to uint64")
	ErrUnableToCastToUint32  = errors.New("unable to cast value to uint32")
	ErrUnableToCastToUint16  = errors.New("unable to cast value to uint16")
	ErrUnableToCastToUint8   = errors.New("unable to cast value to uint8")
	ErrUnableToCastToFloat64 = errors.New("unable to cast value to float64")
	ErrUnableToCastToFloat32 = errors.New("unable to cast value to float32")
	ErrUnableToCastToString  = errors.New("unable to cast value to string")
	ErrUnableToCastToNumber  = errors.New("unable to cast value to number")
	ErrUnableToCastToBool    = errors.New("unable to cast value to bool")
	ErrUnableToCastToBinary  = errors.New("unable to cast value to binary")
	ErrUnableToCastToTime    = errors.New("unable to cast value to time")
	ErrUnableToCast          = errors.New("unable to cast value")
)
