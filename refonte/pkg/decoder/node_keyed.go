package decoder

import (
	"iter"
)

type NodeKeyed struct {
	state *nodeState

	keys      []string
	values    map[string]Node
	redefines map[string]bool
}

func NewNodeKeyed() *NodeKeyed {
	return &NodeKeyed{
		state:     &nodeState{}, //nolint:exhaustruct
		keys:      nil,
		values:    make(map[string]Node),
		redefines: make(map[string]bool),
	}
}

func (n *NodeKeyed) Add(key string, value Node) {
	n.state.prev.Chain(value)
	n.state.prev = []Node{value}
	value._state().next = []Node{n}

	n.keys = append(n.keys, key)
	n.values[key] = value
}

func (n *NodeKeyed) Redefine(key, redefine string, value Node) {
	redefined := n.values[redefine]

	value._state().prev = redefined._state().prev
	value._state().next = redefined._state().next

	n.keys = append(n.keys, key)
	n.values[key] = value
}

func (n *NodeKeyed) _state() *nodeState {
	return n.state
}

func (n *NodeKeyed) Unmarshal(data Buffer) {
	for idx, prev := range n.state.prev {
		prev.Unmarshal(data)

		if idx == 0 {
			n.state.start = prev._state().end
			n.state.end = n.state.start
		} else {
			// assert n.state.end == prev._state().end
			if n.state.end != prev._state().end {
				panic("inconsistent end state")
			}
		}
	}
}

// Implement the Keyed interface

func (n *NodeKeyed) ValueForKey(key string) (any, bool) {
	value, has := n.values[key]

	return value, has
}

func (n *NodeKeyed) HasKey(key string) bool {
	_, has := n.values[key]

	return has
}

func (n *NodeKeyed) Keys() []string {
	return n.keys
}

func (n *NodeKeyed) KeyValuePairs() iter.Seq2[string, any] {
	return func(yield func(string, any) bool) {
		for _, key := range n.keys {
			value := n.values[key]
			if !yield(key, value) {
				return
			}
		}
	}
}
