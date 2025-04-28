package data

type FieldTemplate struct {
	Name     string
	Length   int
	Occurs   int
	Redefine string
	When     ExportPredicate
	Template RecordTemplate
}

type RecordTemplate []FieldTemplate
