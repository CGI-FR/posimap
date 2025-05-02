package deep

type Source interface {
	Read() (any, error)
	Close() error
}
