package data2

import "github.com/cgi-fr/posimap/pkg/flat"

type Schema interface {
	RuneCount() int
	ReadBuffer(source flat.Source, buffer *Buffer) error
	CreateRecord(buffer Buffer, start int) (Record, error)
}
