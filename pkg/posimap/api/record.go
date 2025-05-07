package api

import "github.com/cgi-fr/posimap/pkg/posimap/driven/document"

type Record interface {
	Unmarshal(buffer Buffer) error
	Marshal(buffer Buffer) error
	Export(writer document.Writer) error
	Import(reader document.Reader) error
	AsPrimitive() any
}
