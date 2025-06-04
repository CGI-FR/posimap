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
	ErrInvalidComp3Sign    = errors.New("invalid COMP-3 sign nibble")
	ErrInvalidComp3Nibble  = errors.New("invalid COMP-3 digit")
	ErrInvalidComp3String  = errors.New("invalid COMP-3 string")
	ErrMisplacedDecimalSep = errors.New("mislaced decimal separator in COMP-3 encoding")
	ErrBufferTooShort      = errors.New("buffer too short for COMP-3 encoding")
)

type Comp3 struct {
	intDigits int
	decDigits int
	signed    bool
	size      int // number of bytes needed to store the COMP-3 value
	length    int // number of characters in the string representation (not counting the sign)
}

func NewComp3(intDigits, decDigits int, signed bool) *Comp3 {
	length := intDigits + decDigits
	if decDigits > 0 {
		length++
	}

	return &Comp3{
		intDigits: intDigits,
		decDigits: decDigits,
		signed:    signed,
		size:      (intDigits + decDigits + 2) / 2, //nolint:mnd
		length:    length,
	}
}

const (
	nibbleShift    = 4
	highNibbleMask = byte(0xF0)
	lowNibbleMask  = byte(0x0F)

	signNibblePositive = byte(0xC)
	signNibbleNegative = byte(0xD)
	signNibbleZero     = byte(0xF)
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
				result.WriteRune(convertNibbleToRune(high))
			}

			sign, err := handleSign(byteVal)
			if err != nil {
				return result.String(), err
			}

			return sign + result.String(), nil
		}

		result.WriteRune(convertNibbleToRune((byteVal & highNibbleMask) >> nibbleShift))

		if byteIndex*2+1 == c.intDigits {
			result.WriteRune('.')
		}

		result.WriteRune(convertNibbleToRune(byteVal & lowNibbleMask))
	}

	return result.String(), ErrBufferTooShort // should never happen because short buffer is handled by buffer interface
}

func (c *Comp3) Encode(buffer api.Buffer, offset int, value any) error {
	nibbleSign, str, err := c.detectSignAndAddLeadingZeroes(value)
	if err != nil {
		return err
	}

	if c.decDigits > 0 && str[c.intDigits] != '.' {
		return fmt.Errorf("%w: expected decimal separator at position %d, got %q",
			ErrMisplacedDecimalSep, c.intDigits, str[c.intDigits])
	}

	// ensure that there is only one decimal separator
	if strings.Count(str, ".") > 1 {
		return fmt.Errorf("%w: too many decimal separators in COMP-3 encoding", ErrMisplacedDecimalSep)
	}

	return c.encode(buffer, offset, str, nibbleSign)
}

func (c *Comp3) detectSignAndAddLeadingZeroes(value any) (byte, string, error) {
	str, ok := value.(string)
	if !ok {
		return signNibbleZero, "", fmt.Errorf("%w: got %T", ErrExpectedString, value)
	}

	if len(str) == 0 {
		return signNibbleZero, "", fmt.Errorf("%w: empty string cannot be encoded in COMP-3", ErrInvalidComp3String)
	}

	nibbleSign := signNibbleZero

	if c.signed {
		nibbleSign = signNibblePositive // default to positive sign
	}

	switch str[0] {
	case '+':
		nibbleSign = signNibblePositive
		str = str[1:] // remove sign character
	case '-':
		nibbleSign = signNibbleNegative
		str = str[1:] // remove sign character
	}

	if len(str) < c.length {
		// add leading zeros if the string is too short
		str = strings.Repeat("0", c.length-len(str)) + str
	} else if len(str) != c.length {
		return 0, "", fmt.Errorf("%w: expected %d characters, got %d", ErrInvalidComp3String, c.length, len(str))
	}

	return nibbleSign, str, nil
}

func (c *Comp3) encode(buffer api.Buffer, offset int, str string, nibbleSign byte) error {
	var byteVal byte

	ignoreCount := 0

	for charIndex, char := range str {
		if charIndex-ignoreCount >= c.intDigits+c.decDigits+1 {
			return fmt.Errorf("%w: too many characters in COMP-3 encoding", ErrBufferTooShort)
		}

		if char == '.' {
			if charIndex-ignoreCount != c.intDigits {
				return fmt.Errorf("%w", ErrMisplacedDecimalSep)
			}

			ignoreCount++

			continue
		}

		nibble, err := convertRuneToNibble(char)
		if err != nil {
			return err
		}

		if (charIndex-ignoreCount)%2 == 0 {
			byteVal |= (nibble << nibbleShift) // high nibble
		} else {
			byteVal |= nibble // low nibble
			if (charIndex - ignoreCount) == c.intDigits+c.decDigits {
				// if we are at the last nibble, set the sign nibble
				byteVal |= nibbleSign
			}

			if err := buffer.Write(offset+(charIndex-ignoreCount)/2, []byte{byteVal}); err != nil {
				return fmt.Errorf("%w", err)
			}

			byteVal = 0 // reset for next byte
		}
	}

	// if we have an odd number of nibbles, we need to write the last byte
	if err := buffer.Write(offset+c.size-1, []byte{byteVal | nibbleSign}); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (c *Comp3) Size() int {
	return c.size
}

const (
	maxDecimalNibble  = 9
	minAlphaNibble    = 10
	maxAlphaNibble    = 15
	alphaNibbleOffset = 10
)

func convertNibbleToRune(nibble byte) rune {
	if nibble <= maxDecimalNibble {
		return rune('0' + nibble)
	}

	if nibble >= minAlphaNibble && nibble <= maxAlphaNibble {
		return rune('A' + (nibble - minAlphaNibble))
	}

	return '?'
}

func convertRuneToNibble(runeVal rune) (byte, error) {
	if runeVal >= '0' && runeVal <= '9' {
		return byte(runeVal - '0'), nil
	}

	if runeVal >= 'A' && runeVal <= 'F' {
		return byte(runeVal - 'A' + alphaNibbleOffset), nil
	}

	return 0, fmt.Errorf("%w: invalid rune %q", ErrInvalidComp3Nibble, runeVal)
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
