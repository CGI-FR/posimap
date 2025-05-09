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
	target io.Writer
}

func NewBufferWriter(target io.Writer) *Buffer {
	return &Buffer{
		buffer: make([]byte, 0, defaultBufferSize),
		source: nil,
		target: target,
	}
}

func NewBufferReader(source io.Reader) *Buffer {
	return &Buffer{
		buffer: make([]byte, 0, defaultBufferSize),
		source: source,
		target: nil,
	}
}

func (m *Buffer) Slice(offset, length int) ([]byte, error) {
	if err := m.growTo(offset + length); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return m.buffer[offset : offset+length], nil
}

func (m *Buffer) Write(offset int, data []byte) error {
	if err := m.growTo(offset + len(data)); err != nil {
		return fmt.Errorf("%w", err)
	}

	copy(m.buffer[offset:], data)

	return nil
}

func (m *Buffer) Reset(size int) error {
	if m.target != nil {
		if _, err := m.target.Write(m.buffer); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	if size == 0 {
		size = len(m.buffer)
	}

	m.buffer = m.buffer[:0]

	// immediately refill the buffer
	return m.growTo(size)
}

func (m *Buffer) Fill(byteValue byte) {
	for i := range m.buffer {
		m.buffer[i] = byteValue
	}
}

func (m *Buffer) Bytes() []byte {
	return m.buffer
}

func (m *Buffer) growTo(size int) error {
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
		if _, err := m.source.Read(m.buffer[cursize:]); err != nil {
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
