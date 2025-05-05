package codec

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/cgi-fr/posimap/refonte/api"
	"golang.org/x/text/encoding/charmap"
)

var (
	ErrCannotEncodeRune = errors.New("cannot encode rune")
	ErrExpectedString   = errors.New("expected string value")
)

type String struct {
	charmap charmap.Charmap
	length  int
	trim    bool
}

func NewString(charmap charmap.Charmap, length int, trim bool) *String {
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
		return strings.TrimRight(string(runes), BlankRunes), err
	}

	return string(runes), err
}

func (s *String) Encode(buffer api.Buffer, offset int, value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%w: got %T", ErrExpectedString, value)
	}

	bytes := make([]byte, 0, s.length)

	for idx, rune := range str {
		if idx == s.length {
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

	return nil
}
