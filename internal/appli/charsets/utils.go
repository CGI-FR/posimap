package charsets

import (
	"errors"
	"fmt"
)

var (
	ErrUnsupportedCharset       = errors.New("unsupported charset")
	ErrUnsupportedRuneInCharset = errors.New("unsupported rune in charset")
)

func GetByteInCharset(charset string, char rune) (byte, error) {
	charmap, err := Get(charset)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrUnsupportedCharset, charset)
	}

	b, ok := charmap.EncodeRune(char)
	if !ok {
		return 0, fmt.Errorf("%w: %q in %s", ErrUnsupportedRuneInCharset, char, charset)
	}

	return b, nil
}

func GetBytesInCharset(charset string, value string) ([]byte, error) {
	charmap, err := Get(charset)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedCharset, charset)
	}

	bytes := make([]byte, len(value))

	for i, char := range value {
		var ok bool
		if bytes[i], ok = charmap.EncodeRune(char); !ok {
			return nil, fmt.Errorf("%w: %q in %s", ErrUnsupportedRuneInCharset, char, charset)
		}
	}

	return bytes, nil
}
