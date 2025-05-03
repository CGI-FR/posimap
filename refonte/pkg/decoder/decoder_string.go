package decoder

import (
	"unicode/utf8"

	"golang.org/x/text/encoding"
)

type String struct {
	decoder *encoding.Decoder
	length  int
}

func NewDecoderString(encoding encoding.Encoding, length int) *String {
	return &String{encoding.NewDecoder(), length}
}

func (s *String) Unmarshal(node Node, data Buffer, offset int) (any, int) {
	working := make([]byte, utf8.UTFMax)

	value := make([]rune, s.length)

	nread := 0

	for idx := range s.length {
		raw := data.Peek(offset+nread, utf8.UTFMax)

		nDst, _, _ := s.decoder.Transform(working, raw, false)

		r, size := utf8.DecodeRune(working[:nDst])
		value[idx] = r

		nread += size
	}

	return value, nread
}
