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
	"bufio"
	"encoding/json"
	"io"
)

// JSONRowIterator export rows to JSON format.
type JSONRowIterator struct {
	file     io.ReadCloser
	fscanner *bufio.Scanner
	error    error
	value    map[string]interface{}
}

// NewJSONRowIterator creates a new JSONRowIterator.
func NewJSONRowIterator(file io.ReadCloser) *JSONRowIterator {
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)   //nolint:gomnd
	scanner.Buffer(buf, 10*1024*1024) //nolint:gomnd

	return &JSONRowIterator{file, scanner, nil, nil}
}

// Close file format.
func (re *JSONRowIterator) Close() error {
	if err := re.file.Close(); err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}

// Value return current row.
func (re *JSONRowIterator) Value() map[string]interface{} {
	if re.value != nil {
		return re.value
	}

	panic("Value is not valid after iterator finished")
}

// Error return error catch by next.
func (re *JSONRowIterator) Error() error {
	return re.error
}

// Next try to convert next line to Row.
func (re *JSONRowIterator) Next() bool {
	if !re.fscanner.Scan() {
		if re.fscanner.Err() != nil {
			re.error = re.fscanner.Err()
		}

		return false
	}

	line := re.fscanner.Bytes()

	var row map[string]interface{}

	err2 := json.Unmarshal(line, &row)

	if err2 != nil {
		re.error = err2

		return false
	}

	re.value = row

	return true
}
