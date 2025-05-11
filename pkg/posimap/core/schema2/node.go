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

	element Element

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
		fmt.Printf("\t\"%s\" [label = \"%s\\n%d\"];\n", n.name, n.name, n.element.Size())
		fmt.Printf("\t\"%s\" [label = \"%s\\n%d\"];\n", child.name, child.name, child.element.Size())
		fmt.Printf("\t\"%s\" -> \"%s\";\n", n.name, child.name)
		child.PrintGraph()
	}

	for _, dep := range n.dependsOn {
		fmt.Printf("\t\"%s\" -> \"%s\" [style=dashed constraint=false color=red label=%d];\n", n.name, dep.name, dep.element.Offset()) //nolint:lll
	}
}

type Element interface {
	Offset() int
	Size() int
}

type Field struct {
	*node

	codec api.Codec[any]
}

func NewField(name string, codec api.Codec[any], options ...Option) *Field {
	field := &Field{
		node: &node{
			name:      name,
			redefines: "",
			occurs:    0,
			when:      nil,
			element:   nil,
			redefined: make(map[string]*node),
			children:  []*node{},
			dependsOn: []*node{},
		},
		codec: codec,
	}

	for _, option := range options {
		option(field.node)
	}

	field.node.element = field

	return field
}

func (f *Field) Offset() int {
	offsets := make([]int, len(f.dependsOn))
	for idx, dependent := range f.dependsOn {
		offsets[idx] = dependent.element.Offset()
	}

	if len(offsets) > 1 {
		checkOffset := offsets[0]
		for _, offset := range offsets {
			if offset != checkOffset {
				panic("Offsets are not equal")
			}
		}
	}

	if len(offsets) == 0 {
		return f.codec.Size()
	}

	return offsets[0] + f.codec.Size()
}

func (f *Field) Size() int {
	if f.codec == nil {
		return 0
	}

	return f.codec.Size()
}

type Record struct {
	*node
}

func NewRecord(name string, options ...Option) *Record {
	record := &Record{
		node: &node{
			name:      name,
			redefines: "",
			occurs:    0,
			when:      nil,
			element:   nil,
			redefined: make(map[string]*node),
			children:  []*node{},
			dependsOn: []*node{},
		},
	}

	for _, option := range options {
		option(record.node)
	}

	record.node.element = record

	return record
}

func (r *Record) Offset() int {
	offsets := make([]int, len(r.dependsOn))
	for idx, dependent := range r.dependsOn {
		offsets[idx] = dependent.element.Offset()
	}

	if len(offsets) > 1 {
		checkOffset := offsets[0]
		for _, offset := range offsets {
			if offset != checkOffset {
				panic("Offsets are not equal")
			}
		}
	}

	if len(offsets) == 0 {
		return 0
	}

	return offsets[0]
}

func (r *Record) Size() int {
	if len(r.children) == 0 {
		return 0
	}

	size := 0

	for _, child := range r.children {
		if child.redefines == "" {
			size += child.element.Size()
		}
	}

	return size
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

type Option func(*node)

func Occurs(occurs int) Option {
	return func(n *node) {
		n.occurs = occurs
	}
}

func Condition(when api.Predicate) Option {
	return func(n *node) {
		n.when = when
	}
}

func Redefines(redefines string) Option {
	return func(n *node) {
		n.redefines = redefines
	}
}
