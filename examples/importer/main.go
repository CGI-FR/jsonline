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
	"os"

	"github.com/adrienaury/go-template/pkg/jsonline"
)

func main() {
	template := jsonline.NewTemplate().WithString("name").WithNumeric("age").WithDateTime("birthdate")

	for importer := template.GetImporter(os.Stdin); importer.Import(); { // or importer := jsonline.NewImporter(os.Stdin).WithTemplate(template)
		row, err := importer.GetRow()
		if err != nil {
			fmt.Println("an error occurred!", err)
		} else {
			fmt.Println(row)
		}
	}
}
