package api

type Decoder interface {
	Decode(buffer Buffer, offset int) (any, error)
	Size() int
}

type Encoder interface {
	Encode(buffer Buffer, offset int, value any) error
	Size() int
}

type Codec interface {
	Decoder
	Encoder
}
