package record

import (
	"fmt"

	"github.com/cgi-fr/posimap/refonte/api"
	"github.com/cgi-fr/posimap/refonte/pkg/stoken"
)

type Array struct {
	records []Record
}

func NewArray() *Array {
	return &Array{
		records: make([]Record, 0),
	}
}

func (a *Array) Add(record Record) {
	a.records = append(a.records, record)
}

func (a *Array) Unmarshal(buffer api.Buffer) error {
	for _, record := range a.records {
		if err := record.Unmarshal(buffer); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (a *Array) Marshal(buffer api.Buffer) error {
	for _, record := range a.records {
		if err := record.Marshal(buffer); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (a *Array) Export(writer api.StructWriter) error {
	return a.export(writer, a)
}

func (a *Array) export(writer api.StructWriter, _ Record) error {
	if err := writer.WriteToken(stoken.ArrayStart); err != nil {
		return fmt.Errorf("%w", err)
	}

	for idx, record := range a.records {
		if idx > 0 {
			if err := writer.WriteToken(stoken.Separator); err != nil {
				return fmt.Errorf("%w", err)
			}
		}

		if err := record.Export(writer); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	if err := writer.WriteToken(stoken.ArrayEnd); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (a *Array) Import(reader api.StructReader) error {
	if token, err := reader.ReadToken(); err != nil {
		return fmt.Errorf("%w", err)
	} else if token != stoken.ArrayStart {
		return fmt.Errorf("%w: %q, expected %q", ErrUnexpectedTokenType, token, stoken.ArrayStart)
	}

	for _, record := range a.records {
		if err := record.Import(reader); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	if token, err := reader.ReadToken(); err != nil {
		return fmt.Errorf("%w", err)
	} else if token != stoken.ArrayEnd {
		return fmt.Errorf("%w: %q, expected %q", ErrUnexpectedTokenType, token, stoken.ArrayEnd)
	}

	return nil
}

func (a *Array) AsPrimitive() any {
	primitive := make([]any, len(a.records))

	for idx, record := range a.records {
		primitive[idx] = record.AsPrimitive()
	}

	return primitive
}
