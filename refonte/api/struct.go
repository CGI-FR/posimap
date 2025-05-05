package api

type StructToken rune

const (
	StructTokenOpenRecord  StructToken = '('
	StructTokenCloseRecord StructToken = ')'
	StructTokenOpenObject  StructToken = '{'
	StructTokenCloseObject StructToken = '}'
	StructTokenOpenArray   StructToken = '['
	StructTokenCloseArray  StructToken = ']'
	StructTokenNext        StructToken = ','
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
	Skip() (StructToken, error)
	ReadString() (string, error)
	ReadKey() (string, error)
}
