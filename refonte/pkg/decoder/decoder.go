package decoder

type Decoder interface {
	Unmarshal(d Node, data Buffer) any
	String(node Node, data Buffer) string
}
