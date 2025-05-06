package jsonline

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/cgi-fr/posimap/refonte/driven/document"
)

var (
	ErrTokenNeedValue = errors.New("token need a value")
	ErrUnknownToken   = errors.New("unknown token")
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
func (w *Writer) WriteToken(token document.Token) error {
	var err error

	switch token {
	case document.TokenDocSep:
		if _, err = w.writer.WriteRune('\n'); err == nil {
			err = w.writer.Flush()
		}
	case document.TokenObjStart:
		_, err = w.writer.WriteRune('{')
	case document.TokenObjEnd:
		_, err = w.writer.WriteRune('}')
	case document.TokenArrStart:
		_, err = w.writer.WriteRune('[')
	case document.TokenArrEnd:
		_, err = w.writer.WriteRune(']')
	case document.TokenValSep:
		_, err = w.writer.WriteRune(',')
	case document.TokenTrue:
		err = w.WriteBool(true)
	case document.TokenFalse:
		err = w.WriteBool(false)
	case document.TokenNull:
		err = w.WriteNull()
	case document.TokenEOF:
		err = w.writer.Flush()
	case document.TokenKey:
	case document.TokenString:
	case document.TokenNumber:
		err = fmt.Errorf("%w: %v", ErrTokenNeedValue, token)
	default:
		err = fmt.Errorf("%w: %v", ErrUnknownToken, token)
	}

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

//nolint:cyclop
func (w *Writer) WriteValue(token document.Token, value any) error {
	switch token {
	case document.TokenKey:
		if str, ok := value.(string); ok {
			return w.WriteKey(str)
		}
	case document.TokenString:
		if str, ok := value.(string); ok {
			return w.WriteString(str)
		}
	case document.TokenNumber:
		if num, ok := value.(float64); ok {
			return w.WriteNumber(num)
		}
	case document.TokenDocSep:
	case document.TokenObjStart:
	case document.TokenObjEnd:
	case document.TokenArrStart:
	case document.TokenArrEnd:
	case document.TokenValSep:
	case document.TokenTrue:
	case document.TokenFalse:
	case document.TokenNull:
	case document.TokenEOF:
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
