package refonte_test

import (
	"fmt"

	"github.com/cgi-fr/posimap/refonte"
	"golang.org/x/text/encoding/unicode"
)

func Example() {
	// Example usage of the String struct
	data := refonte.Buffer("Héllo, World!   ")
	encoding := unicode.UTF8

	str := refonte.NewString(encoding, 5).
		Then(refonte.NewString(encoding, 2)).
		Then(refonte.NewString(encoding, 5)).
		Then(refonte.NewString(encoding, 1))

	str.Unmarshal(data)

	// Print the value
	fmt.Print(str)
	// Output: Héllo/, /World/!
}
