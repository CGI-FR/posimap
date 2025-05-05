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

func (o *Object) Unmarshal(buffer api.Buffer) (any, error) {
	decoded := make(map[string]any)

	for _, record := range o.records {
		if err := record.Unmarshal(buffer, decoded); err != nil {
			return nil, fmt.Errorf("%w", err)
		}
	}

	return decoded, nil
}

func (o *Object) Marshal(buffer api.Buffer, value any) error {
	decoded, ok := value.(map[string]any)
	if !ok {
		return fmt.Errorf("%w: %T", ErrUnexpectedObjectType, value)
	}

	for _, record := range o.records {
		if err := record.Marshal(buffer, decoded); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}
