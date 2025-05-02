package data2

import (
	"fmt"

	"github.com/cgi-fr/posimap/pkg/deep"
)

func Export(root Record, sink deep.Sink) error {
	if err := sink.OpenRecord(); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := root.Export(root, sink); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := sink.CloseRecord(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
