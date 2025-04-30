package data

import (
	"strings"
)

const defaultBufferSize = 4 * 1024

const (
	VisibleBlankRunes = " \u00A0\u1680\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200A\u202F\u205F\u3000" //nolint:lll
	ControlBlankRunes = "\t\n\v\f\r\u0085\u2028\u2029"
	BlankRunes        = VisibleBlankRunes + ControlBlankRunes
)

type Buffer struct {
	data []rune
}

func NewBuffer() *Buffer {
	return &Buffer{data: make([]rune, 0, defaultBufferSize)}
}

func NewBufferFrom(data string) *Buffer {
	return &Buffer{data: []rune(data)}
}

func (b *Buffer) ReadTrimmed(start, length int, cutset string) string {
	return strings.TrimRight(b.Read(start, length), cutset)
}

func (b *Buffer) Read(start, length int) string {
	if start >= len(b.data) || start < 0 {
		return string(b.data[0:0])
	}

	if length == 0 {
		return string(b.data[start:])
	}

	if start+length > len(b.data) {
		return string(b.data[start:])
	}

	return string(b.data[start : start+length])
}

func (b *Buffer) Write(start, length int, value string) error {
	if start < 0 {
		return nil
	}

	b.Grow(start + length)

	leftover := length
	done := 0

	for idx, r := range value {
		b.data[start+idx] = r
		done++

		if leftover--; leftover == 0 {
			break
		}
	}

	for idx := range leftover {
		b.data[start+done+idx] = ' '
	}

	return nil
}

func (b *Buffer) String() string {
	return string(b.data)
}

func (b *Buffer) Grow(length int) {
	for len(b.data) < length {
		b.data = append(b.data, ' ')
	}
}
