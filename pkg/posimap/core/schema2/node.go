package schema2

import (
	"fmt"
	"slices"

	"github.com/cgi-fr/posimap/pkg/posimap/api"
	"github.com/cgi-fr/posimap/pkg/posimap/core/codec"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/encoding/charmap"
)

type node struct {
	id        string
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
		child.id = n.id + "." + child.name
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

func (n *node) fixMissingFillers() {
	if len(n.dependsOn) == 0 {
		return
	}

	for _, dependent := range n.dependsOn {
		dependent.fixMissingFillers()
	}

	if len(n.dependsOn) == 1 {
		return
	}

	offsets := make([]int, len(n.dependsOn))
	for idx, dependent := range n.dependsOn {
		offsets[idx] = dependent.element.Offset()
	}

	maxOffset := slices.Max(offsets)

	for idx, dependent := range n.dependsOn {
		dependentOffset := offsets[idx]
		if dependentOffset < maxOffset {
			log.Warn().Msgf("Adding missing filler of len %d for %s", maxOffset-dependentOffset, dependent.name)
			dependent.insertFiller(maxOffset - dependentOffset)
		}
	}
}

func (n *node) insertFiller(size int) {
	switch typed := n.element.(type) {
	case *Field:
		log.Error().Msgf("Cannot add filler to field %s", typed.name)
	case *Record:
		filler := &Field{
			node: &node{
				id:        typed.id + ".FILLER",
				name:      "FILLER",
				redefines: "",
				occurs:    0,
				when:      nil,
				element:   nil,
				redefined: make(map[string]*node),
				children:  []*node{},
				dependsOn: []*node{},
			},
			codec: codec.NewString(charmap.ISO8859_1, size, false),
		}
		filler.node.element = filler
		typed.addChild(filler.node)
		typed.setDependsOn(filler.node)
	}
}

func (n *node) printGraph(showDependsOn bool) {
	for _, child := range n.children {
		fmt.Printf("\t\"%s\" [label = \"%s\\n%d\"];\n", n.id, n.name, n.element.Size())
		fmt.Printf("\t\"%s\" [label = \"%s\\n%d\"];\n", child.id, child.name, child.element.Size())
		fmt.Printf("\t\"%s\" -> \"%s\";\n", n.id, child.id)
		child.printGraph(showDependsOn)
	}

	if showDependsOn {
		for _, dep := range n.dependsOn {
			fmt.Printf("\t\"%s\" -> \"%s\" [style=dashed constraint=false color=red label=%d];\n", n.id, dep.id, dep.element.Offset()) //nolint:lll
		}
	}
}

type Element interface {
	Offset() int
	Size() int
	IsCodec() bool
	IsSchema() bool
	Codec() api.Codec[any]
	Schema() *Record
}

type Field struct {
	*node

	codec api.Codec[any]
}

func NewField(name string, codec api.Codec[any], options ...Option) *Field {
	field := &Field{
		node: &node{
			id:        "", // will be set in compileMarshalingPath
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

	if len(offsets) == 0 {
		return f.codec.Size()
	}

	maxOffset := slices.Max(offsets)

	if len(offsets) > 1 {
		for idx, offset := range offsets {
			if offset < maxOffset {
				log.Error().
					Msgf("Unmapped positions (length=%d) between %s and %s", maxOffset-offset, f.name, f.dependsOn[idx].name)
			}
		}
	}

	if f.occurs > 0 {
		return maxOffset + f.Size()
	}

	return maxOffset + f.codec.Size()
}

func (f *Field) Size() int {
	if f.codec == nil {
		return 0
	}

	if f.occurs > 0 {
		return f.codec.Size() * f.occurs
	}

	return f.codec.Size()
}

func (f *Field) IsCodec() bool {
	return true
}

func (f *Field) IsSchema() bool {
	return false
}

func (f *Field) Codec() api.Codec[any] {
	return f.codec
}

func (f *Field) Schema() *Record {
	return nil
}

type Record struct {
	*node
}

func NewRecord(name string, options ...Option) *Record {
	record := &Record{
		node: &node{
			id:        "ROOT", // will be set in compileMarshalingPath
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

	if len(offsets) == 0 {
		return 0
	}

	maxOffset := slices.Max(offsets)

	if len(offsets) > 1 {
		for idx, offset := range offsets {
			if offset < maxOffset {
				log.Error().
					Msgf("Unmapped positions (length=%d) between %s and %s", maxOffset-offset, r.name, r.dependsOn[idx].name)
			}
		}
	}

	if r.occurs > 0 {
		// 1 occurs is already accounted for in the offset
		return maxOffset + r.Size() - r.Size()/r.occurs
	}

	return maxOffset
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

	if r.occurs > 0 {
		return size * r.occurs
	}

	return size
}

func (r *Record) WithField(name string, codec api.Codec[any], options ...Option) *Record {
	field := NewField(name, codec, options...)

	r.AddField(field)

	return r
}

func (r *Record) WithRecord(name string, record *Record, options ...Option) *Record {
	result := NewRecord(name, options...)

	for _, field := range record.children {
		result.addChild(field)
	}

	r.AddRecord(result)

	return r
}

func (r *Record) Elements() []Element {
	fields := make([]Element, 0)

	for _, child := range r.children {
		if child.redefines == "" {
			fields = append(fields, child.element)
		}
	}

	return fields
}

func (r *Record) AddField(field *Field) {
	r.addChild(field.node)
}

func (r *Record) AddRecord(record *Record) {
	r.addChild(record.node)
}

// Validate ensures the record and all its children are consistent.
// It verifies that all field offsets are correct.
// If fillers are missing, it attempts to add them.
// Any unfixable issues will be logged at the error level.
func (r *Record) Validate() {
	r.node.compileMarshalingPath()
	r.node.fixMissingFillers()
	r.Offset() // will trigger error logs far any unfixable issues
}

func (r *Record) PrintGraph(showDependsOn bool) {
	fmt.Printf("digraph \"%s\" {\n", r.id)

	fmt.Printf("\tnode [shape = box fixedsize=true width=3];\n")

	r.node.compileMarshalingPath()
	r.node.fixMissingFillers()
	r.node.printGraph(showDependsOn)

	fmt.Printf("}\n")
}

func (r *Record) IsCodec() bool {
	return false
}

func (r *Record) IsSchema() bool {
	return true
}

func (r *Record) Codec() api.Codec[any] {
	return nil
}

func (r *Record) Schema() *Record {
	return r
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
