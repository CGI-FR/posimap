package schema

import "github.com/cgi-fr/posimap/refonte/api"

type Definition struct {
	codec  api.Codec
	schema Record
}

func NewCodec(codec api.Codec) Definition {
	return Definition{
		codec:  codec,
		schema: nil,
	}
}

func NewDefinition(schema Record) Definition {
	return Definition{
		codec:  nil,
		schema: schema,
	}
}

func (d Definition) IsCodec() bool {
	return d.codec != nil && d.schema == nil
}

func (d Definition) IsSchema() bool {
	return d.schema != nil && d.codec == nil
}

func (d Definition) Codec() api.Codec { //nolint:ireturn
	if d.IsCodec() {
		return d.codec
	}

	return nil
}

func (d Definition) Schema() Record {
	if d.IsSchema() {
		return d.schema
	}

	return nil
}
