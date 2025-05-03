package decoder_test

import (
	"fmt"

	"github.com/cgi-fr/posimap/refonte/infra/buffer"
	"github.com/cgi-fr/posimap/refonte/pkg/decoder"
	"golang.org/x/text/encoding/unicode"
)

func Example() {
	data := buffer.NewStatic([]byte("Héllo, World!   "))

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

	fmt.Println(record)
	// Output:
	// {"firstWord":[72 233 108 108 111],"comma":[44 32],"secondWord":[87 111 114 108 100],"exclamation":[33]}
}
