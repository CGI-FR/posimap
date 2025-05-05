package api

type Decoder interface {
	Decode(buffer Buffer) (any, error)
}

type Encoder interface {
	Encode(buffer Buffer, value any) error
}
