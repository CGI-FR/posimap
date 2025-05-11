package schema2_test

import (
	"github.com/cgi-fr/posimap/pkg/posimap/core/codec"
	"github.com/cgi-fr/posimap/pkg/posimap/core/schema2"
	"golang.org/x/text/encoding/charmap"
)

func Example() {
	root := schema2.NewRecord("ROOT")

	rootID := schema2.NewField("ID", codec.NewString(charmap.ISO8859_1, 16, false))
	personFirstname := schema2.NewField("FIRSTNAME", codec.NewString(charmap.ISO8859_1, 8, true))
	personLastname := schema2.NewField("LASTNAME", codec.NewString(charmap.ISO8859_1, 8, true))
	companyName := schema2.NewField("NAME", codec.NewString(charmap.ISO8859_1, 16, true))
	addressesLine1 := schema2.NewField("LINE-1", codec.NewString(charmap.ISO8859_1, 25, true))
	addressesLine2 := schema2.NewField("LINE-2", codec.NewString(charmap.ISO8859_1, 25, true))
	rootIscompany := schema2.NewField("ISCOMPANY", codec.NewString(charmap.ISO8859_1, 1, true))
	rootTitles := schema2.NewField("TITLES", codec.NewString(charmap.ISO8859_1, 2, true))

	person := schema2.NewRecord("PERSON", schema2.Redefines("ID"))
	person.AddField(personFirstname)
	person.AddField(personLastname)

	company := schema2.NewRecord("COMPANY", schema2.Redefines("ID"))
	company.AddField(companyName)

	addresses := schema2.NewRecord("ADDRESSES")
	addresses.AddField(addressesLine1)
	addresses.AddField(addressesLine2)

	root.AddField(rootID)
	root.AddRecord(person)
	root.AddRecord(company)
	root.AddRecord(addresses)
	root.AddField(rootIscompany)
	root.AddField(rootTitles)

	root.PrintGraph()

	// Output:
}
