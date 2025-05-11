package schema2_test

import (
	"github.com/cgi-fr/posimap/pkg/posimap/core/codec"
	"github.com/cgi-fr/posimap/pkg/posimap/core/schema2"
	"golang.org/x/text/encoding/charmap"
)

func Example() {
	root := schema2.NewRecord("ROOT")

	rootID := schema2.NewField("ID", codec.NewString(charmap.ISO8859_1, 20, false))
	personFirstname := schema2.NewField("FIRSTNAME", codec.NewString(charmap.ISO8859_1, 8, true))
	personLastname := schema2.NewField("LASTNAME", codec.NewString(charmap.ISO8859_1, 8, true))
	companyName := schema2.NewField("NAME", codec.NewString(charmap.ISO8859_1, 16, true))
	companyType := schema2.NewField("TYPE", codec.NewString(charmap.ISO8859_1, 4, true))
	addressesLine1 := schema2.NewField("LINE-1", codec.NewString(charmap.ISO8859_1, 25, true))
	addressesLine2 := schema2.NewField("LINE-2", codec.NewString(charmap.ISO8859_1, 25, true))
	rootIscompany := schema2.NewField("ISCOMPANY", codec.NewString(charmap.ISO8859_1, 1, true))
	rootTitles := schema2.NewField("TITLES", codec.NewString(charmap.ISO8859_1, 2, true), schema2.Occurs(4))
	rootFiller := schema2.NewField("FILLER", codec.NewString(charmap.ISO8859_1, 21, true))

	person := schema2.NewRecord("PERSON", schema2.Redefines("ID"))
	person.AddField(personFirstname)
	person.AddField(personLastname)

	company := schema2.NewRecord("COMPANY", schema2.Redefines("ID"))
	company.AddField(companyName)
	company.AddField(companyType)

	addresses := schema2.NewRecord("ADDRESSES", schema2.Occurs(2))
	addresses.AddField(addressesLine1)
	addresses.AddField(addressesLine2)

	root.AddField(rootID)
	root.AddRecord(person)
	root.AddRecord(company)
	root.AddRecord(addresses)
	root.AddField(rootIscompany)
	root.AddField(rootTitles)
	root.AddField(rootFiller)

	root.PrintGraph(true)

	// Output:
	// digraph "ROOT" {
	// 	node [shape = box fixedsize=true width=3];
	// 	"ROOT" [label = "ROOT\n150"];
	// 	"ROOT.ID" [label = "ID\n20"];
	// 	"ROOT" -> "ROOT.ID";
	// 	"ROOT" [label = "ROOT\n150"];
	// 	"ROOT.PERSON" [label = "PERSON\n20"];
	// 	"ROOT" -> "ROOT.PERSON";
	// 	"ROOT.PERSON" [label = "PERSON\n20"];
	// 	"ROOT.PERSON.FIRSTNAME" [label = "FIRSTNAME\n8"];
	// 	"ROOT.PERSON" -> "ROOT.PERSON.FIRSTNAME";
	// 	"ROOT.PERSON" [label = "PERSON\n20"];
	// 	"ROOT.PERSON.LASTNAME" [label = "LASTNAME\n8"];
	// 	"ROOT.PERSON" -> "ROOT.PERSON.LASTNAME";
	// 	"ROOT.PERSON.LASTNAME" -> "ROOT.PERSON.FIRSTNAME" [style=dashed constraint=false color=red label=8];
	// 	"ROOT.PERSON" [label = "PERSON\n20"];
	// 	"ROOT.PERSON.FILLER" [label = "FILLER\n4"];
	// 	"ROOT.PERSON" -> "ROOT.PERSON.FILLER";
	// 	"ROOT.PERSON.FILLER" -> "ROOT.PERSON.LASTNAME" [style=dashed constraint=false color=red label=16];
	// 	"ROOT.PERSON" -> "ROOT.PERSON.FILLER" [style=dashed constraint=false color=red label=20];
	// 	"ROOT" [label = "ROOT\n150"];
	// 	"ROOT.COMPANY" [label = "COMPANY\n20"];
	// 	"ROOT" -> "ROOT.COMPANY";
	// 	"ROOT.COMPANY" [label = "COMPANY\n20"];
	// 	"ROOT.COMPANY.NAME" [label = "NAME\n16"];
	// 	"ROOT.COMPANY" -> "ROOT.COMPANY.NAME";
	// 	"ROOT.COMPANY" [label = "COMPANY\n20"];
	// 	"ROOT.COMPANY.TYPE" [label = "TYPE\n4"];
	// 	"ROOT.COMPANY" -> "ROOT.COMPANY.TYPE";
	// 	"ROOT.COMPANY.TYPE" -> "ROOT.COMPANY.NAME" [style=dashed constraint=false color=red label=16];
	// 	"ROOT.COMPANY" -> "ROOT.COMPANY.TYPE" [style=dashed constraint=false color=red label=20];
	// 	"ROOT" [label = "ROOT\n150"];
	// 	"ROOT.ADDRESSES" [label = "ADDRESSES\n100"];
	// 	"ROOT" -> "ROOT.ADDRESSES";
	// 	"ROOT.ADDRESSES" [label = "ADDRESSES\n100"];
	// 	"ROOT.ADDRESSES.LINE-1" [label = "LINE-1\n25"];
	// 	"ROOT.ADDRESSES" -> "ROOT.ADDRESSES.LINE-1";
	// 	"ROOT.ADDRESSES.LINE-1" -> "ROOT.ID" [style=dashed constraint=false color=red label=20];
	// 	"ROOT.ADDRESSES.LINE-1" -> "ROOT.PERSON" [style=dashed constraint=false color=red label=20];
	// 	"ROOT.ADDRESSES.LINE-1" -> "ROOT.COMPANY" [style=dashed constraint=false color=red label=20];
	// 	"ROOT.ADDRESSES" [label = "ADDRESSES\n100"];
	// 	"ROOT.ADDRESSES.LINE-2" [label = "LINE-2\n25"];
	// 	"ROOT.ADDRESSES" -> "ROOT.ADDRESSES.LINE-2";
	// 	"ROOT.ADDRESSES.LINE-2" -> "ROOT.ADDRESSES.LINE-1" [style=dashed constraint=false color=red label=45];
	// 	"ROOT.ADDRESSES" -> "ROOT.ADDRESSES.LINE-2" [style=dashed constraint=false color=red label=70];
	// 	"ROOT" [label = "ROOT\n150"];
	// 	"ROOT.ISCOMPANY" [label = "ISCOMPANY\n1"];
	// 	"ROOT" -> "ROOT.ISCOMPANY";
	// 	"ROOT.ISCOMPANY" -> "ROOT.ADDRESSES" [style=dashed constraint=false color=red label=120];
	// 	"ROOT" [label = "ROOT\n150"];
	// 	"ROOT.TITLES" [label = "TITLES\n8"];
	// 	"ROOT" -> "ROOT.TITLES";
	// 	"ROOT.TITLES" -> "ROOT.ISCOMPANY" [style=dashed constraint=false color=red label=121];
	// 	"ROOT" [label = "ROOT\n150"];
	// 	"ROOT.FILLER" [label = "FILLER\n21"];
	// 	"ROOT" -> "ROOT.FILLER";
	// 	"ROOT.FILLER" -> "ROOT.TITLES" [style=dashed constraint=false color=red label=129];
	// 	"ROOT" -> "ROOT.FILLER" [style=dashed constraint=false color=red label=150];
	// }
}
