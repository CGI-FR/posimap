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

func (f Field) Compile(record schema.Record) {
	if f.IsRecord() {
		record.WithRecord(f.Name, f.Schema.T2.Compile(), f.CompileOptions()...)
	} else {
		record.WithField(f.Name, codec.NewString(f.CompileCharset(), f.Length, f.Trim), f.CompileOptions()...)
	}
}

func (s Schema) Compile() schema.Record {
	record := make(schema.Record, 0, len(s))

	for _, field := range s {
		field.Compile(record)
	}

	return record
}

func (c Config) Compile() schema.Record {
	schema := make(schema.Record, 0, len(c.Schema))

	for _, field := range c.Schema {
		field.Compile(schema)
	}

	return schema
}
