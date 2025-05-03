package decoder

type Decoder interface {
	Unmarshal(data Buffer, offset int) (any, int)
}
