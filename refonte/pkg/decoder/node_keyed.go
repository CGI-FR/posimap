package decoder

import (
	"fmt"
	"iter"
	"strings"
)

type NodeKeyed struct {
	state *nodeState

	keys   []string
	values map[string]Decoder

	element map[string]any
}

func NewNodeKeyed() *NodeKeyed {
	return &NodeKeyed{
		state:   &nodeState{}, //nolint:exhaustruct
		keys:    nil,
		values:  make(map[string]Decoder),
		element: make(map[string]any),
	}
}

func (n *NodeKeyed) Add(key string, value Decoder) {
	n.keys = append(n.keys, key)
	n.values[key] = value
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

	var size int

	for _, key := range n.keys {
		dec := n.values[key]
		n.element[key], size = dec.Unmarshal(data, n.state.end)
		n.state.end += size
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

// Implement the Keyed interface

func (n *NodeKeyed) ValueForKey(key string) (any, bool) {
	value, has := n.element[key]

	return value, has
}

func (n *NodeKeyed) HasKey(key string) bool {
	_, has := n.element[key]

	return has
}

func (n *NodeKeyed) Keys() []string {
	return n.keys
}

func (n *NodeKeyed) KeyValuePairs() iter.Seq2[string, any] {
	return func(yield func(string, any) bool) {
		for _, key := range n.keys {
			value := n.element[key]
			if !yield(key, value) {
				return
			}
		}
	}
}
