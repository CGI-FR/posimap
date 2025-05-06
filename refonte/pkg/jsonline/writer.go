package jsonline

import (
	"bufio"
	"fmt"
	"io"

	"github.com/cgi-fr/posimap/refonte/api"
	"github.com/cgi-fr/posimap/refonte/pkg/stoken"
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
	switch token {
	case stoken.ObjectStart:
		if _, err := w.writer.WriteRune('{'); err != nil {
			return fmt.Errorf("%w", err)
		}
	case stoken.ObjectEnd:
		if _, err := w.writer.WriteRune('}'); err != nil {
			return fmt.Errorf("%w", err)
		}
	case stoken.ArrayStart:
		if _, err := w.writer.WriteRune('['); err != nil {
			return fmt.Errorf("%w", err)
		}
	case stoken.ArrayEnd:
		if _, err := w.writer.WriteRune(']'); err != nil {
			return fmt.Errorf("%w", err)
		}
	case stoken.Separator:
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
