package decoder

import (
	"fmt"
	"strings"
)

type NodeKeyed struct {
	prev Node
	next []Node

	start int
	end   int

	keys   []string
	values map[string]Decoder

	element map[string]any
}

func NewNodeKeyed() *NodeKeyed {
	return &NodeKeyed{
		prev:    nil,
		next:    nil,
		start:   0,
		end:     0,
		keys:    nil,
		values:  make(map[string]Decoder),
		element: make(map[string]any),
	}
}

func (n *NodeKeyed) Add(key string, value Decoder) {
	n.keys = append(n.keys, key)
	n.values[key] = value
}

func (n *NodeKeyed) Chain(node Node) Node { //nolint:ireturn
	n.next = append(n.next, node)
	node.setPrev(n)

	return node
}

func (n *NodeKeyed) getEnd() int {
	return n.end
}

func (n *NodeKeyed) setPrev(node Node) {
	n.prev = node
}

func (n *NodeKeyed) Unmarshal(data Buffer) {
	if n.prev != nil {
		n.prev.Unmarshal(data)
		n.start = n.prev.getEnd()
		n.end = n.start
	}

	var size int

	for _, key := range n.keys {
		dec := n.values[key]
		n.element[key], size = dec.Unmarshal(data, n.end)
		n.end += size
	}
}

func (n *NodeKeyed) String() string {
	myself := strings.Builder{}

	myself.WriteRune('{')

	for idx, key := range n.keys {
		if idx > 0 {
			myself.WriteRune(',')
		}

		myself.WriteRune('"')
		myself.WriteString(key)
		myself.WriteRune('"')
		myself.WriteRune(':')

		myself.WriteString(fmt.Sprint(n.element[key]))
	}

	myself.WriteRune('}')

	return myself.String()
}
