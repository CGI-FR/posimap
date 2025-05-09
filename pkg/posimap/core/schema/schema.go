package schema

import (
	"errors"
	"fmt"

	"github.com/cgi-fr/posimap/pkg/posimap/api"
	"github.com/cgi-fr/posimap/pkg/posimap/core/codec"
	"github.com/cgi-fr/posimap/pkg/posimap/core/record"
	"golang.org/x/text/encoding/charmap"
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

func (r Record) WithField(name string, codec api.Codec[any], options ...Option) Record {
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

// updateOffsetForRedefines checks if the field redefines another field and updates the offset accordingly.
// it returns the length of the redefined field or 0 if no field is redefined.
func (f Field) updateOffsetForRedefines(offset *int, redefines map[string]int) (int, error) {
	if f.redefines == "" {
		return 0, nil
	}

	// Check if the field redefines another field and adjust the offset accordingly.
	if pos, ok := redefines[f.redefines]; ok {
		length := *offset - pos
		*offset = pos

		return length, nil
	}

	return 0, fmt.Errorf("%w: %s", ErrInvalidRedefines, f.redefines)
}

func (f Field) build(rec *record.Object, offset *int, redefines map[string]int) error {
	length, err := f.updateOffsetForRedefines(offset, redefines)
	if err != nil {
		return err
	}

	target := *offset + length

	redefines[f.name] = *offset

	var sub *record.Object

	switch {
	case f.occurs == 0 && f.definition.IsCodec():
		f.buildCodec(rec, offset)
	case f.occurs > 0 && f.definition.IsCodec():
		f.buildCodecArray(rec, offset)
	case f.occurs == 0 && f.definition.IsSchema():
		sub, err = f.buildSchema(rec, offset)
	case f.occurs > 0 && f.definition.IsSchema():
		err = f.buildSchemaArray(rec, offset)
	}

	if *offset < target && length > 0 && sub != nil {
		println("missing filler of", target-*offset, "bytes under", f.name)
		sub.Add("FILLER", record.NewValue(*offset, codec.NewString(charmap.ISO8859_1, target-*offset, true)), nil)
		*offset = target
	}

	return err
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

func (f Field) buildSchema(rec *record.Object, offset *int) (*record.Object, error) {
	sub := record.NewObject()
	if err := f.definition.Schema().build(sub, offset); err != nil {
		return nil, err
	}

	rec.Add(f.name, sub, f.when)

	return sub, nil
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
