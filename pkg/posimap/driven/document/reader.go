package document

type Reader interface {
	Read() (any, error)
	Close() error
}
