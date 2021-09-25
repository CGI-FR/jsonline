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
