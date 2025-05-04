package decoder

import (
	"fmt"
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

func (n *NodeValue) _state() *nodeState {
	return n.state
}

func (n *NodeValue) Unmarshal(data Buffer) {
	for idx, prev := range n.state.prev.Nodes() {
		prev.Unmarshal(data)

		if idx == 0 {
			n.state.start = prev._state().end
			n.state.end = n.state.start
		} else if n.state.end != prev._state().end {
			// assert n.state.end == prev._state().end
			panic("inconsistent end state")
		}
	}

	var size int

	n.element, size = n.decoder.Unmarshal(data, n.state.end)

	n.state.end += size
}

func (n *NodeValue) String() string {
	return fmt.Sprint(n.element)
}
