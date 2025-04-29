package object

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

type JSONLineSource struct {
	scanner *bufio.Scanner
}

func NewJSONLineSource(reader io.Reader, encoding encoding.Encoding) JSONLineSource {
	scanner := bufio.NewScanner(transform.NewReader(reader, encoding.NewDecoder()))
	buf := make([]byte, 0, 64*1024)    //nolint:mnd
	scanner.Buffer(buf, 1024*1024*100) //nolint:mnd // increase buffer up to 100 MB

	return JSONLineSource{
		scanner: scanner,
	}
}

func (s JSONLineSource) Read() (any, error) {
	if !s.scanner.Scan() {
		if err := s.scanner.Err(); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return nil, io.EOF
	}

	var obj any
	if err := json.Unmarshal(s.scanner.Bytes(), &obj); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return obj, nil
}

func (s JSONLineSource) Close() error {
	if err := s.scanner.Err(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
