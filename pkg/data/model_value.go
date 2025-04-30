package data

import "fmt"

type Value struct {
	start    int
	length   int
	trim     bool
	exported ExportPredicate
}

func NewValue(start, length int, predicate ExportPredicate, trim bool) Value {
	if predicate == nil {
		predicate = If(length > 0)
	}

	return Value{
		start:    start,
		length:   length,
		trim:     trim,
		exported: predicate,
	}
}

func (v Value) Materialize(buffer *Buffer) any {
	if v.trim {
		return buffer.ReadTrimmed(v.start, v.length, BlankRunes)
	}

	return buffer.Read(v.start, v.length)
}

func (v Value) Export(_ View, buffer *Buffer, sink ObjectSink) error {
	if v.trim {
		if err := sink.WriteString(buffer.ReadTrimmed(v.start, v.length, BlankRunes)); err != nil {
			return fmt.Errorf("%w", err)
		}
	} else {
		if err := sink.WriteString(buffer.Read(v.start, v.length)); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (v Value) ShouldExport(root View, buffer *Buffer) bool {
	return v.exported == nil || v.exported(root, buffer)
}

func (v Value) Import(value any, buffer *Buffer) error {
	switch typed := value.(type) {
	case string:
		if err := buffer.Write(v.start, v.length, typed); err != nil {
			return fmt.Errorf("%w", err)
		}
	default:
		return fmt.Errorf("%w: expected string, got %T", ErrInvalidType, typed)
	}

	return nil
}
