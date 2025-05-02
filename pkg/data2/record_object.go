package data2

import (
	"fmt"

	"github.com/cgi-fr/posimap/pkg/deep"
)

type RecordObject struct {
	schema  *SchemaObject
	records map[string]Record
	export  Predicate
}

func (ro RecordObject) Materialize() any {
	result := make(map[string]any)
	for _, name := range ro.schema.keys {
		result[name] = ro.records[name].Materialize()
	}

	return result
}

func (ro RecordObject) Export(root Record, sink deep.Sink) error {
	if err := sink.OpenObject(); err != nil {
		return fmt.Errorf("%w", err)
	}

	first := true

	for _, name := range ro.schema.keys {
		if !ro.records[name].VisibleFrom(root) {
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

		if err := ro.records[name].Export(root, sink); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	if err := sink.CloseObject(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (ro RecordObject) VisibleFrom(root Record) bool {
	return ro.export == nil || ro.export(root)
}
