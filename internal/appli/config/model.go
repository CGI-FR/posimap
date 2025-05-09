package config

import (
	"github.com/cgi-fr/posimap/pkg/posimap/core/codec"
	"github.com/cgi-fr/posimap/pkg/posimap/core/predicate"
	"github.com/cgi-fr/posimap/pkg/posimap/core/schema"
	"golang.org/x/text/encoding/charmap"
)

type Config struct {
	Schema Schema `yaml:"schema"`
}

type Schema []Field

type Field struct {
	Name     string `yaml:"name"`
	Occurs   int    `yaml:"occurs,omitempty"`
	Redefine string `yaml:"redefine,omitempty"`
	When     string `yaml:"when,omitempty"`

	Length  int    `yaml:"length"`
	Trim    bool   `yaml:"trim,omitempty"`
	Charset string `yaml:"charset,omitempty"`

	Schema Either[string, Schema] `yaml:"schema"` // either filename for external schema or embedded schema
}

func (f Field) IsRecord() bool {
	return f.Schema.T2 != nil
}

func (f Field) CompileOptions() []schema.Option {
	options := make([]schema.Option, 0)

	if f.Occurs > 0 {
		options = append(options, schema.Occurs(f.Occurs))
	}

	if f.Redefine != "" {
		options = append(options, schema.Redefines(f.Redefine))
	}

	if f.When != "" {
		options = append(options, schema.Condition(predicate.When(f.When)))
	}

	return options
}

func (f Field) CompileCharset() *charmap.Charmap {
	for _, encoding := range charmap.All {
		if charmap, ok := encoding.(*charmap.Charmap); ok && charmap.String() == f.Charset {
			return charmap
		}
	}

	return charmap.ISO8859_1
}

func (f Field) Compile(record schema.Record, defaults ...Default) schema.Record {
	if f.IsRecord() {
		record = record.WithRecord(f.Name, f.Schema.T2.Compile(defaults...), f.CompileOptions()...)
	} else {
		record = record.WithField(f.Name, codec.NewString(f.CompileCharset(), f.Length, f.Trim), f.CompileOptions()...)
	}

	return record
}

func (s Schema) Compile(defaults ...Default) schema.Record {
	record := make(schema.Record, 0, len(s))

	for _, field := range s {
		for _, defaultFunc := range defaults {
			field = defaultFunc(field)
		}

		record = field.Compile(record)
	}

	return record
}

func (c Config) Compile(defaults ...Default) schema.Record {
	schema := make(schema.Record, 0, len(c.Schema))

	for _, field := range c.Schema {
		for _, defaultFunc := range defaults {
			field = defaultFunc(field)
		}

		schema = field.Compile(schema, defaults...)
	}

	return schema
}

type Default func(field Field) Field

func Trim(enable bool) Default {
	return func(field Field) Field {
		field.Trim = enable

		return field
	}
}
