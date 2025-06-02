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
)

var (
	ErrInvalidComp3Sign = errors.New("invalid COMP-3 sign nibble")
	ErrBufferTooShort   = errors.New("buffer too short for COMP-3 encoding")
)

type Comp3 struct {
	intDigits int
	decDigits int
	size      int
}

func NewComp3(intDigits, decDigits int) *Comp3 {
	return &Comp3{
		intDigits: intDigits,
		decDigits: decDigits,
		size:      (intDigits + decDigits + 2) / 2, //nolint:mnd
	}
}

const (
	nibbleShift    = 4
	highNibbleMask = 0xF0
	lowNibbleMask  = 0x0F

	signNibblePositive = 0xC
	signNibbleNegative = 0xD
	signNibbleZero     = 0xF
)

func (c *Comp3) Decode(buffer api.Buffer, offset int) (any, error) {
	result := &strings.Builder{}

	bytes, err := buffer.Slice(offset, c.size)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("%w", err) // return error if buffer is too short
	}

	for byteIndex, byteVal := range bytes {
		if byteIndex*2 == c.intDigits {
			result.WriteRune('.')
		}

		if byteIndex == c.size-1 {
			high := (byteVal & highNibbleMask) >> nibbleShift

			if byteIndex*2 < c.intDigits+c.decDigits {
				result.WriteRune(convertNibleToRune(high))
			}

			sign, err := handleSign(byteVal)
			if err != nil {
				return result.String(), err
			}

			return sign + result.String(), nil
		}

		result.WriteRune(convertNibleToRune((byteVal & highNibbleMask) >> nibbleShift))

		if byteIndex*2+1 == c.intDigits {
			result.WriteRune('.')
		}

		result.WriteRune(convertNibleToRune(byteVal & lowNibbleMask))
	}

	return result.String(), ErrBufferTooShort // should never happen because handled by buffer interface
}

func (c *Comp3) Encode(buffer api.Buffer, offset int, value any) error {
	panic("Comp3 encoding not implemented yet") // Placeholder for actual implementation
}

func (c *Comp3) Size() int {
	return c.size
}

const (
	maxDecimalNibble = 9
	minAlphaNibble   = 10
	maxAlphaNibble   = 15
)

func convertNibleToRune(nible byte) rune {
	if nible <= maxDecimalNibble {
		return rune('0' + nible)
	}

	if nible >= minAlphaNibble && nible <= maxAlphaNibble {
		return rune('A' + (nible - minAlphaNibble))
	}

	return '?'
}

func handleSign(byteVal byte) (string, error) {
	signNibble := byteVal & lowNibbleMask

	if signNibble != signNibblePositive && signNibble != signNibbleNegative && signNibble != signNibbleZero {
		return "", fmt.Errorf("%w: 0x%X", ErrInvalidComp3Sign, signNibble)
	}

	switch signNibble {
	case signNibbleNegative:
		return "-", nil
	case signNibblePositive:
		return "+", nil
	}

	return "", nil
}
