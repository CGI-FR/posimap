package data

const defaultBufferSize = 4 * 1024

type Buffer struct {
	data []rune
}

func NewBuffer() *Buffer {
	return &Buffer{data: make([]rune, 0, defaultBufferSize)}
}

func NewBufferFrom(data string) *Buffer {
	return &Buffer{data: []rune(data)}
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
	if start >= len(b.data) || start < 0 {
		return nil
	}

	b.Grow(start + length)

	done := length

	for idx, r := range value {
		b.data[start+idx] = r

		if done--; done == 0 {
			break
		}
	}

	for idx := range done {
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
