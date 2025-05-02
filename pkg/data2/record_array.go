package data2

import (
	"fmt"

	"github.com/cgi-fr/posimap/pkg/deep"
)

type RecordArray struct {
	schema  *SchemaArray
	records []Record
	export  Predicate
}

func (ra RecordArray) Materialize() any {
	result := make([]any, len(ra.records))
	for i, record := range ra.records {
		result[i] = record.Materialize()
	}

	return result
}

func (ra RecordArray) Export(root Record, sink deep.Sink) error {
	if err := sink.OpenArray(); err != nil {
		return fmt.Errorf("%w", err)
	}

	for idx, record := range ra.records {
		if !ra.records[idx].VisibleFrom(root) {
			continue
		}

		if idx != 0 {
			if err := sink.Next(); err != nil {
				return fmt.Errorf("%w", err)
			}
		}

		if err := record.Export(root, sink); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	if err := sink.CloseArray(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (ra RecordArray) VisibleFrom(root Record) bool {
	return ra.export == nil || ra.export(root)
}
