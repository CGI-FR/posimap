package decoder

import "fmt"

type Node interface {
	Unmarshal(data Buffer)
	_state() *nodeState
	fmt.Stringer
}

type nodeState struct {
	prev Node
	next []Node

	start int
	end   int
}
