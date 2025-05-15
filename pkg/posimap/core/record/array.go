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
	"fmt"

	"github.com/cgi-fr/posimap/pkg/posimap/api"
	"github.com/cgi-fr/posimap/pkg/posimap/driven/document"
)

type Array struct {
	records []Record
}

func NewArray() *Array {
	return &Array{
		records: make([]Record, 0),
	}
}

func (a *Array) Add(record Record) {
	a.records = append(a.records, record)
}

func (a *Array) Unmarshal(buffer api.Buffer) error {
	for _, record := range a.records {
		if err := record.Unmarshal(buffer); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (a *Array) Marshal(buffer api.Buffer) error {
	for _, record := range a.records {
		if err := record.Marshal(buffer); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (a *Array) Export(writer document.Writer) error {
	return a.export(writer, a)
}

func (a *Array) export(writer document.Writer, _ Record) error {
	if err := writer.WriteToken(document.TokenArrStart); err != nil {
		return fmt.Errorf("%w", err)
	}

	for _, record := range a.records {
		if err := record.Export(writer); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	if err := writer.WriteToken(document.TokenArrEnd); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (a *Array) Import(value any) error {
	switch typed := value.(type) {
	case []any:
		for idx, record := range a.records {
			if idx >= len(typed) {
				return fmt.Errorf("%w: expected %d elements, got %d", ErrInvalidType, len(a.records), len(typed))
			}

			if err := record.Import(typed[idx]); err != nil {
				return fmt.Errorf("%w", err)
			}
		}
	default:
		return fmt.Errorf("%w: expected array, got %T", ErrInvalidType, typed)
	}

	return nil
}

func (a *Array) AsPrimitive() any {
	primitive := make([]any, len(a.records))

	for idx, record := range a.records {
		primitive[idx] = record.AsPrimitive()
	}

	return primitive
}

func (a *Array) Reset() {
	for _, record := range a.records {
		record.Reset()
	}
}
