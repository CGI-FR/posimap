package record

import (
	"bufio"
	"fmt"
	"io"

	"github.com/cgi-fr/posimap/pkg/data"
)

type Sink struct {
	writer *bufio.Writer
}

func NewRecordSink(writer io.Writer) Sink {
	return Sink{
		writer: bufio.NewWriter(writer),
	}
}

func (s Sink) Write(b *data.Buffer) error {
	if _, err := s.writer.WriteString(b.String()); err != nil {
		return fmt.Errorf("%w", err)
	}

	if _, err := s.writer.WriteRune('\n'); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := s.writer.Flush(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (s Sink) Close() error {
	return nil
}
