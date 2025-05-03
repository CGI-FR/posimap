package decoder

import (
	"fmt"
	"strings"
)

type NodeValue struct {
	state *nodeState

	decoder Decoder

	element any
}

func NewNode(decoder Decoder) *NodeValue {
	return &NodeValue{
		state:   &nodeState{}, //nolint:exhaustruct
		decoder: decoder,
		element: nil,
	}
}

func (n *NodeValue) Chain(next Node) Node { //nolint:ireturn
	n.state.next = append(n.state.next, next)
	next._state().prev = n

	return next
}

func (n *NodeValue) _state() *nodeState {
	return n.state
}

func (n *NodeValue) Unmarshal(data Buffer) {
	if n.state.prev != nil {
		n.state.prev.Unmarshal(data)
		n.state.start = n.state.prev._state().end
		n.state.end = n.state.start
	}

	var size int

	n.element, size = n.decoder.Unmarshal(data, n.state.end)

	n.state.end += size
}

func (n *NodeValue) String() string {
	myself := strings.Builder{}

	myself.WriteRune('"')
	myself.WriteString(fmt.Sprint(n.element))
	myself.WriteRune('"')

	return myself.String()
}
