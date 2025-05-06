package document

type Token rune

const (
	TokenObjStart Token = '{'
	TokenObjEnd   Token = '}'
	TokenArrStart Token = '['
	TokenArrEnd   Token = ']'
	TokenString   Token = '"' // with string value
	TokenNumber   Token = '0' // with float64 value
	TokenTrue     Token = 't' // with bool value
	TokenFalse    Token = 'f' // with bool value
	TokenNull     Token = 'n'
)
