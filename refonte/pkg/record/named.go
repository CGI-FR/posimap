package record

import (
	"errors"
	"fmt"

	"github.com/cgi-fr/posimap/refonte/api"
)

var ErrUnexpectedKey = errors.New("unexpected key")

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

func (n Named) Unmarshal(buffer api.Buffer) error {
	err := n.record.Unmarshal(buffer)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (n Named) Marshal(buffer api.Buffer) error {
	if err := n.record.Marshal(buffer); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (n Named) Export(writer api.StructWriter) error {
	if err := writer.WriteString(n.name); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := n.record.Export(writer); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (n Named) Import(reader api.StructReader) error {
	if key, err := reader.ReadKey(); err != nil {
		return fmt.Errorf("%w", err)
	} else if key != n.name {
		return fmt.Errorf("%w: %s", ErrUnexpectedKey, key)
	}

	if err := n.record.Import(reader); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
