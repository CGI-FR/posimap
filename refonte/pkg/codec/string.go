package codec

import (
	"errors"
	"fmt"
	"io"

	"github.com/cgi-fr/posimap/refonte/api"
	"golang.org/x/text/encoding/charmap"
)

var (
	ErrCannotEncodeRune = errors.New("cannot encode rune")
	ErrExpectedString   = errors.New("expected string value")
)

type String struct {
	charmap charmap.Charmap
	offset  int
	length  int
}

func NewString(charmap charmap.Charmap, offset int, length int) *String {
	return &String{
		charmap: charmap,
		offset:  offset,
		length:  length,
	}
}

func (s *String) Decode(buffer api.Buffer) (any, error) {
	runes := make([]rune, 0, s.length)

	bytes, err := buffer.Slice(s.offset, s.length)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("%w", err)
	}

	for _, b := range bytes {
		runes = append(runes, s.charmap.DecodeByte(b))
	}

	return string(runes), err
}

func (s *String) Encode(buffer api.Buffer, value any) error {
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

	if err := buffer.Write(s.offset, bytes); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
