package record

import (
	"errors"
	"fmt"

	"github.com/cgi-fr/posimap/pkg/posimap/api"
	"github.com/cgi-fr/posimap/pkg/posimap/driven/document"
)

var ErrUnexpectedValueType = errors.New("unexpected value type")

type Value struct {
	offset  int
	decoder api.Decoder
	encoder api.Encoder
	content any
}

func NewValue(offset int, codec api.Codec) *Value {
	return &Value{
		offset:  offset,
		decoder: codec,
		encoder: codec,
		content: nil,
	}
}

func (v *Value) Unmarshal(buffer api.Buffer) error {
	var err error

	v.content, err = v.decoder.Decode(buffer, v.offset)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (v *Value) Marshal(buffer api.Buffer) error {
	if v.content == nil {
		return nil // document did not have the key set
	}

	err := v.encoder.Encode(buffer, v.offset, v.content)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (v *Value) export(writer document.Writer, _ Record) error {
	return v.Export(writer)
}

func (v *Value) Export(writer document.Writer) error {
	switch typed := v.content.(type) {
	case string:
		if err := writer.WriteString(typed); err != nil {
			return fmt.Errorf("%w", err)
		}
	default:
		return fmt.Errorf("%w: %T", ErrUnexpectedValueType, typed)
	}

	return nil
}

func (v *Value) Import(reader document.Reader) error {
	var err error

	_, v.content, err = reader.ReadValue()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (v *Value) AsPrimitive() any {
	return v.content
}
