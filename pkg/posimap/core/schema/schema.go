package schema

import (
	"errors"
	"fmt"

	"github.com/cgi-fr/posimap/pkg/posimap/api"
	"github.com/cgi-fr/posimap/pkg/posimap/core/record"
)

var ErrInvalidRedefines = errors.New("invalid redefines")

type Option func(Field) Field

func Occurs(occurs int) Option {
	return func(field Field) Field {
		field.occurs = occurs

		return field
	}
}

func Condition(when api.Predicate) Option {
	return func(field Field) Field {
		field.when = when

		return field
	}
}

func Redefines(redefines string) Option {
	return func(field Field) Field {
		field.redefines = redefines

		return field
	}
}

type Field struct {
	name       string
	redefines  string
	occurs     int
	when       api.Predicate
	definition Definition
}

type Record []Field

func NewSchema() Record {
	return make(Record, 0)
}

func (r Record) WithField(name string, codec api.Codec, options ...Option) Record {
	field := Field{
		name:       name,
		redefines:  "",
		occurs:     0,
		when:       nil,
		definition: NewCodec(codec),
	}

	for _, option := range options {
		field = option(field)
	}

	return append(r, field)
}

func (r Record) WithRecord(name string, schema Record, options ...Option) Record {
	field := Field{
		name:       name,
		redefines:  "",
		occurs:     0,
		when:       nil,
		definition: NewDefinition(schema),
	}

	for _, option := range options {
		field = option(field)
	}

	return append(r, field)
}

func (r Record) Build() (*record.Object, error) {
	rec := record.NewObject()
	offset := 0

	if err := r.build(rec, &offset); err != nil {
		return nil, err
	}

	return rec, nil
}

func (r Record) build(rec *record.Object, offset *int) error {
	redefines := make(map[string]int)

	for _, field := range r {
		if err := field.build(rec, offset, redefines); err != nil {
			return err
		}
	}

	return nil
}

func (f Field) updateOffsetForRedefines(offset *int, redefines map[string]int) error {
	if f.redefines == "" {
		return nil
	}

	// Check if the field redefines another field and adjust the offset accordingly.
	if pos, ok := redefines[f.redefines]; ok {
		*offset = pos

		return nil
	}

	return fmt.Errorf("%w: %s", ErrInvalidRedefines, f.redefines)
}

func (f Field) build(rec *record.Object, offset *int, redefines map[string]int) error {
	if err := f.updateOffsetForRedefines(offset, redefines); err != nil {
		return err
	}

	redefines[f.name] = *offset

	switch {
	case f.occurs == 0 && f.definition.IsCodec():
		f.buildCodec(rec, offset)
	case f.occurs > 0 && f.definition.IsCodec():
		f.buildCodecArray(rec, offset)
	case f.occurs == 0 && f.definition.IsSchema():
		return f.buildSchema(rec, offset)
	case f.occurs > 0 && f.definition.IsSchema():
		return f.buildSchemaArray(rec, offset)
	}

	return nil
}

func (f Field) buildCodec(rec *record.Object, offset *int) {
	rec.Add(f.name, record.NewValue(*offset, f.definition.Codec()), f.when)

	*offset += f.definition.Codec().Size()
}

func (f Field) buildCodecArray(rec *record.Object, offset *int) {
	array := record.NewArray()
	for range f.occurs {
		array.Add(record.NewValue(*offset, f.definition.Codec()))
		*offset += f.definition.Codec().Size()
	}

	rec.Add(f.name, array, f.when)
}

func (f Field) buildSchema(rec *record.Object, offset *int) error {
	sub := record.NewObject()
	if err := f.definition.Schema().build(sub, offset); err != nil {
		return err
	}

	rec.Add(f.name, sub, f.when)

	return nil
}

func (f Field) buildSchemaArray(rec *record.Object, offset *int) error {
	array := record.NewArray()

	for range f.occurs {
		sub := record.NewObject()
		if err := f.definition.Schema().build(sub, offset); err != nil {
			return err
		}

		array.Add(sub)
	}

	rec.Add(f.name, array, f.when)

	return nil
}
