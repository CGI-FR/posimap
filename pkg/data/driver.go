package data

import (
	"errors"
	"fmt"
	"io"
)

const DefaultBufferSize = 16 * 1024

func TransformRecordsToObjects(root View, source RecordSource, sink ObjectSink) error {
	for {
		buffer, err := source.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return fmt.Errorf("failed to read record: %w", err)
		}

		if buffer == nil {
			continue
		}

		if err := sink.OpenRecord(); err != nil {
			return fmt.Errorf("failed to open record: %w", err)
		}

		if err := root.Export(root, buffer, sink); err != nil {
			return fmt.Errorf("failed to write record: %w", err)
		}

		if err := sink.CloseRecord(); err != nil {
			return fmt.Errorf("failed to close record: %w", err)
		}
	}

	if err := sink.Close(); err != nil {
		return fmt.Errorf("failed to close sink: %w", err)
	}

	if err := source.Close(); err != nil {
		return fmt.Errorf("failed to close source: %w", err)
	}

	return nil
}

func TransformObjectsToRecords(root View, source ObjectSource, sink RecordSink) error {
	for {
		obj, err := source.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return fmt.Errorf("failed to read object: %w", err)
		}

		if obj == nil {
			continue
		}

		buffer := make(Buffer, DefaultBufferSize)

		if err := root.Import(obj, buffer); err != nil {
			return fmt.Errorf("failed to import record: %w", err)
		}

		if err := sink.Write(buffer); err != nil {
			return fmt.Errorf("failed to write record: %w", err)
		}
	}

	if err := sink.Close(); err != nil {
		return fmt.Errorf("failed to close sink: %w", err)
	}

	if err := source.Close(); err != nil {
		return fmt.Errorf("failed to close source: %w", err)
	}

	return nil
}
