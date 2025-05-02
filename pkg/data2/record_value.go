package data2

import (
	"fmt"

	"github.com/cgi-fr/posimap/pkg/deep"
)

type RecordValue struct {
	schema SchemaValue
	buffer Buffer
	start  int
	trim   bool
}

func (rv RecordValue) Materialize() any {
	return rv.buffer.String(rv.start, rv.schema.length, rv.trim)
}

func (rv RecordValue) Export(sink deep.Sink) error {
	if err := sink.WriteString(rv.buffer.String(rv.start, rv.schema.length, rv.trim)); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
