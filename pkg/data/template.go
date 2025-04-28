package data

type FieldSchema struct {
	Name     string
	Length   int
	Occurs   int
	Redefine string
	When     ExportPredicate
	Schema   RecordSchema
}

type RecordSchema []FieldSchema
