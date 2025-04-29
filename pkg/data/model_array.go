package data

import "fmt"

type Array struct {
	items    []View
	exported ExportPredicate
}

func NewArray(predicate ExportPredicate) *Array {
	return &Array{
		items:    make([]View, 0),
		exported: predicate,
	}
}

func (a *Array) Materialize(buffer *Buffer) any {
	result := make([]any, 0, len(a.items))
	for _, item := range a.items {
		result = append(result, item.Materialize(buffer))
	}

	return result
}

func (a *Array) Export(root View, buffer *Buffer, sink ObjectSink) error {
	if err := sink.OpenArray(); err != nil {
		return fmt.Errorf("%w", err)
	}

	for idx, item := range a.items {
		if !a.items[idx].ShouldExport(root, buffer) {
			continue
		}

		if idx != 0 {
			if err := sink.Next(); err != nil {
				return fmt.Errorf("%w", err)
			}
		}

		if err := item.Export(root, buffer, sink); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	if err := sink.CloseArray(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (a *Array) ShouldExport(root View, buffer *Buffer) bool {
	return a.exported == nil || a.exported(root, buffer)
}

func (a *Array) SetExport(predicate ExportPredicate) {
	a.exported = predicate
}

func (a *Array) Add(item View) {
	a.items = append(a.items, item)
}

func (a *Array) Import(value any, buffer *Buffer) error {
	switch typed := value.(type) {
	case []any:
		for idx, val := range typed {
			if err := a.items[idx].Import(val, buffer); err != nil {
				return fmt.Errorf("%w", err)
			}
		}
	default:
		return fmt.Errorf("%w: expected array, got %T", ErrInvalidType, typed)
	}

	return nil
}
