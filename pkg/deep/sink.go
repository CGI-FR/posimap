package deep

type Sink interface {
	OpenRecord() error
	CloseRecord() error
	OpenObject() error
	CloseObject() error
	OpenArray() error
	CloseArray() error
	Next() error
	WriteString(data string) error
	WriteKey(key string) error
	Close() error
}
