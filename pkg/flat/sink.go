package flat

import (
	"bufio"
	"fmt"
	"io"
)

type Sink struct {
	writer *bufio.Writer
}

func NewSink(writer io.Writer) Sink {
	return Sink{
		writer: bufio.NewWriter(writer),
	}
}

func (s Sink) Write(b []rune) error {
	if _, err := s.writer.WriteString(string(b)); err != nil {
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
	if err := s.writer.Flush(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
