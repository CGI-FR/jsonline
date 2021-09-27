package jsonline

import (
	"fmt"
	"io"
)

const lineSeparator byte = 10

type Exporter interface {
	Export(interface{}) error
}

type exporter struct {
	w io.Writer
	t Template
}

func NewExporter(w io.Writer, t Template) Exporter {
	return &exporter{
		w: w,
		t: t,
	}
}

func (e *exporter) Export(input interface{}) error {
	row := e.t.Create(input)

	b, err := row.MarshalJSON()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	b = append(b, lineSeparator)

	if _, err := e.w.Write(b); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
