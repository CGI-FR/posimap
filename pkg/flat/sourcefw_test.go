package flat_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/cgi-fr/posimap/pkg/flat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/encoding/charmap"
)

func ExampleSourceFixedWidth() {
	data := []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x2C, 0x20, 0x63, 0x61, 0x66, 0xE9, 0x21} // "Hello, café!" in ISO-8859-15
	reader := bytes.NewReader(data)

	source := flat.NewSourceFixedWidth(reader, charmap.ISO8859_15)

	runes, err := source.ReadRunes(5)
	if err != nil {
		panic(err)
	}

	bytes, err := source.ReadBytes(2)
	if err != nil {
		panic(err)
	}

	remaining, err := source.ReadRunes(5)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", runes)
	fmt.Printf("%v\n", bytes)
	fmt.Printf("%v\n", remaining)
	// Output:
	// [72 101 108 108 111]
	// [44 32]
	// [99 97 102 233 33]
}

func TestReadRunes_CommonCharmaps(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		charmap  *charmap.Charmap
		input    []byte
		expected []rune
	}

	tests := []testCase{
		{
			name:     "ISO-8859-1",
			charmap:  charmap.ISO8859_1,
			input:    []byte{0x43, 0x61, 0x66, 0xE9},
			expected: []rune{'C', 'a', 'f', 'é'},
		},
		{
			name:     "ISO-8859-15",
			charmap:  charmap.ISO8859_15,
			input:    []byte{0x43, 0xA4, 0x66, 0xE9},
			expected: []rune{'C', '€', 'f', 'é'},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			reader := bytes.NewReader(test.input)
			source := flat.NewSourceFixedWidth(reader, test.charmap)

			runes, err := source.ReadRunes(len(test.expected))
			require.NoError(t, err)

			assert.Equal(t, test.expected, runes)
		})
	}
}
