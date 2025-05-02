package data2

import (
	"fmt"

	"github.com/cgi-fr/posimap/pkg/deep"
)

type RecordArray struct {
	schema  *SchemaArray
	records []Record
}

func (ra RecordArray) Materialize() any {
	result := make([]any, len(ra.records))
	for i, record := range ra.records {
		result[i] = record.Materialize()
	}

	return result
}

func (ra RecordArray) Export(sink deep.Sink) error {
	if err := sink.OpenArray(); err != nil {
		return fmt.Errorf("%w", err)
	}

	for idx, record := range ra.records {
		if idx != 0 {
			if err := sink.Next(); err != nil {
				return fmt.Errorf("%w", err)
			}
		}

		if err := record.Export(sink); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	if err := sink.CloseArray(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
