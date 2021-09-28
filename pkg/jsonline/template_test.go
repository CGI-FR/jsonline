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
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/adrienaury/go-template/pkg/jsonline"
)

func TestTemplate(t *testing.T) {
	row :=
		jsonline.NewRow().
			Set("address", jsonline.NewValueString("123 Main Street, New York, NY 10030")).
			Set("last-update", jsonline.NewValueDateTime(time.Now()))

	template := jsonline.NewTemplate().WithString("name").WithNumeric("age").WithDateTime("birthdate")
	person1 := template.Create([]interface{}{"Dorothy", 30, time.Date(1991, time.September, 24, 21, 21, 0, 0, time.UTC)})
	person1.Set("house", row)
	fmt.Println(person1)

	person3 := template.Create(
		map[string]interface{}{
			"name":      "Alice",
			"age":       17,
			"birthdate": time.Date(2004, time.June, 15, 21, 8, 47, 0, time.UTC),
			"extra":     true,
		})
	fmt.Println(person3)
}

func TestTemplate2(t *testing.T) {
	template := jsonline.NewTemplate().WithNumeric("age")
	fmt.Println(template.Create(map[string]interface{}{"age": "5"}))

	for importer := template.GetImporter(os.Stdin); importer.Import(); {
		row, err := importer.GetRow()
		if err != nil {
			fmt.Println("an error occurred!", err)
		} else {
			fmt.Println(row)
		}
	}
}
