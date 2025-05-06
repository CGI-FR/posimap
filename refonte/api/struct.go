package api

type StructToken rune

type StructWriter interface {
	Resource
	WriteToken(token StructToken) error
	WriteString(data string) error
	WriteKey(key string) error
}

type StructReader interface {
	Resource
	PeekToken() (StructToken, error)
	ReadToken() (StructToken, error)
	ReadValue() (any, error)
	ReadKey() (string, error)
}
