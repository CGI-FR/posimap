package data2

import (
	"errors"
	"fmt"
	"io"

	"github.com/cgi-fr/posimap/pkg/deep"
	"github.com/cgi-fr/posimap/pkg/flat"
)

func TransformRecordsToObjects(root Schema, source flat.Source, sink deep.Sink) error {
	for {
		buffer := NewBuffer()

		if err := root.ReadBuffer(source, &buffer); errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return fmt.Errorf("failed to read flat source: %w", err)
		}

		record, err := root.CreateRecord(buffer, 0)
		if err != nil {
			return fmt.Errorf("failed to parse record: %w", err)
		}

		if err := sink.OpenRecord(); err != nil {
			return fmt.Errorf("failed to open record: %w", err)
		}

		if err := record.Export(record, sink); err != nil {
			return fmt.Errorf("failed to write record: %w", err)
		}

		if err := sink.CloseRecord(); err != nil {
			return fmt.Errorf("failed to close record: %w", err)
		}
	}

	if err := sink.Close(); err != nil {
		return fmt.Errorf("failed to close sink: %w", err)
	}

	return nil
}
