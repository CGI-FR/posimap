package schema_test

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/cgi-fr/posimap/refonte/pkg/buffer"
	"github.com/cgi-fr/posimap/refonte/pkg/codec"
	"github.com/cgi-fr/posimap/refonte/pkg/jsonline"
	"github.com/cgi-fr/posimap/refonte/pkg/predicate"
	"github.com/cgi-fr/posimap/refonte/pkg/schema"
	"golang.org/x/text/encoding/charmap"
)

//nolint:lll
func Example() {
	schema := schema.NewSchema().
		WithField("ID", codec.NewString(charmap.ISO8859_1, 16, false), schema.Condition(predicate.Never())).
		WithRecord("PERSON", schema.NewSchema().
			WithField("FIRSTNAME", codec.NewString(charmap.ISO8859_1, 8, true)).
			WithField("LASTNAME", codec.NewString(charmap.ISO8859_1, 8, true)),
			schema.Redefines("ID"),
			schema.Condition(predicate.When(`{{ .ISCOMPANY | ne "1" }}`))).
		WithRecord("COMPANY", schema.NewSchema().
			WithField("NAME", codec.NewString(charmap.ISO8859_1, 16, true)),
			schema.Redefines("ID"),
			schema.Condition(predicate.When(`{{ .ISCOMPANY | ne "0" }}`))).
		WithRecord("ADDRESSES", schema.NewSchema().
			WithField("LINE-1", codec.NewString(charmap.ISO8859_1, 25, true)).
			WithField("LINE-2", codec.NewString(charmap.ISO8859_1, 25, true)),
			schema.Occurs(2)).
		WithField("ISCOMPANY", codec.NewString(charmap.ISO8859_1, 1, true)).
		WithField("TITLES", codec.NewString(charmap.ISO8859_1, 2, true), schema.Occurs(4))

	record, err := schema.Build()
	if err != nil {
		panic(err)
	}

	data := "" +
		"JOHN    DOE     1234 ELM STREET          SPRINGFIELD, IL 62704    56 MAPLE AVENUE          RIVERSIDE, CA 92501      0DRPR    " + "\n" + //nolint:lll
		"ACME COMPANY    1234 ELM STREET          SPRINGFIELD, IL 62704    56 MAPLE AVENUE          RIVERSIDE, CA 92501      1DRPR    " //nolint:lll
	source := strings.NewReader(data)
	buffer := buffer.NewMemoryWithSource(source)

	if err := record.Unmarshal(buffer); err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}

	writer := jsonline.NewWriter(os.Stdout)
	if err := record.Export(writer); err != nil {
		panic(err)
	}

	if err := writer.WriteEOF(); err != nil {
		panic(err)
	}

	// Output:
	// {"PERSON":{"FIRSTNAME":"JOHN","LASTNAME":"DOE"},"ADDRESSES":[{"LINE-1":"1234 ELM STREET","LINE-2":"SPRINGFIELD, IL 62704"},{"LINE-1":"56 MAPLE AVENUE","LINE-2":"RIVERSIDE, CA 92501"}],"ISCOMPANY":"0","TITLES":["DR","PR","",""]}
}
