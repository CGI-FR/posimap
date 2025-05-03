package decoder

type Node interface {
	Unmarshal(data Buffer)
	Chain(next Node) Node
	_state() *nodeState
}

type nodeState struct {
	prev Node
	next []Node

	start int
	end   int
}
