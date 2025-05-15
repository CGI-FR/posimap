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

package charsets

import (
	"errors"
	"fmt"
)

var (
	ErrUnsupportedCharset       = errors.New("unsupported charset")
	ErrUnsupportedRuneInCharset = errors.New("unsupported rune in charset")
)

func GetByteInCharset(charset string, char rune) (byte, error) {
	charmap, err := Get(charset)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrUnsupportedCharset, charset)
	}

	b, ok := charmap.EncodeRune(char)
	if !ok {
		return 0, fmt.Errorf("%w: %q in %s", ErrUnsupportedRuneInCharset, char, charset)
	}

	return b, nil
}

func GetBytesInCharset(charset string, value string) ([]byte, error) {
	charmap, err := Get(charset)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedCharset, charset)
	}

	bytes := make([]byte, len(value))

	for i, char := range value {
		var ok bool
		if bytes[i], ok = charmap.EncodeRune(char); !ok {
			return nil, fmt.Errorf("%w: %q in %s", ErrUnsupportedRuneInCharset, char, charset)
		}
	}

	return bytes, nil
}
