package config

import (
	"errors"
	"fmt"

	"github.com/cgi-fr/posimap/pkg/data"
)

type Config struct {
	Schema Schema `yaml:"schema"`
}

type Schema []Field

type Field struct {
	Name     string `yaml:"name"`
	Length   int    `yaml:"length"`
	Occurs   int    `yaml:"occurs,omitempty"`
	Redefine string `yaml:"redefine,omitempty"`
	When     string `yaml:"when,omitempty"`
	Schema   Schema `yaml:"schema,omitempty"`
}

var (
	ErrFieldNameEmpty     = errors.New("field name cannot be empty")
	ErrFieldLengthZero    = errors.New("field length cannot be zero")
	ErrFieldStartNegative = errors.New("field start position cannot be negative")
)

func (f Field) Validate() error {
	if f.Name == "" {
		return ErrFieldNameEmpty
	}

	return nil
}

func (f Field) Build() data.FieldSchema {
	return data.FieldSchema{
		Name:     f.Name,
		Length:   f.Length,
		Occurs:   f.Occurs,
		Redefine: f.Redefine,
		When:     data.When(f.When),
		Schema:   f.Schema.Compile(),
	}
}

func (t Schema) Validate() error {
	for idx, field := range t {
		if err := field.Validate(); err != nil {
			return fmt.Errorf("%v.%w", idx, err)
		}

		if field.Schema != nil {
			if err := field.Schema.Validate(); err != nil {
				return fmt.Errorf("%v.%w", idx, err)
			}
		}
	}

	return nil
}

func (t Schema) Compile() data.RecordSchema {
	if len(t) == 0 {
		return nil
	}

	result := make(data.RecordSchema, len(t))
	for i, field := range t {
		result[i] = field.Build()
	}

	return result
}

func (c Config) Validate() error {
	if err := c.Schema.Validate(); err != nil {
		return fmt.Errorf("schema validation failed: %w", err)
	}

	return nil
}

func (c Config) Compile() data.RecordSchema {
	return c.Schema.Compile()
}
