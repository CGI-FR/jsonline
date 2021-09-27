package jsonline

import "fmt"

type (
	Processor func(Row, error) error
)

//nolint:gochecknoglobals
var (
	DefaultProcessor Processor = func(r Row, e error) error { return e }
)

type Streamer interface {
	WithProcessor(Processor) Streamer
	Stream() error
}

type streamer struct {
	exporter  Exporter
	importer  Importer
	processor Processor
}

func NewStreamer(exporter Exporter, importer Importer) Streamer {
	return &streamer{
		exporter:  exporter,
		importer:  importer,
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
