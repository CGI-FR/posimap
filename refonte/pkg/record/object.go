package record

import (
	"errors"
	"fmt"

	"github.com/cgi-fr/posimap/refonte/api"
)

var ErrUnexpectedObjectType = errors.New("unexpected object type")

type Object struct {
	records []Named
}

func NewObject() *Object {
	return &Object{
		records: make([]Named, 0),
	}
}

func (o *Object) Add(key string, record api.Record, export api.Predicate) {
	o.records = append(o.records, NewNamed(key, record, export))
}

func (o *Object) Unmarshal(buffer api.Buffer) error {
	for _, record := range o.records {
		if err := record.Unmarshal(buffer); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (o *Object) Marshal(buffer api.Buffer) error {
	for _, record := range o.records {
		if err := record.Marshal(buffer); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}
