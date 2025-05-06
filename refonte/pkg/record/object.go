package record

import (
	"errors"
	"fmt"

	"github.com/cgi-fr/posimap/refonte/api"
)

var (
	ErrUnexpectedTokenType = errors.New("unexpected token type")
	ErrUnexpectedKey       = errors.New("unexpected key")
)

type Object struct {
	keys    []string
	records map[string]api.Record
	exports map[string]api.Predicate
}

func NewObject() *Object {
	return &Object{
		keys:    make([]string, 0),
		records: make(map[string]api.Record),
		exports: make(map[string]api.Predicate),
	}
}

func (o *Object) Add(key string, record api.Record, export api.Predicate) {
	o.keys = append(o.keys, key)
	o.records[key] = record
	o.exports[key] = export
}

func (o *Object) Unmarshal(buffer api.Buffer) error {
	for _, key := range o.keys {
		record := o.records[key]

		if err := record.Unmarshal(buffer); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (o *Object) Marshal(buffer api.Buffer) error {
	for _, key := range o.keys {
		record := o.records[key]

		if err := record.Marshal(buffer); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

//nolint:cyclop
func (o *Object) Export(writer api.StructWriter, feedback ...api.Record) error {
	if err := writer.WriteToken(api.StructTokenObjectStart); err != nil {
		return fmt.Errorf("%w", err)
	}

	if len(feedback) == 0 {
		feedback = []api.Record{o}
	}

	first := true

	for _, key := range o.keys {
		record := o.records[key]

		if export, ok := o.exports[key]; ok && export != nil && len(feedback) > 0 && !export(feedback[0]) {
			// Skip the record if the export predicate returns false
			continue
		}

		if !first {
			if err := writer.WriteToken(api.StructTokenSeparator); err != nil {
				return fmt.Errorf("%w", err)
			}
		} else {
			first = false
		}

		if err := writer.WriteKey(key); err != nil {
			return fmt.Errorf("%w", err)
		}

		if err := record.Export(writer); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	if err := writer.WriteToken(api.StructTokenObjectEnd); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (o *Object) Import(reader api.StructReader) error {
	if token, err := reader.ReadToken(); err != nil {
		return fmt.Errorf("%w", err)
	} else if token != api.StructTokenObjectStart {
		return fmt.Errorf("%w: %q, expected %q", ErrUnexpectedTokenType, token, api.StructTokenObjectStart)
	}

	for {
		if token, err := reader.ReadToken(); err != nil {
			return fmt.Errorf("%w", err)
		} else if token == api.StructTokenObjectEnd {
			break
		} else if token != api.StructTokenKey {
			return fmt.Errorf("%w: %q, expected %q", ErrUnexpectedTokenType, token, api.StructTokenKey)
		}

		key, err := reader.ReadKey()
		if err != nil {
			return fmt.Errorf("%w", err)
		} else if _, ok := o.records[key]; !ok {
			return fmt.Errorf("%w: %s", ErrUnexpectedKey, key)
		}

		record := o.records[key]

		if err := record.Import(reader); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (o *Object) AsPrimitive() any {
	primitive := make(map[string]any, len(o.keys))

	for _, key := range o.keys {
		record := o.records[key]

		primitive[key] = record.AsPrimitive()
	}

	return primitive
}
