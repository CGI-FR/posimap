package data2

import "github.com/cgi-fr/posimap/pkg/deep"

type Record interface {
	Export(sink deep.Sink) error
	Materialize() any
}
