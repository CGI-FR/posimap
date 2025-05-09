package schema2

import (
	"fmt"

	"github.com/cgi-fr/posimap/pkg/posimap/api"
)

type Node struct {
	name      string
	redefines string
	occurs    int
	when      api.Predicate
	codec     api.Codec[any]

	redefined map[string]*Node
	children  []*Node
	dependsOn []*Node
}

func NewSchema() *Node {
	return &Node{
		name:      "R",
		redefines: "",
		occurs:    0,
		when:      nil,
		codec:     nil,
		redefined: make(map[string]*Node),
		children:  make([]*Node, 0),
		dependsOn: nil,
	}
}

func (n *Node) SetDependsOn(other *Node) {
	fmt.Println("SetDependsOn", n.name, "->", other.name)
	backup := n.dependsOn
	n.dependsOn = []*Node{other}
	other.dependsOn = backup
}

func (n *Node) AddDependsOn(other *Node) {
	fmt.Println("AddDependsOn", n.name, "->", other.name, "{", other.redefines, "}")
	n.dependsOn = append(n.dependsOn, other)
	dependent := n.redefined[other.name]
	other.dependsOn = dependent.dependsOn
}

func (n *Node) AddChild(other *Node) {
	n.children = append(n.children, other)
	n.redefined[other.name] = other
}

func (n *Node) AddRedefine(other *Node) {
	n.children = append(n.children, other)
	n.redefined[other.name] = n.redefined[other.redefines]
}

func (n *Node) WithField(name string, codec api.Codec[any], options ...Option) *Node {
	field := NewSchema()
	field.name = name
	field.codec = codec

	for _, option := range options {
		field = option(field)
	}

	if field.redefines != "" {
		n.AddRedefine(field)
	} else {
		n.AddChild(field)
	}

	return n
}

func (n *Node) WithRecord(name string, record *Node, options ...Option) *Node {
	record.name = name

	for _, option := range options {
		record = option(record)
	}

	if record.redefines != "" {
		n.AddRedefine(record)
	} else {
		n.AddChild(record)
	}

	return n
}

func (n *Node) Print() {
	for _, dependent := range n.dependsOn {
		println(fmt.Sprintf("\"%s\"", n.name), "->", fmt.Sprintf("\"%s\"", dependent.name))
	}
	for _, child := range n.children {
		child.Print()
	}
}

func (n *Node) FindTips() []*Node {
	tips := make([]*Node, 0)

	for _, dependent := range n.dependsOn {
		if len(dependent.dependsOn) == 0 {
			tips = append(tips, dependent)
		} else {
			tips = append(tips, dependent.FindTips()...)
		}
	}

	return tips
}

type Option func(*Node) *Node

func Occurs(occurs int) Option {
	return func(node *Node) *Node {
		node.occurs = occurs

		return node
	}
}

func Condition(when api.Predicate) Option {
	return func(node *Node) *Node {
		node.when = when

		return node
	}
}

func Redefines(redefines string) Option {
	return func(node *Node) *Node {
		node.redefines = redefines

		return node
	}
}

func CompileDependsOn(node *Node) {
	for _, child := range node.children {
		if child.redefines != "" {
			node.AddDependsOn(child)
		} else {
			node.SetDependsOn(child)
		}
	}
	for _, child := range node.children {
		CompileDependsOn(child)
	}
}
