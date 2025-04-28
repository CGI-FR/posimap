package config

import (
	"errors"
	"fmt"

	"github.com/cgi-fr/posimap/pkg/data"
)

type Template []Field

type Field struct {
	Name     string   `yaml:"name"`
	Length   int      `yaml:"length"`
	Occurs   int      `yaml:"occurs,omitempty"`
	Redefine string   `yaml:"redefine,omitempty"`
	When     string   `yaml:"when,omitempty"`
	Template Template `yaml:"schema,omitempty"`
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

func (f Field) Build() data.FieldTemplate {
	return data.FieldTemplate{
		Name:     f.Name,
		Length:   f.Length,
		Occurs:   f.Occurs,
		Redefine: f.Redefine,
		When:     data.When(f.When),
		Template: f.Template.Build(),
	}
}

func (t Template) Validate() error {
	for idx, field := range t {
		if err := field.Validate(); err != nil {
			return fmt.Errorf("%v.%w", idx, err)
		}

		if field.Template != nil {
			if err := field.Template.Validate(); err != nil {
				return fmt.Errorf("%v.%w", idx, err)
			}
		}
	}

	return nil
}

func (t Template) Build() data.RecordTemplate {
	if len(t) == 0 {
		return nil
	}

	result := make(data.RecordTemplate, len(t))
	for i, field := range t {
		result[i] = field.Build()
	}

	return result
}
