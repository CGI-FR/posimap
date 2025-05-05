package api

type Resource interface {
	Open() error
	Close() error
}
