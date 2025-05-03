package buffer

type Static struct {
	data []byte
}

func NewStatic(data []byte) *Static {
	return &Static{
		data: data,
	}
}

func (b *Static) Peek(start, length int) []byte {
	if start < 0 || start >= len(b.data) {
		return nil
	}

	if start+length > len(b.data) {
		return b.data[start:]
	}

	return b.data[start : start+length]
}
