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
	rootTitles := schema2.NewField("TITLES", codec.NewString(charmap.ISO8859_1, 2, true))

	person := schema2.NewRecord("PERSON", schema2.Redefines("ID"))
	person.AddField(personFirstname)
	person.AddField(personLastname)

	company := schema2.NewRecord("COMPANY", schema2.Redefines("ID"))
	company.AddField(companyName)
	company.AddField(companyType)

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
	// digraph ROOT {
	// 	node [shape = box fixedsize=true width=3];
	// 	"ROOT" [label = "ROOT\n73"];
	// 	"ID" [label = "ID\n20"];
	// 	"ROOT" -> "ID";
	// 	"ROOT" [label = "ROOT\n73"];
	// 	"PERSON" [label = "PERSON\n20"];
	// 	"ROOT" -> "PERSON";
	// 	"PERSON" [label = "PERSON\n20"];
	// 	"FIRSTNAME" [label = "FIRSTNAME\n8"];
	// 	"PERSON" -> "FIRSTNAME";
	// 	"PERSON" [label = "PERSON\n20"];
	// 	"LASTNAME" [label = "LASTNAME\n8"];
	// 	"PERSON" -> "LASTNAME";
	// 	"LASTNAME" -> "FIRSTNAME" [style=dashed constraint=false color=red label=8];
	// 	"PERSON" [label = "PERSON\n20"];
	// 	"FILLER" [label = "FILLER\n4"];
	// 	"PERSON" -> "FILLER";
	// 	"FILLER" -> "LASTNAME" [style=dashed constraint=false color=red label=16];
	// 	"PERSON" -> "FILLER" [style=dashed constraint=false color=red label=20];
	// 	"ROOT" [label = "ROOT\n73"];
	// 	"COMPANY" [label = "COMPANY\n20"];
	// 	"ROOT" -> "COMPANY";
	// 	"COMPANY" [label = "COMPANY\n20"];
	// 	"NAME" [label = "NAME\n16"];
	// 	"COMPANY" -> "NAME";
	// 	"COMPANY" [label = "COMPANY\n20"];
	// 	"TYPE" [label = "TYPE\n4"];
	// 	"COMPANY" -> "TYPE";
	// 	"TYPE" -> "NAME" [style=dashed constraint=false color=red label=16];
	// 	"COMPANY" -> "TYPE" [style=dashed constraint=false color=red label=20];
	// 	"ROOT" [label = "ROOT\n73"];
	// 	"ADDRESSES" [label = "ADDRESSES\n50"];
	// 	"ROOT" -> "ADDRESSES";
	// 	"ADDRESSES" [label = "ADDRESSES\n50"];
	// 	"LINE-1" [label = "LINE-1\n25"];
	// 	"ADDRESSES" -> "LINE-1";
	// 	"LINE-1" -> "ID" [style=dashed constraint=false color=red label=20];
	// 	"LINE-1" -> "PERSON" [style=dashed constraint=false color=red label=20];
	// 	"LINE-1" -> "COMPANY" [style=dashed constraint=false color=red label=20];
	// 	"ADDRESSES" [label = "ADDRESSES\n50"];
	// 	"LINE-2" [label = "LINE-2\n25"];
	// 	"ADDRESSES" -> "LINE-2";
	// 	"LINE-2" -> "LINE-1" [style=dashed constraint=false color=red label=45];
	// 	"ADDRESSES" -> "LINE-2" [style=dashed constraint=false color=red label=70];
	// 	"ROOT" [label = "ROOT\n73"];
	// 	"ISCOMPANY" [label = "ISCOMPANY\n1"];
	// 	"ROOT" -> "ISCOMPANY";
	// 	"ISCOMPANY" -> "ADDRESSES" [style=dashed constraint=false color=red label=70];
	// 	"ROOT" [label = "ROOT\n73"];
	// 	"TITLES" [label = "TITLES\n2"];
	// 	"ROOT" -> "TITLES";
	// 	"TITLES" -> "ISCOMPANY" [style=dashed constraint=false color=red label=71];
	// 	"ROOT" -> "TITLES" [style=dashed constraint=false color=red label=73];
	// }
}
