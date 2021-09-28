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

package jsonline_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/adrienaury/go-template/pkg/jsonline"
	"github.com/stretchr/testify/assert"
)

func TestRow(t *testing.T) {
	r := jsonline.NewRow()
	r.Set("name", jsonline.NewValueString("Dorothy"))
	r.Set("age", jsonline.NewValueNumeric(30))
	r.Set("car", jsonline.NewValueNil())
	r.Set("pet", jsonline.NewRow())
	r.Set("house", jsonline.NewRow().
		Set("address", jsonline.NewValueString("123 Main Street, New York, NY 10030")).
		Set("buy-date", jsonline.NewValueDateTime(time.Now())))
	fmt.Println(r.String())

	row :=
		jsonline.NewRow().
			Set("address", jsonline.NewValueString("123 Main Street, New York, NY 10030")).
			Set("last-update", jsonline.NewValueDateTime(time.Now()))
	fmt.Println(row)
}

func TestRow2(t *testing.T) {
	str := `{"name":"nathan","surname":"Aury","age":5}`
	row := jsonline.NewRow()
	err := json.Unmarshal([]byte(str), row)
	assert.NoError(t, err)
	fmt.Println(row)
}
