package deep

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

type SourceJSONLine struct {
	scanner *bufio.Scanner
}

func NewSourceJSONLine(reader io.Reader, encoding encoding.Encoding) SourceJSONLine {
	scanner := bufio.NewScanner(transform.NewReader(reader, encoding.NewDecoder()))
	buf := make([]byte, 0, 64*1024)    //nolint:mnd
	scanner.Buffer(buf, 1024*1024*100) //nolint:mnd // increase buffer up to 100 MB

	return SourceJSONLine{
		scanner: scanner,
	}
}

func (s SourceJSONLine) Read() (any, error) {
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

func (s SourceJSONLine) Close() error {
	if err := s.scanner.Err(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
