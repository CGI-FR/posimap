package decoder

type Buffer interface {
	Peek(start, length int) []byte
}
