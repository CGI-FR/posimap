package record

import (
	"fmt"

	"github.com/cgi-fr/posimap/refonte/api"
)

type Value struct {
	offset  int
	decoder api.Decoder
	encoder api.Encoder
	export  api.Predicate
}

func NewValue(offset int, codec api.Codec, export api.Predicate) *Value {
	return &Value{
		offset:  offset,
		decoder: codec,
		encoder: codec,
		export:  export,
	}
}

func (v *Value) Unmarshal(buffer api.Buffer) (any, error) {
	decoded, err := v.decoder.Decode(buffer, v.offset)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return decoded, nil
}

func (v *Value) Marshal(buffer api.Buffer, value any) error {
	if err := v.encoder.Encode(buffer, v.offset, value); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (v *Value) Export(buffer api.Buffer, context any, writer api.StructWriter) error {
	if v.export != nil && !v.export(context) {
		return nil
	}

	value, err := v.Unmarshal(buffer)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	str, ok := value.(string)
	if !ok {
		panic("expected string value")
	}

	if err := writer.WriteString(str); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (v *Value) Import(buffer api.Buffer, reader api.StructReader) (any, error) {
	str, err := reader.ReadString()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if err := v.Marshal(buffer, str); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return str, nil
}
