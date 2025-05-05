package buffer

import (
	"fmt"
	"io"
)

const defaultBufferSize = 4 * 1024

type Reader struct {
	source io.Reader
	buffer []byte
}

func NewReader(source io.Reader) *Reader {
	return &Reader{
		source: source,
		buffer: make([]byte, 0, defaultBufferSize),
	}
}

func (r *Reader) Slice(offset, length int) ([]byte, error) {
	var err error

	if len(r.buffer) < offset+length {
		_, err = r.source.Read(r.buffer[len(r.buffer) : offset+length])
	}

	return r.buffer[len(r.buffer) : offset+length], fmt.Errorf("%w", err)
}

func (r *Reader) Write(_ int, _ []byte) error {
	panic("cannot write to this buffer, it is read-only")
}
