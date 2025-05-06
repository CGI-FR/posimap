package document

type Writer interface {
	WriteToken(token Token) error
	WriteValue(token Token, value any) error
	WriteKey(key string) error
	WriteString(value string) error
	WriteNumber(value float64) error
	WriteBool(value bool) error
	WriteNull() error
	WriteEOF() error
}
