package buffer

import (
	"fmt"
	"io"
)

const (
	growthFactor      = 2
	defaultBufferSize = 4 * 1024
)

type Buffer struct {
	buffer []byte
	source io.Reader
}

func NewBuffer() *Buffer {
	return &Buffer{
		buffer: make([]byte, 0, defaultBufferSize),
		source: nil,
	}
}

func NewBufferReader(source io.Reader) *Buffer {
	return &Buffer{
		buffer: make([]byte, 0, defaultBufferSize),
		source: source,
	}
}

func (m *Buffer) Slice(offset, length int) ([]byte, error) {
	if err := m.Required(offset + length); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return m.buffer[offset : offset+length], nil
}

func (m *Buffer) Write(offset int, data []byte) error {
	if err := m.Required(offset + len(data)); err != nil {
		return fmt.Errorf("%w", err)
	}

	copy(m.buffer[offset:], data)

	return nil
}

func (m *Buffer) Reset() error {
	size := len(m.buffer)
	m.buffer = m.buffer[:0]

	// immediately refill the buffer
	return m.Required(size)
}

func (m *Buffer) Bytes() []byte {
	return m.buffer
}

func (m *Buffer) Required(size int) error {
	cursize := len(m.buffer)
	if cursize >= size {
		return nil
	}

	if !m.tryReslice(size) {
		capacity := max(size, growthFactor*cap(m.buffer))
		newBuffer := append([]byte(nil), make([]byte, capacity)...)
		copy(newBuffer, m.buffer)
		m.buffer = m.buffer[:size]
	}

	if m.source != nil {
		_, err := m.source.Read(m.buffer[cursize:])
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

// tryReslice is an inlineable version for the fast-case where the
// internal buffer only needs to be resliced.
// It returns whether it succeeded.
func (m *Buffer) tryReslice(size int) bool {
	if size <= cap(m.buffer) {
		m.buffer = m.buffer[:size]

		return true
	}

	return false
}
