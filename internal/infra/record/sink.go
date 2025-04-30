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

package record

import (
	"bufio"
	"fmt"
	"io"

	"github.com/cgi-fr/posimap/pkg/data"
)

type Sink struct {
	writer *bufio.Writer
}

func NewRecordSink(writer io.Writer) Sink {
	return Sink{
		writer: bufio.NewWriter(writer),
	}
}

func (s Sink) Write(b *data.Buffer) error {
	if _, err := s.writer.WriteString(b.String()); err != nil {
		return fmt.Errorf("%w", err)
	}

	if _, err := s.writer.WriteRune('\n'); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := s.writer.Flush(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (s Sink) Close() error {
	if err := s.writer.Flush(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
