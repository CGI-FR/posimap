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

package schema

import (
	"errors"
	"fmt"

	"github.com/cgi-fr/posimap/pkg/posimap/core/record"
)

var ErrInvalidRedefines = errors.New("invalid redefines")

func (r *Record) Build() (*record.Object, error) {
	rec := record.NewObject()
	offset := 0

	if err := r.build(rec, &offset); err != nil {
		return nil, err
	}

	return rec, nil
}

func (r *Record) build(rec *record.Object, offset *int) error {
	redefines := make(map[string]int)

	if r.feedback {
		rec.SetFeedback()
	}

	for _, node := range r.children {
		if err := node.build(rec, offset, redefines); err != nil {
			return err
		}
	}

	return nil
}

func (n *node) updateOffsetForRedefines(offset *int, redefines map[string]int) error {
	if n.redefines == "" {
		return nil
	}

	// Check if the field redefines another field and adjust the offset accordingly.
	if pos, ok := redefines[n.redefines]; ok {
		*offset = pos

		return nil
	}

	return fmt.Errorf("%w: %s", ErrInvalidRedefines, n.redefines)
}

func (n *node) build(rec *record.Object, offset *int, redefines map[string]int) error {
	if err := n.updateOffsetForRedefines(offset, redefines); err != nil {
		return err
	}

	redefines[n.name] = *offset

	switch {
	case n.occurs == 0 && n.element.IsCodec():
		n.buildCodec(rec, offset)
	case n.occurs > 0 && n.element.IsCodec():
		n.buildCodecArray(rec, offset)
	case n.occurs == 0 && n.element.IsSchema():
		return n.buildSchema(rec, offset)
	case n.occurs > 0 && n.element.IsSchema():
		return n.buildSchemaArray(rec, offset)
	}

	return nil
}

func (n *node) buildCodec(rec *record.Object, offset *int) {
	rec.Add(n.name, record.NewValue(*offset, n.element.Codec()), n.when)

	*offset += n.element.Codec().Size()
}

func (n *node) buildCodecArray(rec *record.Object, offset *int) {
	array := record.NewArray()
	for range n.occurs {
		array.Add(record.NewValue(*offset, n.element.Codec()))
		*offset += n.element.Codec().Size()
	}

	rec.Add(n.name, array, n.when)
}

func (n *node) buildSchema(rec *record.Object, offset *int) error {
	sub := record.NewObject()
	if err := n.element.Schema().build(sub, offset); err != nil {
		return err
	}

	rec.Add(n.name, sub, n.when)

	return nil
}

func (n *node) buildSchemaArray(rec *record.Object, offset *int) error {
	array := record.NewArray()

	for range n.occurs {
		sub := record.NewObject()
		if err := n.element.Schema().build(sub, offset); err != nil {
			return err
		}

		array.Add(sub)
	}

	rec.Add(n.name, array, n.when)

	return nil
}
