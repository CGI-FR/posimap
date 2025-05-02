package data2

import "github.com/cgi-fr/posimap/pkg/deep"

type Record interface {
	Export(root Record, sink deep.Sink) error
	Materialize() any
	VisibleFrom(root Record) bool
}
