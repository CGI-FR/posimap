package decoder

import (
	"fmt"
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
	if len(n.keys) == 0 && n.state.prev != nil {
		n.state.prev.Chain(value)
	} else if len(n.keys) > 0 {
		n.values[n.keys[len(n.keys)-1]].Chain(value)
	}

	value.Chain(n)

	n.keys = append(n.keys, key)
	n.values[key] = value
}

func (n *NodeKeyed) Redefine(key, redefined string, value Node) {
	// assert redefined == last key
	if n.keys[len(n.keys)-1] != redefined {
		panic(fmt.Sprintf("redefined key %s is not the last key %s", redefined, n.keys[len(n.keys)-1]))
	}

	n.Add(key, value)
	n.redefines[key] = true
}

func (n *NodeKeyed) Chain(next Node) Node { //nolint:ireturn
	n.state.next = append(n.state.next, next)
	next._state().prev = n

	return next
}

func (n *NodeKeyed) _state() *nodeState {
	return n.state
}

func (n *NodeKeyed) Unmarshal(data Buffer) {
	if n.state.prev != nil {
		n.state.prev.Unmarshal(data)
		n.state.start = n.state.prev._state().end
		n.state.end = n.state.start
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
