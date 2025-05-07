package buffer

import "io"

const (
	growthFactor      = 2
	defaultBufferSize = 4 * 1024
)

type Memory struct {
	buffer []byte
	source io.Reader
}

func NewMemory() *Memory {
	return &Memory{
		buffer: make([]byte, 0, defaultBufferSize),
		source: nil,
	}
}

func NewMemoryWithSource(source io.Reader) *Memory {
	return &Memory{
		buffer: make([]byte, 0, defaultBufferSize),
		source: source,
	}
}

func (m *Memory) Slice(offset, length int) ([]byte, error) {
	m.Required(offset + length)

	return m.buffer[offset : offset+length], nil
}

func (m *Memory) Write(offset int, data []byte) error {
	m.Required(offset + len(data))

	copy(m.buffer[offset:], data)

	return nil
}

func (m *Memory) Reset() {
	size := len(m.buffer)
	m.buffer = m.buffer[:0]
	m.Required(size) // immediately refill the buffer
}

func (m *Memory) Bytes() []byte {
	return m.buffer
}

func (m *Memory) Required(size int) {
	cursize := len(m.buffer)
	if cursize >= size {
		return
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
			panic(err)
		}
	}
}

// tryReslice is an inlineable version for the fast-case where the
// internal buffer only needs to be resliced.
// It returns whether it succeeded.
func (m *Memory) tryReslice(size int) bool {
	if size <= cap(m.buffer) {
		m.buffer = m.buffer[:size]

		return true
	}

	return false
}
