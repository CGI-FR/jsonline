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

	"github.com/adrienaury/go-template/pkg/jsonline"
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
	Columns []ColumnDefinition `yaml:"columns"`
}

type ColumnDefinition struct {
	Name    string             `yaml:"name"`
	Type    string             `yaml:"type"`
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

func ParseRowDefinition(filename string) (jsonline.Template, error) {
	def, err := ReadRowDefinition(filename)
	if err != nil {
		return nil, err
	}

	template := jsonline.NewTemplate()

	return parse(template, def.Columns)
}

//nolint:cyclop
func parse(tmpl jsonline.Template, columns []ColumnDefinition) (jsonline.Template, error) {
	for _, column := range columns {
		switch column.Type {
		case String:
			tmpl = tmpl.WithString(column.Name)
		case Numeric:
			tmpl = tmpl.WithNumeric(column.Name)
		case Boolean:
			tmpl = tmpl.WithBoolean(column.Name)
		case Binary:
			tmpl = tmpl.WithBinary(column.Name)
		case DateTime:
			tmpl = tmpl.WithDateTime(column.Name)
		case Timestamp:
			tmpl = tmpl.WithTimestamp(column.Name)
		case Auto:
			tmpl = tmpl.WithAuto(column.Name)
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
	switch coltype {
	case String:
		tmpl = tmpl.WithString(colname)
	case Numeric:
		tmpl = tmpl.WithNumeric(colname)
	case Boolean:
		tmpl = tmpl.WithBoolean(colname)
	case Binary:
		tmpl = tmpl.WithBinary(colname)
	case DateTime:
		tmpl = tmpl.WithDateTime(colname)
	case Timestamp:
		tmpl = tmpl.WithTimestamp(colname)
	case Auto:
		tmpl = tmpl.WithAuto(colname)
	case Hidden:
		tmpl = tmpl.WithHidden(colname)
	}

	return tmpl
}
