package data2

import (
	"fmt"

	"github.com/cgi-fr/posimap/pkg/flat"
)

type SchemaValue struct {
	length int
	trim   bool
	export Predicate
}

func NewSchemaValue(length int, trim bool, export Predicate) SchemaValue {
	return SchemaValue{
		length: length,
		trim:   trim,
		export: export,
	}
}

func (v SchemaValue) RuneCount() int {
	return v.length
}

func (v SchemaValue) ReadBuffer(source flat.Source, buffer *Buffer) error {
	runes, err := source.ReadRunes(v.length)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	*buffer = append(*buffer, runes...)

	return nil
}

func (v SchemaValue) CreateRecord(buffer Buffer, start int) (Record, error) { //nolint:ireturn
	return RecordValue{
		schema: v,
		buffer: buffer,
		start:  start,
		trim:   v.trim,
		export: v.export,
	}, nil
}
