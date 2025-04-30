package data

type Builder struct {
	pos     int
	indexes map[string]int
}

func NewBuilder() *Builder {
	return &Builder{
		pos:     0,
		indexes: make(map[string]int),
	}
}

func (b *Builder) Build(schema RecordSchema) *Object {
	return b.build(schema, nil)
}

func (b *Builder) buildArray(schema RecordSchema, occurs int, when ExportPredicate) *Array {
	array := NewArray(when)
	for range occurs {
		array.Add(b.build(schema, nil))
	}

	return array
}

func (b *Builder) buildArrayScalar(length, occurs int, when ExportPredicate, trim bool) *Array {
	array := NewArray(when)
	for range occurs {
		array.Add(NewValue(b.pos, length, nil, trim))
		b.pos += length
	}

	return array
}

func (b *Builder) build(schema RecordSchema, when ExportPredicate) *Object {
	object := NewObject(when)

	for _, field := range schema {
		if field.Redefine != "" {
			if pos, ok := b.indexes[field.Redefine]; ok {
				b.pos = pos
			}
		}

		b.indexes[field.Name] = b.pos

		switch {
		case field.Occurs == 0 && field.Schema != nil:
			object.Add(field.Name, b.build(field.Schema, field.When))
		case field.Occurs > 0 && field.Schema != nil:
			object.Add(field.Name, b.buildArray(field.Schema, field.Occurs, field.When))
		case field.Occurs == 0:
			object.Add(field.Name, NewValue(b.pos, field.Length, field.When, field.Trim))
			b.pos += field.Length
		case field.Occurs > 0:
			object.Add(field.Name, b.buildArrayScalar(field.Length, field.Occurs, field.When, field.Trim))
		}
	}

	return object
}
