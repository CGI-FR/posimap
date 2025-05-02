package flat

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"golang.org/x/text/encoding/charmap"
)

// SourceFixedWidth implements the Source interface for fixed-width encodings.
type SourceFixedWidth struct {
	reader  *bufio.Reader
	charmap *charmap.Charmap
}

func NewSourceFixedWidth(reader io.Reader, encoding *charmap.Charmap) *SourceFixedWidth {
	if encoding == nil {
		encoding = charmap.ISO8859_1
	}

	return &SourceFixedWidth{
		reader:  bufio.NewReader(reader),
		charmap: encoding,
	}
}

func (s *SourceFixedWidth) ReadRunes(length int) ([]rune, error) {
	runes := make([]rune, 0, length)

	for len(runes) < length {
		raw, err := s.reader.Peek(1) // raw will contain at most 1 byte
		if err != nil {
			return runes, fmt.Errorf("%w", err)
		}

		if len(raw) == 0 {
			return runes, io.EOF
		}

		r := s.charmap.DecodeByte(raw[0])
		runes = append(runes, r)

		// Discard the bytes that were read from the source.
		_, _ = s.reader.Discard(1)
	}

	return runes, nil
}

func (s *SourceFixedWidth) ReadBytes(length int) ([]byte, error) {
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

func (s *SourceFixedWidth) IsFixedWidth() bool {
	return true
}
