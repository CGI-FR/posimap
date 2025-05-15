// Copyright (C) 2025 CGI France
//
// This file is part of posimap.
//
// posimap is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// posimap is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with posimap.  If not, see <http://www.gnu.org/licenses/>.

package jsonline

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/cgi-fr/posimap/pkg/posimap/driven/document"
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

func (w *Writer) shift() error {
	sep := w.pointer.Shift()
	if sep != 0 {
		if _, err := w.writer.WriteRune(sep); err != nil {
			return fmt.Errorf("%w", err)
		}

		if sep == '\n' {
			if err := w.writer.Flush(); err != nil {
				return fmt.Errorf("%w", err)
			}
		}
	}

	return nil
}

func (w *Writer) handleSeparator(token document.Token) error {
	switch token {
	case
		document.TokenObjStart,
		document.TokenArrStart,
		document.TokenString,
		document.TokenNumber,
		document.TokenTrue,
		document.TokenFalse,
		document.TokenNull:
		return w.shift()
	case document.TokenObjEnd, document.TokenArrEnd:
		return nil
	}

	return fmt.Errorf("%w: %v", ErrUnknownToken, token)
}

func (w *Writer) handleLevel(token document.Token) (bool, error) {
	switch token { //nolint:exhaustive
	case document.TokenObjStart:
		w.pointer.OpenObject()
	case document.TokenObjEnd:
		if err := w.pointer.CloseObject(); err != nil {
			return false, fmt.Errorf("%w", err)
		}
	case document.TokenArrStart:
		w.pointer.OpenArray()
	case document.TokenArrEnd:
		if err := w.pointer.CloseArray(); err != nil {
			return false, fmt.Errorf("%w", err)
		}
	default:
		return false, nil
	}

	if _, err := w.writer.WriteRune(rune(token)); err != nil {
		return false, fmt.Errorf("%w", err)
	}

	return true, nil
}

func (w *Writer) handleString(value any) error {
	if str, ok := value.(string); ok {
		if err := w.WriteString(str); err != nil {
			return fmt.Errorf("%w", err)
		}
	} else {
		return fmt.Errorf("%w: expected string, got %T", ErrUnexpectedType, value)
	}

	return nil
}

func (w *Writer) handleNumber(value any) error {
	if num, ok := value.(float64); ok {
		if err := w.WriteNumber(num); err != nil {
			return fmt.Errorf("%w", err)
		}
	} else {
		return fmt.Errorf("%w: expected float64, got %T", ErrUnexpectedType, value)
	}

	return nil
}

func (w *Writer) handleBool(value any) error {
	if boolVal, ok := value.(bool); ok {
		if err := w.WriteBool(boolVal); err != nil {
			return fmt.Errorf("%w", err)
		}
	} else {
		return fmt.Errorf("%w: expected bool, got %T", ErrUnexpectedType, value)
	}

	return nil
}

func (w *Writer) WriteValue(token document.Token, value any) error {
	switch token {
	case document.TokenString:
		return w.handleString(value)
	case document.TokenNumber:
		return w.handleNumber(value)
	case document.TokenTrue:
		return w.handleBool(value)
	case document.TokenFalse:
		return w.handleBool(value)
	case document.TokenNull:
		return w.WriteNull()
	case document.TokenObjStart, document.TokenObjEnd, document.TokenArrStart, document.TokenArrEnd:
		return w.WriteToken(token)
	default:
		return fmt.Errorf("%w: %v", ErrUnknownToken, token)
	}
}

func (w *Writer) WriteToken(token document.Token) error {
	if err := w.handleSeparator(token); err != nil {
		return fmt.Errorf("%w", err)
	}

	if handled, err := w.handleLevel(token); err != nil {
		return fmt.Errorf("%w", err)
	} else if handled {
		return nil
	}

	return w.WriteValue(token, nil)
}

func (w *Writer) WriteString(data string) error {
	if err := w.shift(); err != nil {
		return err
	}

	// when encoding/json/v2 is ready to use, switch to it
	// for now, we use a simple escape from encoding/json

	escapedBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if _, err := w.writer.Write(escapedBytes); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (w *Writer) WriteNumber(value float64) error {
	if err := w.shift(); err != nil {
		return err
	}

	if _, err := w.writer.WriteString(fmt.Sprintf("%f", value)); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (w *Writer) WriteBool(value bool) error {
	if err := w.shift(); err != nil {
		return err
	}

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
	if err := w.shift(); err != nil {
		return err
	}

	if _, err := w.writer.WriteString("null"); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (w *Writer) Close() error {
	if _, err := w.writer.WriteRune('\n'); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := w.writer.Flush(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
