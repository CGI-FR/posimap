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
	record.Add("firstWord", decoder.NewNode(firstWord))
	record.Add("comma", decoder.NewNode(comma))
	record.Add("secondWord", decoder.NewNode(secondWord))
	record.Add("exclamation", decoder.NewNode(exclamation))
	last, _ := record.ValueForKey("exclamation")
	last.(decoder.Node).Unmarshal(data)

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

// func TestRedefine(t *testing.T) {
// 	data := buffer.NewStatic([]byte("20231001"))

// 	date := decoder.NewDecoderString(unicode.UTF8, 8)
// 	year := decoder.NewDecoderString(unicode.UTF8, 4)
// 	month := decoder.NewDecoderString(unicode.UTF8, 2)
// 	day := decoder.NewDecoderString(unicode.UTF8, 2)

// 	recordDate := decoder.NewNodeKeyed()
// 	recordDate.Add("year", year)
// 	recordDate.Add("month", month)
// 	recordDate.Add("day", day)

// 	record := decoder.NewNodeKeyed()
// 	record.Add("fulldate", date)
// 	record.Redefine("date", "fulldate", recordDate)
// 	record.Redefined("secondWord", "comma", secondWord)
// 	record.Add("exclamation", exclamation)
// 	record.Unmarshal(data)

// 	fmt.Println("{")

// 	for key, value := range record.KeyValuePairs() {
// 		fmt.Printf(`    "%s": "%v"`, key, value)
// 		fmt.Println()
// 	}

// 	fmt.Println("}")

// 	// Output:
// 	// {
// 	//     "firstWord": "Héllo"
// 	//     "comma": ", "
// 	//     "secondWord": "World"
// 	//     "exclamation": "!"
// 	// }
// }
