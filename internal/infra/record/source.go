package record

import (
	"bufio"
	"fmt"
	"io"

	"github.com/cgi-fr/posimap/pkg/data"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

type Source struct {
	scanner *bufio.Scanner
}

func NewRecordSource(reader io.Reader, encoding encoding.Encoding) Source {
	scanner := bufio.NewScanner(transform.NewReader(reader, encoding.NewDecoder()))
	buf := make([]byte, 0, 64*1024)    //nolint:mnd
	scanner.Buffer(buf, 1024*1024*100) //nolint:mnd // increase buffer up to 100 MB

	return Source{
		scanner: scanner,
	}
}

func (s Source) Read() (data.Buffer, error) {
	if !s.scanner.Scan() {
		if err := s.scanner.Err(); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return nil, io.EOF
	}

	return data.Buffer(s.scanner.Text()), nil
}

func (s Source) Close() error {
	if err := s.scanner.Err(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
