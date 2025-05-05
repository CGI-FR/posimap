package record

import (
	"errors"
	"fmt"

	"github.com/cgi-fr/posimap/refonte/api"
)

var ErrUnexpectedObjectType = errors.New("unexpected object type")

type Object struct {
	offset  int
	keys    []string
	records map[string]api.Record
	export  api.Predicate
}

func NewObject(offset int, export api.Predicate) *Object {
	return &Object{
		offset:  offset,
		keys:    make([]string, 0),
		records: make(map[string]api.Record),
		export:  export,
	}
}

func (o *Object) Add(key string, record api.Record) {
	o.keys = append(o.keys, key)
	o.records[key] = record
}

func (o *Object) Unmarshal(buffer api.Buffer) (any, error) {
	decoded := make(map[string]any)

	for _, key := range o.keys {
		record := o.records[key]

		value, err := record.Unmarshal(buffer)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		decoded[key] = value
	}

	return decoded, nil
}

func (o *Object) Marshal(buffer api.Buffer, value any) error {
	decoded, ok := value.(map[string]any)
	if !ok {
		return fmt.Errorf("%w: %T", ErrUnexpectedObjectType, value)
	}

	for _, key := range o.keys {
		record := o.records[key]

		val, exists := decoded[key]
		if !exists {
			continue
		}

		if err := record.Marshal(buffer, val); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (o *Object) Export(buffer api.Buffer, context any, writer api.StructWriter) error {
	if o.export != nil && !o.export(context) {
		return nil
	}

	if err := writer.WriteToken(api.StructTokenObjectStart); err != nil {
		return fmt.Errorf("%w", err)
	}

	for idx, key := range o.keys {
		record := o.records[key]

		if idx > 0 {
			if err := writer.WriteToken(api.StructTokenSeparator); err != nil {
				return fmt.Errorf("%w", err)
			}
		}

		if err := writer.WriteKey(key); err != nil {
			return fmt.Errorf("%w", err)
		}

		if err := record.Export(buffer, context, writer); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	if err := writer.WriteToken(api.StructTokenObjectEnd); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
