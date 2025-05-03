package decoder

import "fmt"

type Node interface {
	Unmarshal(data Buffer)
	State() *NodeState
	fmt.Stringer
}

type NodeState struct {
	prev Node
	next []Node

	start int
	end   int
}
