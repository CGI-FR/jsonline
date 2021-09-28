// Copyright (C) 2021 CGI France
//
// This file is part of the jsonline library.
//
// The jsonline library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The jsonline library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the jsonline library.  If not, see <http://www.gnu.org/licenses/>.
//
// Linking this library statically or dynamically with other modules is
// making a combined work based on this library.  Thus, the terms and
// conditions of the GNU General Public License cover the whole
// combination.
//
// As a special exception, the copyright holders of this library give you
// permission to link this library with independent modules to produce an
// executable, regardless of the license terms of these independent
// modules, and to copy and distribute the resulting executable under
// terms of your choice, provided that you also meet, for each linked
// independent module, the terms and conditions of the license of that
// module.  An independent module is a module which is not derived from
// or based on this library.  If you modify this library, you may extend
// this exception to your version of the library, but you are not
// obligated to do so.  If you do not wish to do so, delete this
// exception statement from your version.

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
	Import() bool
	GetRow() (Row, error)
}

type importer struct {
	r io.Reader
	s *bufio.Scanner
	t Template
}

func NewImporter(r io.Reader) Importer {
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

func (i *importer) Import() bool {
	return i.s.Scan()
}

func (i *importer) GetRow() (Row, error) {
	if i.s.Err() != nil {
		return nil, fmt.Errorf("%w", i.s.Err())
	}

	b := i.s.Bytes()

	row := i.t.CreateRowEmpty()
	if err := row.UnmarshalJSON(b); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return row, nil
}
