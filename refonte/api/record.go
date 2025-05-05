package api

type Record interface {
	Unmarshal(buffer Buffer) error
	Marshal(buffer Buffer) error
	Export(writer StructWriter) error
	Import(reader StructReader) error
	AsPrimitive() any
}
