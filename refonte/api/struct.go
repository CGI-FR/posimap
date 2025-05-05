package api

type StructToken rune

const (
	StructTokenRecordStart StructToken = '('
	StructTokenRecordEnd   StructToken = ')'
	StructTokenObjectStart StructToken = '{'
	StructTokenObjectEnd   StructToken = '}'
	StructTokenArrayStart  StructToken = '['
	StructTokenArrayEnd    StructToken = ']'
	StructTokenSeparator   StructToken = ','
	StructTokenKey         StructToken = ':'
	StructTokenString      StructToken = '"'
)

type StructWriter interface {
	Resource
	WriteToken(token StructToken) error
	WriteString(data string) error
	WriteKey(key string) error
}

type StructReader interface {
	Resource
	Peek() (StructToken, error)
	Read() (StructToken, error)
	ReadString() (string, error)
	ReadKey() (string, error)
}
