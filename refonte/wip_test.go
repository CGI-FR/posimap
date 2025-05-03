package refonte_test

import (
	"fmt"

	"github.com/cgi-fr/posimap/refonte"
	"golang.org/x/text/encoding/unicode"
)

func Example() {
	// Example usage of the String struct
	data := refonte.NewStaticBuffer([]byte("Héllo, World!   "))
	encoding := unicode.UTF8

	str := refonte.NewString(encoding, 5).
		Then(refonte.NewString(encoding, 2)).
		Then(refonte.NewString(encoding, 5)).
		Then(refonte.NewString(encoding, 1))

	str.Unmarshal(data)

	// Print the value
	fmt.Println(str)

	str.Set("?")
	fmt.Println(str)
	// Output:
	// Héllo/, /World/!
	// Héllo/, /World/?
}
