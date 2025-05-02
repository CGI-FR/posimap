package data2_test

import (
	"fmt"
	"strings"

	"github.com/cgi-fr/posimap/pkg/data2"
	"github.com/cgi-fr/posimap/pkg/deep"
	"github.com/cgi-fr/posimap/pkg/flat"
)

func ExampleSchemaValue() {
	// Create a new SchemaValue with a length of 5
	schemaValue := data2.NewSchemaValue(5)

	// Print the rune count
	fmt.Println("Rune Count:", schemaValue.RuneCount())

	// Create a source and buffer
	source := flat.NewSourceVariableWidth(strings.NewReader("Hello, World!"), nil)
	buffer := data2.NewBuffer()

	// Read from the source into the buffer
	if err := schemaValue.ReadBuffer(source, &buffer); err != nil {
		panic(err)
	}

	// Create a record from position 0
	record, err := schemaValue.CreateRecord(buffer, 0)
	if err != nil {
		panic(err)
	}

	// Create a new output buffer
	output := &strings.Builder{}

	// Create a sink to write the output
	sink := deep.NewSinkJSONLine(output)

	// Export the record to stdout
	if err := record.Export(sink); err != nil {
		panic(err)
	}

	// Don't forget to close the sink
	sink.Close()

	// Print the output
	fmt.Println("Buffer Data:", output.String())
	// Output:
	// Rune Count: 5
	// Buffer Data: "Hello"
}
