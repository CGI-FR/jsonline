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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/cgi-fr/jsonline/pkg/jsonline"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

const (
	String    = "string"
	Numeric   = "numeric"
	Boolean   = "boolean"
	Binary    = "binary"
	Date      = "date"
	DateTime  = "datetime"
	Timestamp = "timestamp"
	Auto      = "auto"
	Hidden    = "hidden"
)

type RowDefinition struct {
	Columns []ColumnDefinition `yaml:"columns"`
}

type ColumnDefinition struct {
	Name    string             `yaml:"name"`
	Input   string             `yaml:"input,omitempty"`
	Output  string             `yaml:"output,omitempty"`
	Columns []ColumnDefinition `yaml:"columns,omitempty"`
}

func ReadRowDefinition(filename string) (*RowDefinition, error) {
	def := &RowDefinition{
		Columns: []ColumnDefinition{},
	}

	if _, err := os.Stat(filename); err == nil {
		dat, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		err = yaml.Unmarshal(dat, def)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
	} else {
		log.Warn().Str("filename", filename).Msg("can't read row definition from file")
	}

	return def, nil
}

func ParseRowDefinition(filename string) (jsonline.Template, jsonline.Template, error) {
	def, err := ReadRowDefinition(filename)
	if err != nil {
		return nil, nil, err
	}

	ti, to, err := parse(jsonline.NewTemplate(), jsonline.NewTemplate(), def.Columns)
	if err != nil {
		return nil, nil, err
	}

	return ti, to, nil
}

func parseDescriptor(def string) (jsonline.Format, jsonline.RawType) {
	var rawtype jsonline.RawType

	r := regexp.MustCompile(`^([^\(]+)(?:\(([^\)]+)\))?$`) // "<FORMAT>(<TYPE>)" or "FORMAT"
	submatches := r.FindStringSubmatch(def)
	format := jsonline.Auto

	if len(submatches) > 1 {
		if f, ok := formatRegistry[submatches[1]]; ok {
			format = f
		}
	}

	//nolint:gomnd
	if len(submatches) > 2 {
		rawtype = typeRegistry[submatches[2]]
	}

	return format, rawtype
}

func parse(ti jsonline.Template, to jsonline.Template,
	columns []ColumnDefinition) (jsonline.Template, jsonline.Template, error) {
	for _, column := range columns {
		iformat, irawtype := parseDescriptor(column.Input)
		oformat, orawtype := parseDescriptor(column.Output)

		ti.With(column.Name, iformat, irawtype)
		to.With(column.Name, oformat, orawtype)

		if len(column.Columns) > 0 {
			rowti, rowto, err := parse(jsonline.NewTemplate(), jsonline.NewTemplate(), column.Columns)
			if err != nil {
				return ti, to, err
			}

			ti = ti.WithRow(column.Name, rowti)
			to = to.WithRow(column.Name, rowto)
		}
	}

	return ti, to, nil
}

func createTemplateFromString(input string) (jsonline.Template, jsonline.Template, error) {
	row := jsonline.NewRow()

	if err := json.Unmarshal([]byte(input), row); err != nil {
		return nil, nil, fmt.Errorf("%w", err)
	}

	return createTemplateFromRow(row)
}

func createTemplateFromRow(row jsonline.Row) (jsonline.Template, jsonline.Template, error) {
	ti := jsonline.NewTemplate()
	to := jsonline.NewTemplate()

	iter := row.IterValues()

	for colname, v, ok := iter(); ok; colname, v, ok = iter() {
		valExported, err := v.Export()
		if err != nil {
			return ti, to, fmt.Errorf("%w", err)
		}

		switch coldef := valExported.(type) {
		case string:
			parts := strings.SplitN(coldef, ":", 2) //nolint:gomnd
			iformat, irawtype := parseDescriptor(parts[0])
			ti.With(colname, iformat, irawtype)

			if len(parts) > 1 {
				oformat, orawtype := parseDescriptor(parts[1])
				to.With(colname, oformat, orawtype)
			} else {
				to.With(colname, iformat, irawtype)
			}

		case jsonline.Row:
			rowti, rowto, err := createTemplateFromRow(coldef)
			if err != nil {
				return ti, to, err
			}

			ti = ti.WithRow(colname, rowti)
			to = to.WithRow(colname, rowto)
		}
	}

	return ti, to, nil
}

//nolint:gochecknoglobals
var typeRegistry = map[string]jsonline.RawType{
	"int":         int(0),
	"int64":       int64(0),
	"int32":       int32(0),
	"int16":       int16(0),
	"int8":        int8(0),
	"uint":        uint(0),
	"uint64":      uint64(0),
	"uint32":      uint32(0),
	"uint16":      uint16(0),
	"uint8":       uint8(0),
	"float64":     float64(0),
	"float32":     float32(0),
	"bool":        bool(true),
	"byte":        byte(0),
	"rune":        rune(0),
	"string":      string(""),
	"[]byte":      []byte{},
	"time.Time":   time.Time{},
	"json.Number": json.Number(""),
}

//nolint:gochecknoglobals
var formatRegistry = map[string]jsonline.Format{
	String:    jsonline.String,
	Numeric:   jsonline.Numeric,
	Boolean:   jsonline.Boolean,
	Binary:    jsonline.Binary,
	Date:      jsonline.Date,
	DateTime:  jsonline.DateTime,
	Timestamp: jsonline.Timestamp,
	Auto:      jsonline.Auto,
	Hidden:    jsonline.Hidden,
}
