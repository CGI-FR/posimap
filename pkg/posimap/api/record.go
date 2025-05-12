package api

import "github.com/cgi-fr/posimap/pkg/posimap/driven/document"

type Record interface {
	Unmarshal(buffer Buffer) error
	Marshal(buffer Buffer) error
	Export(writer document.Writer) error
	Import(value any) error
	Reset()
	AsPrimitive() any
}
