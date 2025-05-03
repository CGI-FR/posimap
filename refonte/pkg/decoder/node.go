package decoder

type Node struct {
	prev *Node
	next []Node

	start int
	end   int

	decoder Decoder
}

func (n *Node) Then(next *Node) *Node {
	n.next = append(n.next, *next)
	next.prev = n

	return next
}

func (n *Node) Unmarshal(data Buffer) {
	if n.prev != nil {
		n.prev.Unmarshal(data)
		n.start = n.prev.end
		n.end = n.start
	}

	n.decoder.Unmarshal(n, data)
}

func (n *Node) Get() any {
	return n.decoder.Get()
}

func (n *Node) Set(value any) {
	n.decoder.Set(value)
}

func (n *Node) String() string {
	if n.prev != nil {
		return n.prev.String() + "/" + n.decoder.String()
	}

	return n.decoder.String()
}
