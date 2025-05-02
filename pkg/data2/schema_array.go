package data2

import (
	"fmt"

	"github.com/cgi-fr/posimap/pkg/flat"
)

type SchemaArray struct {
	schema Schema
	occurs int
}

func NewSchemaArray(schema Schema, occurs int) *SchemaArray {
	return &SchemaArray{
		schema: schema,
		occurs: occurs,
	}
}

func (a *SchemaArray) RuneCount() int {
	return a.schema.RuneCount() * a.occurs
}

func (a *SchemaArray) ReadBuffer(source flat.Source, buffer *Buffer) error {
	if source.IsFixedWidth() {
		runes, err := source.ReadRunes(a.RuneCount())
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		*buffer = append(*buffer, runes...)

		return nil
	}

	for range a.occurs {
		if err := a.schema.ReadBuffer(source, buffer); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (a *SchemaArray) CreateRecord(buffer Buffer, start int) (Record, error) { //nolint:ireturn
	records := make([]Record, a.occurs)

	pos := start

	for idx := range a.occurs {
		record, err := a.schema.CreateRecord(buffer, pos)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		records[idx] = record

		pos += a.schema.RuneCount()
	}

	return RecordArray{
		schema:  a,
		records: records,
	}, nil
}
