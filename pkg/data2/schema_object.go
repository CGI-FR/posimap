package data2

import (
	"fmt"

	"github.com/cgi-fr/posimap/pkg/flat"
)

type SchemaObject struct {
	keys   []string
	values map[string]Schema
	export Predicate
}

func NewSchemaObject(export Predicate) *SchemaObject {
	return &SchemaObject{
		keys:   make([]string, 0),
		values: make(map[string]Schema),
		export: export,
	}
}

func (o *SchemaObject) Add(key string, value Schema) {
	o.keys = append(o.keys, key)
	o.values[key] = value
}

func (o *SchemaObject) RuneCount() int {
	length := 0
	for _, key := range o.keys {
		length += o.values[key].RuneCount()
	}

	return length
}

func (o *SchemaObject) ReadBuffer(source flat.Source, buffer *Buffer) error {
	if source.IsFixedWidth() {
		runes, err := source.ReadRunes(o.RuneCount())
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		*buffer = append(*buffer, runes...)
	} else {
		for _, key := range o.keys {
			if err := o.values[key].ReadBuffer(source, buffer); err != nil {
				return fmt.Errorf("%w", err)
			}
		}
	}

	return nil
}

func (o *SchemaObject) CreateRecord(buffer Buffer, start int) (Record, error) { //nolint:ireturn
	records := make(map[string]Record, len(o.keys))

	pos := start

	for _, key := range o.keys {
		schema := o.values[key]

		r, err := schema.CreateRecord(buffer, pos)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		records[key] = r
		pos += schema.RuneCount()
	}

	return RecordObject{
		schema:  o,
		records: records,
		export:  o.export,
	}, nil
}
