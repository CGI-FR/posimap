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

var ErrUnexpectedValueType = errors.New("unexpected value type")

type Value struct {
	offset  int
	codec   api.Codec[any]
	content any
}

func NewValue(offset int, codec api.Codec[any]) *Value {
	return &Value{
		offset:  offset,
		codec:   codec,
		content: nil,
	}
}

func (v *Value) Unmarshal(buffer api.Buffer) error {
	var err error

	v.content, err = v.codec.Decode(buffer, v.offset)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (v *Value) Marshal(buffer api.Buffer) error {
	if v.content == nil {
		return nil // document did not have the key set
	}

	err := v.codec.Encode(buffer, v.offset, v.content)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (v *Value) export(writer document.Writer, _ Record) error {
	return v.Export(writer)
}

func (v *Value) Export(writer document.Writer) error {
	switch typed := v.content.(type) {
	case string:
		if err := writer.WriteString(typed); err != nil {
			return fmt.Errorf("%w", err)
		}
	default:
		return fmt.Errorf("%w: %T", ErrUnexpectedValueType, typed)
	}

	return nil
}

func (v *Value) Import(value any) error {
	v.content = value

	return nil
}

func (v *Value) AsPrimitive() any {
	return v.content
}

func (v *Value) Reset() {
	v.content = nil
}
