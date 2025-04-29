package data

import "fmt"

type Object struct {
	keys     []string
	items    map[string]View
	exported ExportPredicate
}

func NewObject(predicate ExportPredicate) *Object {
	return &Object{
		keys:     make([]string, 0),
		items:    make(map[string]View),
		exported: predicate,
	}
}

func (o *Object) Materialize(buffer *Buffer) any {
	result := make(map[string]any)
	for _, name := range o.keys {
		result[name] = o.items[name].Materialize(buffer)
	}

	return result
}

func (o *Object) Export(root View, buffer *Buffer, sink ObjectSink) error {
	if err := sink.OpenObject(); err != nil {
		return fmt.Errorf("%w", err)
	}

	first := true

	for _, name := range o.keys {
		if !o.items[name].ShouldExport(root, buffer) {
			continue
		}

		if !first {
			if err := sink.Next(); err != nil {
				return fmt.Errorf("%w", err)
			}
		} else {
			first = false
		}

		if err := sink.WriteKey(name); err != nil {
			return fmt.Errorf("%w", err)
		}

		if err := o.items[name].Export(root, buffer, sink); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	if err := sink.CloseObject(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (o *Object) ShouldExport(root View, buffer *Buffer) bool {
	return o.exported == nil || o.exported(root, buffer)
}

func (o *Object) SetExport(predicate ExportPredicate) {
	o.exported = predicate
}

func (o *Object) Add(name string, value View) {
	o.keys = append(o.keys, name)
	o.items[name] = value
}

func (o *Object) Import(value any, buffer *Buffer) error {
	switch typed := value.(type) {
	case map[string]any:
		for key, val := range typed {
			if err := o.items[key].Import(val, buffer); err != nil {
				return fmt.Errorf("%w", err)
			}
		}
	default:
		return nil
	}

	return nil
}
