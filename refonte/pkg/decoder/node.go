package decoder

import "fmt"

type Node interface {
	Unmarshal(data Buffer)
	getEnd() int
	setPrev(node Node)
	fmt.Stringer
}
