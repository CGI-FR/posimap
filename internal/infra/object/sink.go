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
