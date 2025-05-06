package jsonline

import (
	"bufio"
	"fmt"
	"io"

	"github.com/cgi-fr/posimap/refonte/api"
)

type Writer struct {
	writer *bufio.Writer
}

func NewWriter(writer io.Writer) *Writer {
	return &Writer{
		writer: bufio.NewWriter(writer),
	}
}

func (w *Writer) Open() error {
	return nil
}

func (w *Writer) Close() error {
	if err := w.writer.Flush(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

//nolint:cyclop
func (w *Writer) WriteToken(token api.StructToken) error {
	switch token { //nolint:exhaustive
	case api.StructTokenObjectStart:
		if _, err := w.writer.WriteRune('{'); err != nil {
			return fmt.Errorf("%w", err)
		}
	case api.StructTokenObjectEnd:
		if _, err := w.writer.WriteRune('}'); err != nil {
			return fmt.Errorf("%w", err)
		}
	case api.StructTokenArrayStart:
		if _, err := w.writer.WriteRune('['); err != nil {
			return fmt.Errorf("%w", err)
		}
	case api.StructTokenArrayEnd:
		if _, err := w.writer.WriteRune(']'); err != nil {
			return fmt.Errorf("%w", err)
		}
	case api.StructTokenSeparator:
		if _, err := w.writer.WriteRune(','); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (w *Writer) WriteString(data string) error {
	if _, err := w.writer.WriteRune('"'); err != nil {
		return fmt.Errorf("%w", err)
	}

	if _, err := w.writer.WriteString(data); err != nil {
		return fmt.Errorf("%w", err)
	}

	if _, err := w.writer.WriteRune('"'); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (w *Writer) WriteKey(key string) error {
	if err := w.WriteString(key); err != nil {
		return fmt.Errorf("%w", err)
	}

	if _, err := w.writer.WriteRune(':'); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
