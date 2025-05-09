package schema2_test

import (
	"github.com/cgi-fr/posimap/pkg/posimap/core/codec"
	"github.com/cgi-fr/posimap/pkg/posimap/core/predicate"
	"github.com/cgi-fr/posimap/pkg/posimap/core/schema2"
	"golang.org/x/text/encoding/charmap"
)

func Example() {
	schema := schema2.NewSchema().
		WithField("ID", codec.NewString(charmap.ISO8859_1, 16, false), schema2.Condition(predicate.Never())).
		WithRecord("PERSON", schema2.NewSchema().
			WithField("FIRSTNAME", codec.NewString(charmap.ISO8859_1, 8, true)).
			WithField("LASTNAME", codec.NewString(charmap.ISO8859_1, 8, true)),
			schema2.Redefines("ID"),
			schema2.Condition(predicate.When(`{{ .ISCOMPANY | ne "1" }}`))).
		WithRecord("COMPANY", schema2.NewSchema().
			WithField("NAME", codec.NewString(charmap.ISO8859_1, 16, true)),
			schema2.Redefines("ID"),
			schema2.Condition(predicate.When(`{{ .ISCOMPANY | ne "0" }}`))).
		WithRecord("ADDRESSES", schema2.NewSchema().
			WithField("LINE-1", codec.NewString(charmap.ISO8859_1, 25, true)).
			WithField("LINE-2", codec.NewString(charmap.ISO8859_1, 25, true)),
			schema2.Occurs(2)).
		WithField("ISCOMPANY", codec.NewString(charmap.ISO8859_1, 1, true)).
		WithField("TITLES", codec.NewString(charmap.ISO8859_1, 2, true), schema2.Occurs(4))

	schema2.CompileDependsOn(schema)

	schema.Print()

	// for idx, tip := range schema.FindTips() {
	// 	fmt.Print(idx)
	// 	tip.Print()
	// }

	// Output:
	// root -> child1
}
