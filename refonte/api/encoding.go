package api

type Decoder interface {
	Decode(buffer Buffer, offset int) (any, error)
}

type Encoder interface {
	Encode(buffer Buffer, offset int, value any) error
}

type Codec interface {
	Decoder
	Encoder
}
