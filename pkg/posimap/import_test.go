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

func ExampleRecord_Import() {
	buffer := buffer.NewBuffer()

	rec := record.NewObject()
	rec.Add("NAME", record.NewValue(0, codec.NewString(charmap.ISO8859_1, 35, true)), nil)
	rec.Add("ADDRESS-LINE1", record.NewValue(35, codec.NewString(charmap.ISO8859_1, 30, true)), nil)
	rec.Add("ADDRESS-LINE2", record.NewValue(65, codec.NewString(charmap.ISO8859_1, 30, true)), nil)

	data := strings.NewReader(`{"NAME":"JOHN DOE","ADDRESS-LINE1":"1234 ELM STREET","ADDRESS-LINE2":"SPRINGFIELD, IL 62704"}`) //nolint:lll
	reader := jsonline.NewReader(data)

	if err := rec.Import(reader); err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}

	if err := rec.Marshal(buffer); err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}

	os.Stdout.Write(buffer.Bytes())

	// Output:
	// JOHN DOE                           1234 ELM STREET               SPRINGFIELD, IL 62704
}
