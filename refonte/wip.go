package refonte

import (
	"unicode/utf8"

	"golang.org/x/text/encoding"
)

type Buffer []byte

type String struct {
	prev *String
	next []String

	decoder *encoding.Decoder

	start int
	end   int

	length int
	value  []rune
}

func NewString(encoding encoding.Encoding, length int) *String {
	return &String{
		prev:    nil,
		next:    nil,
		decoder: encoding.NewDecoder(),
		start:   0,
		end:     0,
		length:  length,
		value:   nil,
	}
}

func (d *String) Then(next *String) *String {
	d.next = append(d.next, *next)
	next.prev = d

	return next
}

func (d *String) Unmarshal(data Buffer) {
	if d.prev != nil {
		d.prev.Unmarshal(data)
		d.start = d.prev.end
		d.end = d.start
	}

	d.value = nil

	working := make([]byte, utf8.UTFMax)

	for range d.length {
		raw := data[d.end : d.end+utf8.UTFMax]

		nDst, _, _ := d.decoder.Transform(working, raw, false)

		r, size := utf8.DecodeRune(working[:nDst])
		d.value = append(d.value, r)

		d.end += size
	}
}

func (d *String) Get() any {
	return d.value
}

func (d *String) String() string {
	if d.prev != nil {
		return d.prev.String() + "/" + string(d.value)
	}

	return string(d.value)
}
