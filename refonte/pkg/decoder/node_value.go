package decoder

import (
	"fmt"
	"strings"
)

type NodeValue struct {
	prev Node
	next []Node

	start int
	end   int

	decoder Decoder

	element any
}

func NewNode(name string, decoder Decoder) *NodeValue {
	return &NodeValue{
		prev:    nil,
		next:    nil,
		start:   0,
		end:     0,
		decoder: decoder,
		element: nil,
	}
}

func (n *NodeValue) Chain(node Node) Node { //nolint:ireturn
	n.next = append(n.next, node)
	node.setPrev(n)

	return node
}

func (n *NodeValue) getEnd() int {
	return n.end
}

func (n *NodeValue) setPrev(node Node) {
	n.prev = node
}

func (n *NodeValue) Unmarshal(data Buffer) {
	if n.prev != nil {
		n.prev.Unmarshal(data)
		n.start = n.prev.getEnd()
		n.end = n.start
	}

	var size int

	n.element, size = n.decoder.Unmarshal(n, data, n.end)

	n.end += size
}

func (n *NodeValue) String() string {
	myself := strings.Builder{}

	myself.WriteRune('"')
	myself.WriteString(fmt.Sprint(n.element))
	myself.WriteRune('"')

	return myself.String()
}
