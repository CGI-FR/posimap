package data_test

import (
	"fmt"

	"github.com/cgi-fr/posimap/pkg/data"
)

func Example() {
	// Example usage of the Buffer type
	buffer := data.NewBuffer()
	buffer.Grow(11)

	if err := buffer.Write(0, 5, "Hello"); err != nil {
		panic(err)
	}

	if err := buffer.Write(6, 5, "World"); err != nil {
		panic(err)
	}

	fmt.Print(buffer.String()) // Output: Hello World
}
