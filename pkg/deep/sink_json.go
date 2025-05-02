package deep

import (
	"bufio"
	"fmt"
	"io"
)

type SinkJSONLine struct {
	writer *bufio.Writer
}

func NewSinkJSONLine(writer io.Writer) *SinkJSONLine {
	return &SinkJSONLine{
		writer: bufio.NewWriter(writer),
	}
}

func (m *SinkJSONLine) OpenRecord() error {
	return nil
}

func (m *SinkJSONLine) CloseRecord() error {
	if _, err := m.writer.WriteRune('\n'); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := m.writer.Flush(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (m *SinkJSONLine) OpenObject() error {
	if _, err := m.writer.WriteRune('{'); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (m *SinkJSONLine) CloseObject() error {
	if _, err := m.writer.WriteRune('}'); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (m *SinkJSONLine) OpenArray() error {
	if _, err := m.writer.WriteRune('['); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (m *SinkJSONLine) CloseArray() error {
	if _, err := m.writer.WriteRune(']'); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (m *SinkJSONLine) WriteString(data string) error {
	if _, err := m.writer.WriteRune('"'); err != nil {
		return fmt.Errorf("%w", err)
	}

	if _, err := m.writer.WriteString(data); err != nil {
		return fmt.Errorf("%w", err)
	}

	if _, err := m.writer.WriteRune('"'); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (m *SinkJSONLine) WriteKey(key string) error {
	if err := m.WriteString(key); err != nil {
		return err
	}

	if _, err := m.writer.WriteRune(':'); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (m *SinkJSONLine) Next() error {
	if _, err := m.writer.WriteRune(','); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (m *SinkJSONLine) Close() error {
	if err := m.writer.Flush(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
