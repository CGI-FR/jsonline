package jsonline

import (
	"bufio"
	"fmt"
	"io"
)

const (
	initialBufferSize = 64 * 1024        // 64 Kb
	maximumBufferSize = 10 * 1024 * 1024 // 10 Mb
)

type Importer interface {
	WithTemplate(Template) Importer
	Import() (Row, error)
}

type importer struct {
	r io.Reader
	s *bufio.Scanner
	t Template
}

func NewImporter(r io.Reader, t Template) Importer {
	buf := make([]byte, 0, initialBufferSize)
	s := bufio.NewScanner(r)
	s.Buffer(buf, maximumBufferSize)

	return &importer{
		r: r,
		s: s,
		t: NewTemplate(),
	}
}

func (i *importer) WithTemplate(t Template) Importer {
	i.t = t

	return i
}

func (i *importer) Import() (Row, error) {
	if !i.s.Scan() {
		if i.s.Err() != nil {
			return nil, fmt.Errorf("%w", i.s.Err())
		}

		return nil, nil
	}

	b := i.s.Bytes()

	row := i.t.CreateEmpty()
	if err := row.UnmarshalJSON(b); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return row, nil
}
