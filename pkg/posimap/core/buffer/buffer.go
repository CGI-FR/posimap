package buffer

import (
	"errors"
	"fmt"
	"io"
)

var ErrIncompleteBuffer = errors.New("incomplete buffer")

const (
	growthFactor      = 2
	defaultBufferSize = 4 * 1024
)

type Buffer struct {
	buffer []byte
	source io.Reader
	target io.Writer
	locked bool
}

func NewBufferWriter(target io.Writer) *Buffer {
	return &Buffer{
		buffer: make([]byte, 0, defaultBufferSize),
		source: nil,
		target: target,
		locked: false,
	}
}

func NewBufferReader(source io.Reader) *Buffer {
	return &Buffer{
		buffer: make([]byte, 0, defaultBufferSize),
		source: source,
		target: nil,
		locked: false,
	}
}

func (m *Buffer) Slice(offset, length int) ([]byte, error) {
	if err := m.growTo(offset + length); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if offset > len(m.buffer) {
		return []byte{}, nil
	}

	if offset+length > len(m.buffer) {
		return m.buffer[offset:], nil
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

func (m *Buffer) LockTo(separator []byte) error {
	if len(separator) == 0 {
		return nil
	}

	size := len(m.buffer)

search:
	for {
		size++
		if err := m.growTo(size); err != nil {
			return fmt.Errorf("%w", err)
		}

		for idx, sep := range separator {
			if m.buffer[size-len(separator)+idx] != sep {
				continue search
			}
		}

		break
	}

	m.buffer = m.buffer[:size-len(separator)]
	m.locked = true

	return nil
}

func (m *Buffer) Reset(size int, separator ...byte) error {
	if m.target != nil {
		if _, err := m.target.Write(m.buffer); err != nil {
			return fmt.Errorf("%w", err)
		}

		if _, err := m.target.Write(separator); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	if size == 0 && !m.locked {
		size = len(m.buffer)
	}

	m.locked = false
	m.buffer = m.buffer[:0]

	// immediately refill the buffer either by locking to a separator
	// or by growing to specified length
	if err := m.LockTo(separator); err != nil {
		return fmt.Errorf("%w", err)
	}

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
	if cursize >= size || m.locked {
		return nil
	}

	if !m.tryReslice(size) {
		capacity := max(size, growthFactor*cap(m.buffer))
		newBuffer := append([]byte(nil), make([]byte, capacity)...)
		copy(newBuffer, m.buffer)
		m.buffer = m.buffer[:size]
	}

	if m.source != nil {
		if n, err := m.source.Read(m.buffer[cursize:]); err != nil {
			return fmt.Errorf("%w", err)
		} else if n < size-cursize {
			return fmt.Errorf("%w", ErrIncompleteBuffer)
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
