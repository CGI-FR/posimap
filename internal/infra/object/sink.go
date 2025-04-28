package object

import (
	"bufio"
	"fmt"
	"io"
)

type JSON struct {
	writer *bufio.Writer
}

func NewJSON(writer io.Writer) *JSON {
	return &JSON{
		writer: bufio.NewWriter(writer),
	}
}

func (m *JSON) OpenRecord() error {
	return nil
}

func (m *JSON) CloseRecord() error {
	if _, err := m.writer.WriteRune('\n'); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := m.writer.Flush(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (m *JSON) OpenObject() error {
	if _, err := m.writer.WriteRune('{'); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (m *JSON) CloseObject() error {
	if _, err := m.writer.WriteRune('}'); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (m *JSON) OpenArray() error {
	if _, err := m.writer.WriteRune('['); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (m *JSON) CloseArray() error {
	if _, err := m.writer.WriteRune(']'); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (m *JSON) WriteString(data string) error {
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

func (m *JSON) WriteKey(key string) error {
	if err := m.WriteString(key); err != nil {
		return err
	}

	if _, err := m.writer.WriteRune(':'); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (m *JSON) Next() error {
	if _, err := m.writer.WriteRune(','); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (m *JSON) Close() error {
	if err := m.writer.Flush(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
