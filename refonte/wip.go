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

func (s *String) Unmarshal(node *Node, data Buffer) {
	working := make([]byte, utf8.UTFMax)

	for idx := range s.length {
		raw := data[node.end : node.end+utf8.UTFMax]

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

type Decoder interface {
	Unmarshal(d *Node, data Buffer)
	Get() any
	Set(value any)
	fmt.Stringer
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

func (n *Node) Then(next *Node) *Node {
	n.next = append(n.next, *next)
	next.prev = n

	return next
}

func (n *Node) Unmarshal(data Buffer) {
	if n.prev != nil {
		n.prev.Unmarshal(data)
		n.start = n.prev.end
		n.end = n.start
	}

	n.decoder.Unmarshal(n, data)
}

func (n *Node) Get() any {
	return n.decoder.Get()
}

func (n *Node) Set(value any) {
	n.decoder.Set(value)
}

func (n *Node) String() string {
	if n.prev != nil {
		return n.prev.String() + "/" + n.decoder.String()
	}

	return n.decoder.String()
}
