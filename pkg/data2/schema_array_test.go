package data2_test

import (
	"fmt"
	"strings"

	"github.com/cgi-fr/posimap/pkg/data2"
	"github.com/cgi-fr/posimap/pkg/deep"
	"github.com/cgi-fr/posimap/pkg/flat"
)

func ExampleSchemaArray() {
	// Create a new SchemaValue with a length of 5
	schemaValue := data2.NewSchemaValue(5, true, nil)

	// Create a new SchemaArray with the schemaValue and occurs set to 3
	schemaArray := data2.NewSchemaArray(schemaValue, 3, nil)

	// Print the rune length
	fmt.Println("Rune Count:", schemaArray.RuneCount())

	// Create a source and buffer
	source := flat.NewSourceVariableWidth(strings.NewReader("Hello, World!!!"), nil)
	buffer := data2.NewBuffer()

	// Read from the source into the buffer
	if err := schemaArray.ReadBuffer(source, &buffer); err != nil {
		panic(err)
	}

	// Create a record from position 0
	record, err := schemaArray.CreateRecord(buffer, 0)
	if err != nil {
		panic(err)
	}

	// Create a new output buffer
	output := &strings.Builder{}

	// Create a sink to write the output
	sink := deep.NewSinkJSONLine(output)

	// Export the record to stdout
	if err := record.Export(record, sink); err != nil {
		panic(err)
	}

	// Don't forget to close the sink
	sink.Close()

	// Print the output
	fmt.Println("Buffer Data:", output.String())
	// Output:
	// Rune Count: 15
	// Buffer Data: ["Hello",", Wor","ld!!!"]
}
