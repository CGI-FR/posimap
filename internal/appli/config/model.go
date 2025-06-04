// Copyright (C) 2025 CGI France
//
// This file is part of posimap.
//
// posimap is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// posimap is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with posimap.  If not, see <http://www.gnu.org/licenses/>.

package config

import (
	"errors"
	"fmt"

	"github.com/cgi-fr/posimap/internal/appli/charsets"
	"github.com/cgi-fr/posimap/pkg/posimap/core/codec"
	"github.com/cgi-fr/posimap/pkg/posimap/core/predicate"
	"github.com/cgi-fr/posimap/pkg/posimap/core/schema"
	"golang.org/x/text/encoding/charmap"
)

var (
	ErrUnsupportedCharset      = errors.New("unsupported charset")
	ErrEitherLengthOrSeparator = errors.New("either length or separator must be set")
	ErrInvalidLength           = errors.New("length must be greater than 0")
)

type Config struct {
	Length    int    `yaml:"length,omitempty"`    // Length is the fixed length of the record, optional
	Separator string `yaml:"separator,omitempty"` // Separator is the separator between records, optional
	Feedback  *bool  `yaml:"feedback,omitempty"`
	Schema    Schema `yaml:"schema"`
}

type Schema []Field

type Field struct {
	Name     string `yaml:"name"`
	Occurs   int    `yaml:"occurs,omitempty"`
	Redefine string `yaml:"redefine,omitempty"`
	When     string `yaml:"when,omitempty"`
	Feedback bool   `yaml:"feedback,omitempty"`

	Length  int     `yaml:"length"`
	Trim    *bool   `yaml:"trim,omitempty"`    // Trim can be nil if not set in the configuration file
	Charset *string `yaml:"charset,omitempty"` // Charset can be nil if not set in the configuration file
	Picture Picture `yaml:"picture,omitempty"` // Picture is an optional string representation of the format
	Codec   string  `yaml:"codec,omitempty"`   // Codec is the codec to use for this field, default to String

	Schema Either[string, Schema] `yaml:"schema"` // Schema is either a filename (external schema) or an embedded schema
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

	if f.Feedback {
		options = append(options, schema.Feedback())
	}

	return options
}

func (f Field) CompileCharset() (*charmap.Charmap, error) {
	charset, err := charsets.Get(*f.Charset)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedCharset, *f.Charset)
	}

	return charset, nil
}

func (f Field) Compile(record *schema.Record, defaults ...Default) (*schema.Record, error) {
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

		switch f.Codec {
		case "COMP-3":
			format, err := f.Picture.Compile()
			if err != nil {
				return nil, fmt.Errorf("failed to compile picture for field %s: %w", f.Name, err)
			}

			record = record.WithField(f.Name, codec.NewComp3(format.Length, format.Decimal, format.Signed),
				f.CompileOptions()...)
		default:
			record = record.WithField(f.Name, codec.NewString(charset, f.Length, *f.Trim), f.CompileOptions()...)
		}
	}

	return record, nil
}

func (s Schema) Compile(defaults ...Default) (*schema.Record, error) {
	var err error

	record := schema.NewRecord("ROOT")

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

func (c Config) Compile(defaults ...Default) (*schema.Record, error) {
	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate configuration: %w", err)
	}

	return c.Schema.Compile(defaults...)
}

func (c Config) Validate() error {
	// Either length or separator must be set
	if c.Length == 0 && c.Separator == "" || c.Length != 0 && c.Separator != "" {
		return fmt.Errorf("%w", ErrEitherLengthOrSeparator)
	} else if c.Length < 0 {
		return fmt.Errorf("%w: got %d", ErrInvalidLength, c.Length)
	}

	return nil
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
