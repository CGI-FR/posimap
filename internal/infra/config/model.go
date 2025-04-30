package config

import (
	"errors"
	"fmt"

	"github.com/cgi-fr/posimap/pkg/data"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Schema Schema `yaml:"schema"`
	Trim   bool   `yaml:"trim,omitempty"`
}

type Schema []Field

type Either[T1 any, T2 any] struct {
	T1 T1
	T2 T2
}

func (e *Either[T1, T2]) UnmarshalYAML(value *yaml.Node) error {
	out, err := yaml.Marshal(value)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := yaml.Unmarshal(out, &e.T1); err == nil {
		return nil
	}

	if err := yaml.Unmarshal(out, &e.T2); err == nil {
		return nil
	}

	return fmt.Errorf("%w", err)
}

type Field struct {
	Name     string                 `yaml:"name"`
	Length   int                    `yaml:"length"`
	Occurs   int                    `yaml:"occurs,omitempty"`
	Redefine string                 `yaml:"redefine,omitempty"`
	Trim     bool                   `yaml:"trim,omitempty"`
	When     string                 `yaml:"when,omitempty"`
	Schema   Either[string, Schema] `yaml:"schema,omitempty"` // either filename for external schema or inlined schema
	schema   Schema
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

func (f Field) Build(trim bool) data.FieldSchema {
	return data.FieldSchema{
		Name:     f.Name,
		Length:   f.Length,
		Occurs:   f.Occurs,
		Redefine: f.Redefine,
		Trim:     f.Trim || trim,
		When:     data.When(f.When),
		Schema:   f.schema.Compile(trim),
	}
}

func (s Schema) Validate() error {
	for idx, field := range s {
		if err := field.Validate(); err != nil {
			return fmt.Errorf("%v.%w", idx, err)
		}

		if field.schema != nil {
			if err := field.schema.Validate(); err != nil {
				return fmt.Errorf("%v.%w", idx, err)
			}
		}
	}

	return nil
}

func (s Schema) Compile(trim bool) data.RecordSchema {
	if len(s) == 0 {
		return nil
	}

	result := make(data.RecordSchema, len(s))
	for i, field := range s {
		result[i] = field.Build(trim)
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
	return c.Schema.Compile(c.Trim)
}
