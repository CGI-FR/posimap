package api

type Buffer interface {
	Slice(offset, length int) ([]byte, error)
	Write(offset int, data []byte) error
}
