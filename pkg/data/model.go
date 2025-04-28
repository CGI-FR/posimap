package data

import "fmt"

type View interface {
	// Materialize the buffer as primitive types, this operation consumes memory.
	Materialize(buffer Buffer) any
	// Export the buffer into a record sink.
	Export(root View, buffer Buffer, sink ObjectSink) error
	// ShouldExport returns true if the view will export data with the given context.
	ShouldExport(root View, buffer Buffer) bool
}

type Record struct {
	buffer Buffer
	root   View
}

func NewRecord(buffer Buffer, view View) Record {
	return Record{buffer: buffer, root: view}
}

func (r Record) Export(sink ObjectSink) error {
	if err := r.root.Export(r.root, r.buffer, sink); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
