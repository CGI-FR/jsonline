package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/adrienaury/go-template/pkg/jsonline"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
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

	return parse(template, def)
}

//nolint:cyclop
func parse(tmpl jsonline.Template, def *RowDefinition) (jsonline.Template, error) {
	for _, column := range def.Columns {
		switch column.Type {
		case "string":
			tmpl = tmpl.WithString(column.Name)
		case "numeric":
			tmpl = tmpl.WithNumeric(column.Name)
		case "boolean":
			tmpl = tmpl.WithBoolean(column.Name)
		case "binary":
			tmpl = tmpl.WithBinary(column.Name)
		case "datetime":
			tmpl = tmpl.WithDateTime(column.Name)
		case "time":
			tmpl = tmpl.WithTime(column.Name)
		case "timestamp":
			tmpl = tmpl.WithTimestamp(column.Name)
		case "auto":
			tmpl = tmpl.WithAuto(column.Name)
		case "hidden":
			tmpl = tmpl.WithHidden(column.Name)
		}
	}

	return tmpl, nil
}
