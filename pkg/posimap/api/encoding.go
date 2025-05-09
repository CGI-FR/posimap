package api

type Decoder[T any] interface {
	Decode(buffer Buffer, offset int) (T, error)
	Size() int
}

type Encoder[T any] interface {
	Encode(buffer Buffer, offset int, value T) error
	Size() int
}

type Codec[T any] interface {
	Decoder[T]
	Encoder[T]
}
