package data2

import (
	"fmt"
	"iter"

	"github.com/cgi-fr/posimap/pkg/flat"
)

type SchemaObject struct {
	keys      []string
	values    map[string]Schema
	redefines map[string]string
	export    Predicate
}

func NewSchemaObject(export Predicate) *SchemaObject {
	return &SchemaObject{
		keys:      make([]string, 0),
		values:    make(map[string]Schema),
		redefines: make(map[string]string),
		export:    export,
	}
}

func (o *SchemaObject) Add(key string, value Schema, redefine string) {
	o.keys = append(o.keys, key)
	o.values[key] = value

	if redefine != "" {
		o.redefines[key] = redefine
	}
}

func (o *SchemaObject) MainKeys() iter.Seq[string] {
	return func(yield func(key string) bool) {
		for _, key := range o.keys {
			if _, ok := o.redefines[key]; !ok {
				if !yield(key) {
					return
				}
			}
		}
	}
}

func (o *SchemaObject) RuneCount() int {
	length := 0
	for key := range o.MainKeys() {
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
		for key := range o.MainKeys() {
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
	indexes := make(map[string]int)

	for _, key := range o.keys {
		schema := o.values[key]

		if redefine, ok := o.redefines[key]; ok {
			rpos, ok := indexes[redefine]
			if !ok {
				return nil, fmt.Errorf("redefine %s not found", redefine) //nolint:err113
			}

			pos = rpos
		}

		indexes[key] = pos

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
