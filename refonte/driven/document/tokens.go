package document

type Token rune

const (
	TokenDocSep   Token = '\n'
	TokenObjStart Token = '{'
	TokenObjEnd   Token = '}'
	TokenArrStart Token = '['
	TokenArrEnd   Token = ']'
	TokenValSep   Token = ','
	TokenKey      Token = ':' // with String value
	TokenString   Token = '"' // with String value
	TokenNumber   Token = '0' // with Number value
	TokenTrue     Token = 't'
	TokenFalse    Token = 'f'
	TokenNull     Token = 'n'
	TokenEOF      Token = 0
)
