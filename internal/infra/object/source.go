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
	"encoding/json"
	"fmt"
	"io"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

type JSONLineSource struct {
	scanner *bufio.Scanner
}

func NewJSONLineSource(reader io.Reader, encoding encoding.Encoding) JSONLineSource {
	scanner := bufio.NewScanner(transform.NewReader(reader, encoding.NewDecoder()))
	buf := make([]byte, 0, 64*1024)    //nolint:mnd
	scanner.Buffer(buf, 1024*1024*100) //nolint:mnd // increase buffer up to 100 MB

	return JSONLineSource{
		scanner: scanner,
	}
}

func (s JSONLineSource) Read() (any, error) {
	if !s.scanner.Scan() {
		if err := s.scanner.Err(); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return nil, io.EOF
	}

	var obj any
	if err := json.Unmarshal(s.scanner.Bytes(), &obj); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return obj, nil
}

func (s JSONLineSource) Close() error {
	if err := s.scanner.Err(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
