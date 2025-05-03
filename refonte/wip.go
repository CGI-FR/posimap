package refonte

import (
	"fmt"
	"unicode/utf8"

	"golang.org/x/text/encoding"
)

type Buffer []byte

type String struct {
	decoder *encoding.Decoder
	length  int
	value   []rune
}

func (s *String) Unmarshal(d *Node, data Buffer) {
	working := make([]byte, utf8.UTFMax)

	for idx := 0; idx < s.length; idx++ {
		raw := data[d.end : d.end+utf8.UTFMax]

		nDst, _, _ := s.decoder.Transform(working, raw, false)

		r, size := utf8.DecodeRune(working[:nDst])
		s.value[idx] = r

		d.end += size
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

type Decoder interface {
	Unmarshal(d *Node, data Buffer)
	Get() any
	Set(value any)
}

type Node struct {
	prev *Node
	next []Node

	start int
	end   int

	decoder Decoder
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

func (d *Node) Then(next *Node) *Node {
	d.next = append(d.next, *next)
	next.prev = d

	return next
}

func (d *Node) Unmarshal(data Buffer) {
	if d.prev != nil {
		d.prev.Unmarshal(data)
		d.start = d.prev.end
		d.end = d.start
	}

	d.decoder.Unmarshal(d, data)
}

func (d *Node) Get() any {
	return d.decoder.Get()
}

func (d *Node) Set(value any) {
	d.decoder.Set(value)
}

func (d *Node) String() string {
	if d.prev != nil {
		return d.prev.String() + "/" + fmt.Sprintf("%v", d.Get())
	}

	return fmt.Sprintf("%v", d.Get())
}
