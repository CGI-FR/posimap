package decoder

import (
	"fmt"
	"strings"
)

type NodeKeyed struct {
	state *NodeState

	keys   []string
	values map[string]Decoder

	element map[string]any
}

func NewNodeKeyed() *NodeKeyed {
	return &NodeKeyed{
		state:   &NodeState{}, //nolint:exhaustruct
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
	next.State().prev = n

	return next
}

func (n *NodeKeyed) State() *NodeState {
	return n.state
}

func (n *NodeKeyed) Unmarshal(data Buffer) {
	if n.state.prev != nil {
		n.state.prev.Unmarshal(data)
		n.state.start = n.state.prev.State().end
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
