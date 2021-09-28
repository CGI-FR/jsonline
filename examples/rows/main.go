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

	"github.com/adrienaury/go-template/pkg/jsonline"
)

func main() {
	row := jsonline.NewRow()
	row.Set("address", jsonline.NewValueString("123 Main Street, New York, NY 10030"))
	row.Set("last-update", jsonline.NewValueDateTime(time.Now()))

	// Will print : {"address":"123 Main Street, New York, NY 10030","last-update":"2021-09-25T08:51:10+02:00"}
	fmt.Println(row)
}
