package jsonline

import (
	"bufio"
	"fmt"
	"io"

	"github.com/cgi-fr/posimap/refonte/driven/document"
)

type Writer struct {
	writer  *bufio.Writer
	pointer *Pointer
}

func NewWriter(writer io.Writer) *Writer {
	return &Writer{
		writer:  bufio.NewWriter(writer),
		pointer: NewPointer(),
	}
}

//nolint:cyclop,funlen
func (w *Writer) WriteValue(token document.Token, value any) error {
	var err error

	//nolint:exhaustive
	switch token {
	case document.TokenObjStart:
		w.pointer.OpenObject()
	case document.TokenArrStart:
		w.pointer.OpenArray()
	case document.TokenObjEnd:
		if err := w.pointer.CloseObject(); err != nil {
			return fmt.Errorf("%w", err)
		}
	case document.TokenArrEnd:
		if err := w.pointer.CloseArray(); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	if sep := w.pointer.Shift(); sep != 0 {
		if _, err := w.writer.WriteRune(sep); err != nil {
			return fmt.Errorf("%w", err)
		}

		if sep == '\n' {
			if err := w.writer.Flush(); err != nil {
				return fmt.Errorf("%w", err)
			}
		}
	}

	switch token {
	case document.TokenObjStart:
		_, err = w.writer.WriteRune('{')
	case document.TokenObjEnd:
		_, err = w.writer.WriteRune('}')
	case document.TokenArrStart:
		_, err = w.writer.WriteRune('[')
	case document.TokenArrEnd:
		_, err = w.writer.WriteRune(']')
	case document.TokenTrue:
		err = w.WriteBool(true)
	case document.TokenFalse:
		err = w.WriteBool(false)
	case document.TokenNull:
		err = w.WriteNull()
	case document.TokenString:
		if str, ok := value.(string); ok {
			err = w.WriteString(str)
		} else {
			err = fmt.Errorf("%w: %T", ErrUnexpectedType, value)
		}
	case document.TokenNumber:
		if num, ok := value.(float64); ok {
			err = w.WriteNumber(num)
		} else {
			err = fmt.Errorf("%w: %T", ErrUnexpectedType, value)
		}
	default:
		err = fmt.Errorf("%w: %v", ErrUnknownToken, token)
	}

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (w *Writer) WriteToken(token document.Token) error {
	switch token {
	case document.TokenString:
		return fmt.Errorf("%w: %v", ErrTokenNeedValue, token)
	case document.TokenNumber:
		return fmt.Errorf("%w: %v", ErrTokenNeedValue, token)
	case document.TokenObjStart:
	case document.TokenObjEnd:
	case document.TokenArrStart:
	case document.TokenArrEnd:
	case document.TokenTrue:
	case document.TokenFalse:
	case document.TokenNull:
		return w.WriteToken(token)
	}

	return fmt.Errorf("%w: %q", ErrUnknownToken, token)
}

func (w *Writer) WriteKey(key string) error {
	if err := w.WriteString(key); err != nil {
		return err
	}

	if _, err := w.writer.WriteRune(':'); err != nil {
		return fmt.Errorf("%w", err)
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

func (w *Writer) WriteNumber(value float64) error {
	if _, err := w.writer.WriteString(fmt.Sprintf("%f", value)); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (w *Writer) WriteBool(value bool) error {
	if value {
		if _, err := w.writer.WriteString("true"); err != nil {
			return fmt.Errorf("%w", err)
		}
	} else {
		if _, err := w.writer.WriteString("false"); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

func (w *Writer) WriteNull() error {
	if _, err := w.writer.WriteString("null"); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (w *Writer) WriteEOF() error {
	if err := w.writer.Flush(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
