// Copyright (C) 2021 CGI France
//
// This file is part of JL.
//
// JL is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// JL is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with JL.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"time"

	"github.com/cgi-fr/jsonline/pkg/jsonline"
)

func main() {
	template := jsonline.NewTemplate().WithString("name").WithNumeric("age").WithDateTime("birthdate")

	person1, err := template.CreateRow([]interface{}{"Dorothy", 30, time.Date(1991, time.September, 24, 21, 21, 0, 0, time.UTC)})
	if err != nil {
		fmt.Println(person1) // {"name":"Dorothy","age":30,"birthdate":"1991-09-24T21:21:00Z"}
	} else {
		fmt.Println("ERROR:", err)
	}

	row := jsonline.NewRow()
	row.Set("address", jsonline.NewValueString("123 Main Street, New York, NY 10030"))
	row.Set("last-update", jsonline.NewValueDateTime(time.Now()))

	template = template.WithRow("house", jsonline.NewTemplate().WithString("address").WithDateTime("last-update"))
	person1.Set("house", row)
	fmt.Println(person1) // {"name":"Dorothy","age":30,"birthdate":"1991-09-24T21:21:00Z","house":{"address":"123 Main Street, New York, NY 10030","last-update":"2021-09-25T09:22:54+02:00"}}

	b, err := person1.MarshalJSON()
	fmt.Println(string(b)) // same result as fmt.Println(person1)

	person2 := jsonline.NewRow().UnmarshalJSON(b)
	fmt.Println(person2) // same result as fmt.Println(person1)

	person3, err := template.CreateRow(map[string]interface{}{"name": "Alice", "extra": true, "age": 17, "birthdate": time.Date(2004, time.June, 15, 21, 8, 47, 0, time.UTC)})
	if err != nil {
		fmt.Println(person3) // {"name":"Alice","age":17,"birthdate":"2004-06-15T21:08:47Z","extra":true}
	} else {
		fmt.Println("ERROR:", err)
	}
}
