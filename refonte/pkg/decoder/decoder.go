package decoder

import "fmt"

type Decoder interface {
	Unmarshal(d *Node, data Buffer)
	Get() any
	Set(value any)
	fmt.Stringer
}
