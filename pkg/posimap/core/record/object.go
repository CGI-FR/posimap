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

package record

import (
	"errors"
	"fmt"

	"github.com/cgi-fr/posimap/pkg/posimap/api"
	"github.com/cgi-fr/posimap/pkg/posimap/driven/document"
)

var (
	ErrUnexpectedTokenType = errors.New("unexpected token type")
	ErrUnexpectedKey       = errors.New("unexpected key")
	ErrInvalidType         = errors.New("invalid type")
)

type Object struct {
	keys     []string
	records  map[string]Record
	exports  map[string]api.Predicate
	feedback bool
}

func NewObject() *Object {
	return &Object{
		keys:     make([]string, 0),
		records:  make(map[string]Record),
		exports:  make(map[string]api.Predicate),
		feedback: false,
	}
}

func (o *Object) SetFeedback() {
	o.feedback = true
}

func (o *Object) Add(key string, record Record, export api.Predicate) {
	o.keys = append(o.keys, key)
	o.records[key] = record
	o.exports[key] = export
}

func (o *Object) Unmarshal(buffer api.Buffer) error {
	for _, key := range o.keys {
		record := o.records[key]

		if err := record.Unmarshal(buffer); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (o *Object) Marshal(buffer api.Buffer) error {
	for _, key := range o.keys {
		record := o.records[key]

		if err := record.Marshal(buffer); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (o *Object) Export(writer document.Writer) error {
	return o.export(writer, o)
}

//nolint:cyclop
func (o *Object) export(writer document.Writer, feedback Record) error {
	if err := writer.WriteToken(document.TokenObjStart); err != nil {
		return fmt.Errorf("%w", err)
	}

	if o.feedback {
		feedback = o
	}

	for _, key := range o.keys {
		record := o.records[key]

		if export, ok := o.exports[key]; ok && export != nil && feedback != nil {
			// Skip the record if the export predicate returns false
			if ok, err := export(feedback); err != nil {
				return fmt.Errorf("%w", err)
			} else if !ok {
				continue
			}
		}

		if err := writer.WriteString(key); err != nil {
			return fmt.Errorf("%w", err)
		}

		if err := record.export(writer, feedback); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	if err := writer.WriteToken(document.TokenObjEnd); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (o *Object) Import(value any) error {
	switch typed := value.(type) {
	case map[string]any:
		for key, record := range o.records {
			if err := record.Import(typed[key]); err != nil {
				return fmt.Errorf("%w", err)
			}
		}
	case nil: // skip nil values
	default:
		return fmt.Errorf("%w: expected object, got %T", ErrInvalidType, typed)
	}

	return nil
}

func (o *Object) AsPrimitive() any {
	primitive := make(map[string]any, len(o.keys))

	for _, key := range o.keys {
		record := o.records[key]

		primitive[key] = record.AsPrimitive()
	}

	return primitive
}

func (o *Object) Reset() {
	for _, record := range o.records {
		record.Reset()
	}
}
