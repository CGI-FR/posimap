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
	codec   api.Codec[any]
	content any
}

func NewValue(offset int, codec api.Codec[any]) *Value {
	return &Value{
		offset:  offset,
		codec:   codec,
		content: nil,
	}
}

func (v *Value) Unmarshal(buffer api.Buffer) error {
	var err error

	v.content, err = v.codec.Decode(buffer, v.offset)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (v *Value) Marshal(buffer api.Buffer) error {
	if v.content == nil {
		return nil // document did not have the key set
	}

	err := v.codec.Encode(buffer, v.offset, v.content)
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

func (v *Value) Import(value any) error {
	v.content = value

	return nil
}

func (v *Value) AsPrimitive() any {
	return v.content
}

func (v *Value) Reset() {
	v.content = nil
}
