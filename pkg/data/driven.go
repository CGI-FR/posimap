package data

type RecordSource interface {
	Read() (*Buffer, error)
	Close() error
}

type RecordSink interface {
	Write(b *Buffer) error
	Close() error
}

type ObjectSource interface {
	Read() (any, error)
	Close() error
}

type ObjectSink interface {
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

type ExportPredicate func(root View, context *Buffer) bool
