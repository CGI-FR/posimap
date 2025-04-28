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

func (b *Builder) Build(template RecordTemplate) *Object {
	return b.build(template, nil)
}

func (b *Builder) buildArray(template RecordTemplate, occurs int, when ExportPredicate) *Array {
	array := NewArray(when)
	for range occurs {
		array.Add(b.build(template, nil))
	}

	return array
}

func (b *Builder) buildArrayScalar(length, occurs int, when ExportPredicate) *Array {
	array := NewArray(when)
	for range occurs {
		array.Add(NewValue(b.pos, length, nil))
		b.pos += length
	}

	return array
}

func (b *Builder) build(template RecordTemplate, when ExportPredicate) *Object {
	object := NewObject(when)

	for _, field := range template {
		if field.Redefine != "" {
			if pos, ok := b.indexes[field.Redefine]; ok {
				b.pos = pos
			}
		}

		b.indexes[field.Name] = b.pos

		switch {
		case field.Occurs == 0 && field.Template != nil:
			object.Add(field.Name, b.build(field.Template, field.When))
		case field.Occurs > 0 && field.Template != nil:
			object.Add(field.Name, b.buildArray(field.Template, field.Occurs, field.When))
		case field.Occurs == 0:
			object.Add(field.Name, NewValue(b.pos, field.Length, field.When))
			b.pos += field.Length
		case field.Occurs > 0:
			object.Add(field.Name, b.buildArrayScalar(field.Length, field.Occurs, field.When))
		}
	}

	return object
}
