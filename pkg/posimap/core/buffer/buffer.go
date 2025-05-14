package buffer

import (
	"errors"
	"fmt"
	"io"
	"math"

	"github.com/rs/zerolog/log"
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
	log.Trace().Int("size", 0).Int("cap", defaultBufferSize).Msg("Creating new buffer writer")

	return &Buffer{
		buffer: make([]byte, 0, defaultBufferSize),
		source: nil,
		target: target,
		locked: false,
	}
}

func NewBufferReader(source io.Reader) *Buffer {
	log.Trace().Int("size", 0).Int("cap", defaultBufferSize).Msg("Creating new buffer reader")

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
	if len(separator) == 0 || m.source == nil {
		return nil
	}

	size := len(m.buffer)

search:
	for inc := range math.MaxInt {
		if err := m.growTo(size + inc + 1); errors.Is(err, io.EOF) && inc > 0 {
			break // last separator is optional, if data has been read then we can stop
		} else if err != nil {
			return fmt.Errorf("%w", err)
		}

		for idx, sep := range separator {
			if m.buffer[size+inc+1-len(separator)+idx] != sep {
				continue search
			}
		}

		break
	}

	size = len(m.buffer)

	m.buffer = m.buffer[:size-len(separator)]
	m.locked = true

	log.Trace().Int("size", len(m.buffer)).Int("cap", cap(m.buffer)).Msg("Buffer locked to separator")

	return nil
}

func (m *Buffer) Reset(size int, separator ...byte) error {
	log.Trace().
		Int("size", len(m.buffer)).
		Int("cap", cap(m.buffer)).
		Msgf("Resetting buffer to %d bytes locking separator %q", size, separator)

	defer func() {
		log.Trace().Int("size", len(m.buffer)).Int("cap", cap(m.buffer)).Msg("Buffer reset")
	}()

	if m.target != nil && len(m.buffer) > 0 {
		if _, err := m.target.Write(m.buffer); err != nil {
			return fmt.Errorf("%w", err)
		}

		if _, err := m.target.Write(separator); err != nil {
			return fmt.Errorf("%w", err)
		}
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
