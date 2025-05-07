package document

type Reader interface {
	ReadToken() (token Token, err error)
	ReadValue() (token Token, value any, err error)
	ReadString() (string, error)
	ReadNumber() (float64, error)
	ReadBool() (bool, error)
	ReadNull() error
}
