package record

import (
	"fmt"

	"github.com/cgi-fr/posimap/pkg/posimap/api"
	"github.com/cgi-fr/posimap/pkg/posimap/driven/document"
)

type Filler struct {
	offset  int
	encoder api.Encoder[string]
}

func NewFiller(offset int, codec api.Codec[string]) *Filler {
	return &Filler{
		offset:  offset,
		encoder: codec,
	}
}

func (f *Filler) Unmarshal(_ api.Buffer) error {
	return nil
}

func (f *Filler) Marshal(buffer api.Buffer) error {
	err := f.encoder.Encode(buffer, f.offset, "")
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// func (f *Filler) export(_ document.Writer, _ Record) error {
// 	return nil
// }

func (f *Filler) Export(_ document.Writer) error {
	return nil
}
