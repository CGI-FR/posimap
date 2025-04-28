package data

import "fmt"

type Value struct {
	start    int
	length   int
	exported ExportPredicate
}

func NewValue(start, length int, predicate ExportPredicate) Value {
	if predicate == nil {
		predicate = If(length > 0)
	}

	return Value{
		start:    start,
		length:   length,
		exported: predicate,
	}
}

func (v Value) Materialize(buffer Buffer) any {
	return buffer.Read(v.start, v.length).String()
}

func (v Value) Export(_ View, buffer Buffer, sink ObjectSink) error {
	if err := sink.WriteString(buffer.Read(v.start, v.length).String()); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (v Value) ShouldExport(root View, buffer Buffer) bool {
	return v.exported == nil || v.exported(root, buffer)
}

func (v Value) Import(value any, buffer Buffer) error {
	switch typed := value.(type) {
	case string:
		if err := buffer.Write(v.start, v.length, typed); err != nil {
			return fmt.Errorf("%w", err)
		}
	default:
		return nil
	}

	return nil
}
