package decoder_test

import (
	"fmt"

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
	record.Add("firstWord", firstWord)
	record.Add("comma", comma)
	record.Add("secondWord", secondWord)
	record.Add("exclamation", exclamation)
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
