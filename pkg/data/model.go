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

package data

import "fmt"

type View interface {
	// Materialize the buffer as primitive types, this operation consumes memory.
	Materialize(buffer *Buffer) any
	// Export the buffer into a record sink.
	Export(root View, buffer *Buffer, sink ObjectSink) error
	// ShouldExport returns true if the view will export data with the given context.
	ShouldExport(root View, buffer *Buffer) bool

	Import(value any, buffer *Buffer) error
}

type Record struct {
	buffer *Buffer
	root   View
}

func NewRecord(buffer *Buffer, view View) Record {
	return Record{buffer: buffer, root: view}
}

func (r Record) Export(sink ObjectSink) error {
	if err := r.root.Export(r.root, r.buffer, sink); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
