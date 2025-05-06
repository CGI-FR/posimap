package buffer

type Memory struct {
	buffer []byte
}

func NewMemory() *Memory {
	return &Memory{
		buffer: make([]byte, 0, defaultBufferSize),
	}
}

func (m *Memory) Slice(offset, length int) ([]byte, error) {
	return m.buffer[offset : offset+length], nil
}

func (m *Memory) Write(offset int, data []byte) error {
	requiredLength := offset + len(data)
	if len(m.buffer) < requiredLength {
		newBuffer := make([]byte, requiredLength)
		copy(newBuffer, m.buffer)

		for i := len(m.buffer); i < requiredLength; i++ {
			newBuffer[i] = ' '
		}

		m.buffer = newBuffer
	}

	copy(m.buffer[offset:], data)

	return nil
}

func (m *Memory) Bytes() []byte {
	return m.buffer
}
