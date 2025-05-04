package decoder

import "iter"

type Node interface {
	Unmarshal(data Buffer)
	_state() *nodeState
}

type Nodes []Node

type nodeState struct {
	prev *Nodes
	next *Nodes

	start int
	end   int
}

func (s *nodeState) addNext(next Node) {
	*s.next = append(*s.next, next)
}

func (s *nodeState) addPrev(prev Node) {
	*s.prev = append(*s.prev, prev)
}

func NodesOf(nodes ...Node) *Nodes {
	res := make(Nodes, len(nodes))
	copy(res, nodes)

	return &res
}

func (set *Nodes) Nodes() iter.Seq2[int, Node] {
	if set == nil {
		return func(_ func(idx int, node Node) bool) {}
	}

	return func(yield func(idx int, node Node) bool) {
		for idx, node := range *set {
			if !yield(idx, node) {
				return
			}
		}
	}
}

func (set *Nodes) Chain(next ...Node) {
	if set == nil {
		return
	}

	nodes := make(Nodes, len(*set))
	copy(nodes, *set)

	for _, node := range *set {
		node._state().next = &nodes
	}

	for _, node := range next {
		node._state().prev = set
	}
}
