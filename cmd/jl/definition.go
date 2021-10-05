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
	DateTime  = "datetime"
	Timestamp = "timestamp"
	Auto      = "auto"
	Hidden    = "hidden"
	Row       = "row"
)

type RowDefinition struct {
	Input  []ColumnDefinition `yaml:"input"`
	Output []ColumnDefinition `yaml:"output"`
}

type ColumnDefinition struct {
	Name    string             `yaml:"name"`
	Format  string             `yaml:"format"`
	Type    string             `yaml:"type"`
	Columns []ColumnDefinition `yaml:"columns,omitempty"`
}

func ReadRowDefinition(filename string) (*RowDefinition, error) {
	def := &RowDefinition{
		Input:  []ColumnDefinition{},
		Output: []ColumnDefinition{},
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

	ti, err := parse(jsonline.NewTemplate(), def.Input)
	if err != nil {
		return nil, nil, err
	}

	to, err := parse(jsonline.NewTemplate(), def.Output)
	if err != nil {
		return nil, nil, err
	}

	return ti, to, nil
}

//nolint:cyclop
func parse(tmpl jsonline.Template, columns []ColumnDefinition) (jsonline.Template, error) {
	for _, column := range columns {
		coltype, ok := typeRegistry[column.Type]
		if !ok && len(column.Type) > 0 {
			return nil, fmt.Errorf("%w: %v", ErrInvalidRawType, column.Type)
		}

		switch column.Format {
		case String:
			tmpl = tmpl.WithMappedString(column.Name, coltype)
		case Numeric:
			tmpl = tmpl.WithMappedNumeric(column.Name, coltype)
		case Boolean:
			tmpl = tmpl.WithMappedBoolean(column.Name, coltype)
		case Binary:
			tmpl = tmpl.WithMappedBinary(column.Name, coltype)
		case DateTime:
			tmpl = tmpl.WithMappedDateTime(column.Name, coltype)
		case Timestamp:
			tmpl = tmpl.WithMappedTimestamp(column.Name, coltype)
		case Auto:
			tmpl = tmpl.WithMappedAuto(column.Name, coltype)
		case Hidden:
			tmpl = tmpl.WithHidden(column.Name)
		case Row:
			rowt, err := parse(jsonline.NewTemplate(), column.Columns)
			if err != nil {
				return tmpl, err
			}

			tmpl = tmpl.WithRow(column.Name, rowt)
		}
	}

	return tmpl, nil
}

func createTemplateFromString(input string) (jsonline.Template, error) {
	row := jsonline.NewRow()

	if err := json.Unmarshal([]byte(input), row); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return createTemplateFromRow(row)
}

func createTemplateFromRow(row jsonline.Row) (jsonline.Template, error) {
	tmpl := jsonline.NewTemplate()

	iter := row.Iter()

	for colname, v, ok := iter(); ok; colname, v, ok = iter() {
		valExported, err := v.Export()
		if err != nil {
			return tmpl, fmt.Errorf("%w", err)
		}

		switch coltype := valExported.(type) {
		case string:
			tmpl = setColumnInTemplate(tmpl, coltype, colname)

		case jsonline.Row:
			rowt, err := createTemplateFromRow(coltype)
			if err != nil {
				return tmpl, err
			}

			tmpl = tmpl.WithRow(colname, rowt)
		}
	}

	return tmpl, nil
}

func setColumnInTemplate(tmpl jsonline.Template, coltype string, colname string) jsonline.Template {
	format := strings.SplitN(coltype, ":", 2) //nolint:gomnd

	var colrawtype interface{}
	if len(format) > 1 {
		colrawtype = typeRegistry[format[1]]
	}

	switch format[0] {
	case String:
		tmpl = tmpl.WithMappedString(colname, colrawtype)
	case Numeric:
		tmpl = tmpl.WithMappedNumeric(colname, colrawtype)
	case Boolean:
		tmpl = tmpl.WithMappedBoolean(colname, colrawtype)
	case Binary:
		tmpl = tmpl.WithMappedBinary(colname, colrawtype)
	case DateTime:
		tmpl = tmpl.WithMappedDateTime(colname, colrawtype)
	case Timestamp:
		tmpl = tmpl.WithMappedTimestamp(colname, colrawtype)
	case Auto:
		tmpl = tmpl.WithMappedAuto(colname, colrawtype)
	case Hidden:
		tmpl = tmpl.WithHidden(colname)
	}

	return tmpl
}

//nolint:gochecknoglobals
var typeRegistry = map[string]interface{}{
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
