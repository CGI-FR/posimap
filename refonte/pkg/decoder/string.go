package decoder

import (
	"unicode/utf8"

	"golang.org/x/text/encoding"
)

type String struct {
	decoder *encoding.Decoder
	length  int
	value   []rune
}

func NewString(encoding encoding.Encoding, length int) *Node {
	return &Node{
		prev:    nil,
		next:    nil,
		start:   0,
		end:     0,
		decoder: &String{encoding.NewDecoder(), length, make([]rune, length)},
	}
}

func (s *String) Unmarshal(node *Node, data Buffer) {
	working := make([]byte, utf8.UTFMax)

	for idx := range s.length {
		raw := data.Peek(node.end, utf8.UTFMax)

		nDst, _, _ := s.decoder.Transform(working, raw, false)

		r, size := utf8.DecodeRune(working[:nDst])
		s.value[idx] = r

		node.end += size
	}
}

func (s *String) Get() any {
	return string(s.value)
}

func (s *String) Set(value any) {
	str, ok := value.(string)
	if !ok {
		panic("value must be a string")
	}

	for idx, r := range str {
		s.value[idx] = r

		if len(s.value) == s.length {
			break
		}
	}

	for range len(str) - len(s.value) {
		s.value = append(s.value, ' ')
	}
}

func (s *String) String() string {
	return string(s.value)
}
