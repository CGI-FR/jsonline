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

	"github.com/cgi-fr/jsonline/pkg/jsonline"
)

func main() {
	value := jsonline.NewValueString("123 Main Street, New York, NY 10030")

	// Will print : 123 Main Street, New York, NY 10030
	fmt.Println(value.String())

	value = jsonline.NewValueNumeric(123)

	// Will print : 123
	fmt.Println(value.String())
}
