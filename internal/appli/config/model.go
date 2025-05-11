package config

import (
	"errors"
	"fmt"

	"github.com/cgi-fr/posimap/pkg/posimap/core/codec"
	"github.com/cgi-fr/posimap/pkg/posimap/core/predicate"
	"github.com/cgi-fr/posimap/pkg/posimap/core/schema2"
	"golang.org/x/text/encoding/charmap"
)

var ErrUnsupportedCharset = errors.New("unsupported charset")

type Config struct {
	Length int    `yaml:"length,omitempty"` // Length is the length of the record, optional
	Schema Schema `yaml:"schema"`
}

type Schema []Field

type Field struct {
	Name     string `yaml:"name"`
	Occurs   int    `yaml:"occurs,omitempty"`
	Redefine string `yaml:"redefine,omitempty"`
	When     string `yaml:"when,omitempty"`

	Length  int     `yaml:"length"`
	Trim    *bool   `yaml:"trim,omitempty"`    // Trim can be nil if not set in the configuration file
	Charset *string `yaml:"charset,omitempty"` // Charset can be nil if not set in the configuration file

	Schema Either[string, Schema] `yaml:"schema"` // Schema is either a filename (external schema) or an embedded schema
}

func (f Field) IsRecord() bool {
	return f.Schema.T2 != nil
}

func (f Field) CompileOptions() []schema2.Option {
	options := make([]schema2.Option, 0)

	if f.Occurs > 0 {
		options = append(options, schema2.Occurs(f.Occurs))
	}

	if f.Redefine != "" {
		options = append(options, schema2.Redefines(f.Redefine))
	}

	if f.When != "" {
		options = append(options, schema2.Condition(predicate.When(f.When)))
	}

	return options
}

func (f Field) CompileCharset() (*charmap.Charmap, error) {
	for _, encoding := range charmap.All {
		if charmap, ok := encoding.(*charmap.Charmap); ok && charmap.String() == *f.Charset {
			return charmap, nil
		}
	}

	return nil, fmt.Errorf("%w: %s", ErrUnsupportedCharset, *f.Charset)
}

func (f Field) Compile(record *schema2.Record, defaults ...Default) (*schema2.Record, error) {
	if f.IsRecord() {
		schema, err := f.Schema.T2.Compile(defaults...)
		if err != nil {
			return nil, err
		}

		record = record.WithRecord(f.Name, schema, f.CompileOptions()...)
	} else {
		charset, err := f.CompileCharset()
		if err != nil {
			return nil, err
		}

		record = record.WithField(f.Name, codec.NewString(charset, f.Length, *f.Trim), f.CompileOptions()...)
	}

	return record, nil
}

func (s Schema) Compile(defaults ...Default) (*schema2.Record, error) {
	var err error

	record := schema2.NewRecord("ROOT")

	for _, field := range s {
		for _, defaultFunc := range defaults {
			field = defaultFunc(field)
		}

		record, err = field.Compile(record, defaults...)
		if err != nil {
			return nil, fmt.Errorf("failed to compile field %s: %w", field.Name, err)
		}
	}

	return record, nil
}

func (c Config) Compile(defaults ...Default) (*schema2.Record, error) {
	return c.Schema.Compile(defaults...)
}

type Default func(field Field) Field

func Trim(enable bool) Default {
	return func(field Field) Field {
		if field.Trim == nil {
			field.Trim = &enable
		}

		return field
	}
}

func Charset(name string) Default {
	return func(field Field) Field {
		if field.Charset == nil {
			field.Charset = &name
		}

		return field
	}
}
