package schema2

import (
	"fmt"

	"github.com/cgi-fr/posimap/pkg/posimap/api"
)

type node struct {
	name      string
	redefines string
	occurs    int
	when      api.Predicate

	redefined map[string]*node
	children  []*node
	dependsOn []*node
}

func (n *node) setDependsOn(other *node) {
	backup := n.dependsOn
	n.dependsOn = []*node{other}
	other.dependsOn = backup
}

func (n *node) addDependsOn(other *node) {
	n.dependsOn = append(n.dependsOn, other)
	dependent := n.redefined[other.name]
	other.dependsOn = dependent.dependsOn
}

func (n *node) addChild(other *node) {
	n.children = append(n.children, other)

	if other.redefines == "" {
		n.redefined[other.name] = other
	} else {
		n.redefined[other.name] = n.redefined[other.redefines]
	}
}

func (n *node) compileMarshalingPath() {
	for _, child := range n.children {
		if child.redefines != "" {
			n.addDependsOn(child)
		} else {
			n.setDependsOn(child)
		}
	}

	for _, child := range n.children {
		child.compileMarshalingPath()
	}
}

func (n *node) PrintGraph() {
	for _, child := range n.children {
		fmt.Printf("\t\"%s\" -> \"%s\";\n", n.name, child.name)
		child.PrintGraph()
	}

	for _, dep := range n.dependsOn {
		fmt.Printf("\t\"%s\" -> \"%s\" [style=dashed constraint=false color=red];\n", n.name, dep.name)
	}
}

type Field struct {
	*node

	codec api.Codec[any]
}

func NewField(name string, codec api.Codec[any]) *Field {
	return &Field{
		node: &node{
			name:      name,
			redefines: "",
			occurs:    0,
			when:      nil,
			redefined: make(map[string]*node),
			children:  []*node{},
			dependsOn: []*node{},
		},
		codec: codec,
	}
}

type Record struct {
	*node
}

func NewRecord(name string) *Record {
	return &Record{
		node: &node{
			name:      name,
			redefines: "",
			occurs:    0,
			when:      nil,
			redefined: make(map[string]*node),
			children:  []*node{},
			dependsOn: []*node{},
		},
	}
}

func (r *Record) AddField(field *Field) {
	r.addChild(field.node)
}

func (r *Record) AddRecord(record *Record) {
	r.addChild(record.node)
}

func (r *Record) PrintGraph() {
	fmt.Printf("digraph %s {\n", r.name)

	fmt.Printf("\tnode [shape = box fixedsize=true width=3];\n")

	r.node.compileMarshalingPath()
	r.node.PrintGraph()

	fmt.Printf("}\n")
}
