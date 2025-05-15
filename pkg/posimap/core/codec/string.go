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

package codec

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/cgi-fr/posimap/pkg/posimap/api"
	"golang.org/x/text/encoding/charmap"
)

var (
	ErrCannotEncodeRune = errors.New("cannot encode rune")
	ErrExpectedString   = errors.New("expected string value")
)

type String struct {
	charmap *charmap.Charmap
	length  int
	trim    bool
}

func NewString(charmap *charmap.Charmap, length int, trim bool) *String {
	return &String{
		charmap: charmap,
		length:  length,
		trim:    trim,
	}
}

func (s *String) Decode(buffer api.Buffer, offset int) (any, error) {
	runes := make([]rune, 0, s.length)

	bytes, err := buffer.Slice(offset, s.length)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("%w", err)
	}

	for _, b := range bytes {
		runes = append(runes, s.charmap.DecodeByte(b))
	}

	if s.trim {
		return strings.TrimRight(string(runes), VisibleBlankRunes), err
	}

	return string(runes), err
}

func (s *String) Encode(buffer api.Buffer, offset int, value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%w: got %T", ErrExpectedString, value)
	}

	bytes := make([]byte, 0, s.length)

	for _, rune := range str {
		if len(bytes) == s.length {
			break
		}

		b, ok := s.charmap.EncodeRune(rune)
		if !ok {
			return fmt.Errorf("%w: rune %q to %s", ErrCannotEncodeRune, rune, s.charmap.String())
		}

		bytes = append(bytes, b)
	}

	if err := buffer.Write(offset, bytes); err != nil {
		return fmt.Errorf("%w", err)
	}

	if s.length-len(bytes) > 0 {
		space, _ := s.charmap.EncodeRune(' ')
		for idx := range s.length - len(bytes) {
			if err := buffer.Write(offset+len(bytes)+idx, []byte{space}); err != nil {
				return fmt.Errorf("%w", err)
			}
		}
	}

	return nil
}

func (s *String) Size() int {
	return s.length
}
