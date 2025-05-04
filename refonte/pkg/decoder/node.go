package decoder

type Node interface {
	Unmarshal(data Buffer)
	_state() *nodeState
}

type Nodes []Node

type nodeState struct {
	prev Nodes
	next Nodes

	start int
	end   int
}

func (set Nodes) Chain(next ...Node) {
	for _, node := range set {
		node._state().next = next
	}

	for _, node := range next {
		node._state().prev = set
	}
}
