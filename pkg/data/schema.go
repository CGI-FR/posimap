package data

type FieldSchema struct {
	Name     string
	Length   int
	Occurs   int
	Redefine string
	Trim     bool
	When     ExportPredicate
	Schema   RecordSchema
}

type RecordSchema []FieldSchema
