package decoder

type Decoder interface {
	Unmarshal(d Node, data Buffer, offset int) (any, int)
}
