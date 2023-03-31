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

//nolint:ireturn
package jsonline

import (
	"fmt"
)

type (
	Processor func(Row, error) error
)

func DefaultProcessor(r Row, e error) error { return e }

func NoFailureProcessor(r Row, e error) error {
	return nil
}

type Streamer interface {
	WithProcessor(Processor) Streamer
	Stream() error
}

type streamer struct {
	importer  Importer
	exporter  Exporter
	processor Processor
}

func NewStreamer(importer Importer, exporter Exporter) Streamer {
	return &streamer{
		importer:  importer,
		exporter:  exporter,
		processor: DefaultProcessor,
	}
}

func (s *streamer) WithProcessor(p Processor) Streamer {
	if p == nil {
		s.processor = DefaultProcessor
	} else {
		s.processor = p
	}

	return s
}

func (s *streamer) Stream() error {
	for i := 1; s.importer.Import(); i++ {
		row, err := s.importer.GetRow()
		if err != nil {
			if err := s.processor(row, fmt.Errorf("%w", err)); err != nil {
				return err
			}

			continue
		}

		if err := s.processor(row, nil); err != nil {
			return err
		}

		if err := s.exporter.Export(row); err != nil {
			if err := s.processor(row, fmt.Errorf("%w", err)); err != nil {
				return err
			}
		}
	}

	return nil
}
