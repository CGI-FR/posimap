package api

type Record interface {
	Unmarshal(buffer Buffer) (any, error)
	Marshal(buffer Buffer, value any) error
	Export(buffer Buffer, writer StructWriter) error
	Import(buffer Buffer, reader StructReader) (any, error)
}
