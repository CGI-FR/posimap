//nolint:lll
package posimap_test

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/cgi-fr/posimap/internal/infra/jsonline"
	"github.com/cgi-fr/posimap/pkg/posimap/core/buffer"
	"github.com/cgi-fr/posimap/pkg/posimap/core/codec"
	"github.com/cgi-fr/posimap/pkg/posimap/core/record"
	"golang.org/x/text/encoding/charmap"
)

func ExampleRecord_Export() {
	data := strings.NewReader("JOHN DOE                           1234 ELM STREET               SPRINGFIELD, IL 62704         ")
	buffer := buffer.NewBufferReader(data)

	rec := record.NewObject()
	rec.Add("NAME", record.NewValue(0, codec.NewString(charmap.ISO8859_1, 35, true)), nil)
	rec.Add("ADDRESS-LINE1", record.NewValue(35, codec.NewString(charmap.ISO8859_1, 30, true)), nil)
	rec.Add("ADDRESS-LINE2", record.NewValue(65, codec.NewString(charmap.ISO8859_1, 30, true)), nil)

	if err := rec.Unmarshal(buffer); err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}

	writer := jsonline.NewWriter(os.Stdout)
	if err := rec.Export(writer); err != nil {
		panic(err)
	}

	if err := writer.Close(); err != nil {
		panic(err)
	}

	// Output:
	// {"NAME":"JOHN DOE","ADDRESS-LINE1":"1234 ELM STREET","ADDRESS-LINE2":"SPRINGFIELD, IL 62704"}
}
