package record

import (
	"fmt"

	"github.com/cgi-fr/posimap/refonte/api"
)

type Named struct {
	name   string
	record api.Record
	export api.Predicate
}

func NewNamed(name string, record api.Record, export api.Predicate) Named {
	return Named{
		name:   name,
		record: record,
		export: export,
	}
}

func (n Named) Unmarshal(buffer api.Buffer, obj map[string]any) error {
	decoded, err := n.record.Unmarshal(buffer)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	obj[n.name] = decoded

	return nil
}

func (n Named) Marshal(buffer api.Buffer, obj map[string]any) error {
	value, exists := obj[n.name]
	if !exists {
		return nil
	}

	if err := n.record.Marshal(buffer, value); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
