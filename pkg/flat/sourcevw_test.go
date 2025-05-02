package flat_test

import (
	"bytes"
	"fmt"
	"testing"
	"unicode/utf8"

	"github.com/cgi-fr/posimap/pkg/flat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
)

func ExampleSourceVariableWidth() {
	data := []byte("Hello, café!")
	reader := bytes.NewReader(data)

	source := flat.NewSourceVariableWidth(reader, unicode.UTF8)

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

func TestReadRunes_CommonEncodings(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		input    []byte
		encoding encoding.Encoding
		expected []rune
	}

	tests := []testCase{
		{
			name:     "UTF-8",
			input:    []byte{0x43, 0x61, 0x66, 0xC3, 0xA9},
			encoding: nil, // UTF-8 natif
			expected: []rune{'C', 'a', 'f', 'é'},
		},
		{
			name:     "UTF-16LE",
			input:    []byte{0x43, 0x00, 0x61, 0x00, 0x66, 0x00, 0xE9, 0x00},
			encoding: unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM),
			expected: []rune{'C', 'a', 'f', 'é'},
		},
		{
			name:     "UTF-16BE",
			input:    []byte{0x00, 0x43, 0x00, 0x61, 0x00, 0x66, 0x00, 0xE9},
			encoding: unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM),
			expected: []rune{'C', 'a', 'f', 'é'},
		},
		{
			name:     "ISO-8859-1",
			input:    []byte{0x43, 0x61, 0x66, 0xE9},
			encoding: charmap.ISO8859_1,
			expected: []rune{'C', 'a', 'f', 'é'},
		},
		{
			name:     "Windows-1252",
			input:    []byte{0x43, 0x80, 0x66, 0xE9},
			encoding: charmap.Windows1252,
			expected: []rune{'C', '€', 'f', 'é'},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			reader := bytes.NewReader(test.input)
			source := flat.NewSourceVariableWidth(reader, test.encoding)

			runes, err := source.ReadRunes(len(test.expected))
			require.NoError(t, err)

			assert.Equal(t, test.expected, runes)

			// Optionnel : vérifier que chaque rune est UTF-8 valide
			for _, r := range runes {
				buf := make([]byte, utf8.RuneLen(r))
				utf8.EncodeRune(buf, r)
				require.True(t, utf8.Valid(buf))
			}
		})
	}
}
