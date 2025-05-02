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

package copybook

import "github.com/cgi-fr/posimap/pkg/data2"

type Builder struct {
	pos     int
	indexes map[string]int
}

func NewBuilder() *Builder {
	return &Builder{
		pos:     0,
		indexes: make(map[string]int),
	}
}

func (b *Builder) Build(schema RecordSchema) *data2.SchemaObject {
	return b.build(schema, nil)
}

func (b *Builder) build(schema RecordSchema, when data2.Predicate) *data2.SchemaObject {
	object := data2.NewSchemaObject(when)

	for _, field := range schema {
		switch {
		case field.Occurs == 0 && field.Schema != nil:
			object.Add(field.Name, b.build(field.Schema, field.When), field.Redefine)
		case field.Occurs > 0 && field.Schema != nil:
			object.Add(field.Name, data2.NewSchemaArray(b.build(field.Schema, nil), field.Occurs, field.When), field.Redefine)
		case field.Occurs == 0:
			object.Add(field.Name, data2.NewSchemaValue(field.Length, field.Trim, field.When), field.Redefine)
		case field.Occurs > 0:
			object.Add(field.Name, data2.NewSchemaArray(data2.NewSchemaValue(field.Length, field.Trim, nil), field.Occurs, field.When), field.Redefine) //nolint:lll
		}
	}

	return object
}
