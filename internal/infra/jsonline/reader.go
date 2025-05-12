package jsonline

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

type Reader struct {
	scanner *bufio.Scanner
}

func NewReader(reader io.Reader) *Reader {
	scanner := bufio.NewScanner(reader)
	buf := make([]byte, 0, 64*1024)    //nolint:mnd
	scanner.Buffer(buf, 1024*1024*100) //nolint:mnd // increase buffer up to 100 MB

	return &Reader{
		scanner: scanner,
	}
}

func (r *Reader) Read() (any, error) {
	if !r.scanner.Scan() {
		if err := r.scanner.Err(); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return nil, io.EOF
	}

	var obj any
	if err := json.Unmarshal(r.scanner.Bytes(), &obj); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return obj, nil
}

func (r *Reader) Close() error {
	if err := r.scanner.Err(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
