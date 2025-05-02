package data2

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"unicode/utf8"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

type Source struct {
	reader  *bufio.Reader
	encoder *encoding.Encoder
	decoder *encoding.Decoder
	working []byte // working buffer for decoding
}

func NewSource(reader io.Reader, encoding encoding.Encoding) *Source {
	if encoding == nil {
		encoding = unicode.UTF8
	}

	return &Source{
		reader:  bufio.NewReader(reader),
		encoder: encoding.NewEncoder(),
		decoder: encoding.NewDecoder(),
		working: make([]byte, utf8.UTFMax),
	}
}

func (s *Source) ReadRunes(length int) ([]rune, error) {
	runes := make([]rune, 0, length)

	for len(runes) < length {
		// Read a maximum of utf8.UTFMax bytes from the reader to decode a single rune.
		raw, err := s.reader.Peek(utf8.UTFMax) // raw will contain at most 4 runes, at least 1
		if err != nil && !errors.Is(err, io.EOF) {
			return runes, fmt.Errorf("%w", err)
		}

		if len(raw) == 0 {
			return runes, io.EOF
		}

		nDst, _, _ := s.decoder.Transform(s.working, raw, false)

		// Decode the bytes into 1 rune.
		r, _ := utf8.DecodeRune(s.working[:nDst])
		runes = append(runes, r)

		// Reencode the rune back to bytes to compute len to discard.
		size, _, _ := s.encoder.Transform(s.working, []byte(string(r)), true)

		// Discard the bytes that were read from the source.
		_, _ = s.reader.Discard(size)
	}

	return runes, nil
}

func (s *Source) ReadBytes(length int) ([]byte, error) {
	result := make([]byte, length)

	n, err := s.reader.Read(result)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return result[:n], io.EOF
		}

		return nil, fmt.Errorf("%w", err)
	}

	return result, nil
}
