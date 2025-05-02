package copybook

import "github.com/cgi-fr/posimap/pkg/data2"

type FieldSchema struct {
	Name     string
	Length   int
	Occurs   int
	Redefine string
	Trim     bool
	When     data2.Predicate
	Schema   RecordSchema
}

type RecordSchema []FieldSchema
