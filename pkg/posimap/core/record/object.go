package record

import (
	"errors"
	"fmt"

	"github.com/cgi-fr/posimap/pkg/posimap/api"
	"github.com/cgi-fr/posimap/pkg/posimap/driven/document"
)

var (
	ErrUnexpectedTokenType = errors.New("unexpected token type")
	ErrUnexpectedKey       = errors.New("unexpected key")
)

type Object struct {
	keys    []string
	records map[string]Record
	exports map[string]api.Predicate
}

func NewObject() *Object {
	return &Object{
		keys:    make([]string, 0),
		records: make(map[string]Record),
		exports: make(map[string]api.Predicate),
	}
}

func (o *Object) Add(key string, record Record, export api.Predicate) {
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

func (o *Object) Export(writer document.Writer) error {
	return o.export(writer, o)
}

//nolint:cyclop
func (o *Object) export(writer document.Writer, feedback Record) error {
	if err := writer.WriteToken(document.TokenObjStart); err != nil {
		return fmt.Errorf("%w", err)
	}

	for _, key := range o.keys {
		record := o.records[key]

		if export, ok := o.exports[key]; ok && export != nil && feedback != nil {
			// Skip the record if the export predicate returns false
			if ok, err := export(feedback); err != nil {
				return fmt.Errorf("%w", err)
			} else if !ok {
				continue
			}
		}

		if err := writer.WriteString(key); err != nil {
			return fmt.Errorf("%w", err)
		}

		if err := record.export(writer, feedback); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	if err := writer.WriteToken(document.TokenObjEnd); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (o *Object) Import(reader document.Reader) error {
	if token, err := reader.ReadToken(); err != nil {
		return fmt.Errorf("%w", err)
	} else if token != document.TokenObjStart {
		return fmt.Errorf("%w: %q, expected %q", ErrUnexpectedTokenType, token, document.TokenObjStart)
	}

loop:
	for {
		token, value, err := reader.ReadValue()

		switch {
		case err != nil:
			return fmt.Errorf("%w", err)
		case token == document.TokenObjEnd:
			break loop
		case token != document.TokenString:
			return fmt.Errorf("%w: %q, expected %q", ErrUnexpectedTokenType, token, document.TokenString)
		}

		key, _ := value.(string)
		if _, ok := o.records[key]; !ok {
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
