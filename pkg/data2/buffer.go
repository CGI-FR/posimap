package data2

import (
	"strings"
)

const defaultBufferSize = 4 * 1024

const (
	visibleBlankRunes = " \u00A0\u1680\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200A\u202F\u205F\u3000" //nolint:lll
	controlBlankRunes = "\t\n\v\f\r\u0085\u2028\u2029"
	blankRunes        = visibleBlankRunes + controlBlankRunes
)

type Buffer []rune

func NewBuffer() Buffer {
	return make([]rune, 0, defaultBufferSize)
}

func (b Buffer) String(start, length int, trim bool) string {
	if trim {
		return strings.TrimRight(b.Read(start, length), blankRunes)
	}

	return b.Read(start, length)
}

func (b Buffer) Read(start, length int) string {
	if start >= len(b) || start < 0 {
		return string(b[0:0])
	}

	if length == 0 {
		return string(b[start:])
	}

	if start+length > len(b) {
		return string(b[start:])
	}

	return string(b[start : start+length])
}

func (b Buffer) Length() int {
	return len(b)
}
