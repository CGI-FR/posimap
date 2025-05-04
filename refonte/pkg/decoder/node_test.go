package decoder_test

import (
	"fmt"
	"testing"

	"github.com/cgi-fr/posimap/refonte/infra/buffer"
	"github.com/cgi-fr/posimap/refonte/pkg/decoder"
	"golang.org/x/text/encoding/unicode"
)

func Example() {
	data := buffer.NewStatic([]byte("Héllo, World!"))

	firstWord := decoder.NewDecoderString(unicode.UTF8, 5)
	comma := decoder.NewDecoderString(unicode.UTF8, 2)
	secondWord := decoder.NewDecoderString(unicode.UTF8, 5)
	exclamation := decoder.NewDecoderString(unicode.UTF8, 1)

	record := decoder.NewNodeKeyed()
	record.Add("firstWord", decoder.NewNode(firstWord))
	record.Add("comma", decoder.NewNode(comma))
	record.Add("secondWord", decoder.NewNode(secondWord))
	record.Add("exclamation", decoder.NewNode(exclamation))
	record.Unmarshal(data)

	fmt.Println("{")

	for key, value := range record.KeyValuePairs() {
		fmt.Printf(`    "%s": "%v"`, key, value)
		fmt.Println()
	}

	fmt.Println("}")

	// Output:
	// {
	//     "firstWord": "Héllo"
	//     "comma": ", "
	//     "secondWord": "World"
	//     "exclamation": "!"
	// }
}

func TestRedefine(t *testing.T) {
	data := buffer.NewStatic([]byte("AABB1122"))

	group1 := decoder.NewDecoderString(unicode.UTF8, 2)
	group2 := decoder.NewDecoderString(unicode.UTF8, 2)
	group3 := decoder.NewDecoderString(unicode.UTF8, 2)
	group4 := decoder.NewDecoderString(unicode.UTF8, 2)
	groups := decoder.NewNodeKeyed()
	groups.Add("group1", decoder.NewNode(group1)) // group1<->groups
	groups.Add("group2", decoder.NewNode(group2)) // group1<->group2<->groups
	groups.Add("group3", decoder.NewNode(group3)) // group1<->group2<->group3<->groups
	groups.Add("group4", decoder.NewNode(group4)) // group1<->group2<->group3<->group4<->groups

	letters := decoder.NewDecoderString(unicode.UTF8, 4)
	numbers := decoder.NewDecoderString(unicode.UTF8, 4)
	chars := decoder.NewNodeKeyed()
	chars.Add("letters", decoder.NewNode(letters))
	chars.Add("numbers", decoder.NewNode(numbers))

	record := decoder.NewNodeKeyed()
	record.Add("groups", groups) // groups<->record
	// record.Redefine("chars", "groups", chars)
	record.Unmarshal(data)

	fmt.Println("{")

	for key, value := range groups.KeyValuePairs() {
		fmt.Printf(`    "%s": "%v"`, key, value)
		fmt.Println()
	}

	for key, value := range chars.KeyValuePairs() {
		fmt.Printf(`    "%s": "%v"`, key, value)
		fmt.Println()
	}

	fmt.Println("}")
}
