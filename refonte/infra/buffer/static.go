package buffer

type Static struct {
	data []byte
}

func NewBufferStatic(data []byte) *Static {
	return &Static{
		data: data,
	}
}

func (b *Static) Peek(start, length int) []byte {
	if start+length > len(b.data) {
		return nil
	}

	return b.data[start : start+length]
}
