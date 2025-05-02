package flat

type Source interface {
	// ReadRunes reads a specified number of runes from the source.
	ReadRunes(length int) ([]rune, error)
	// ReadRunesUntil reads runes from the source until a specified separator is found.
	ReadRunesUntil(sep []byte) ([]rune, error)
	// ReadBytes reads a specified number of bytes from the source.
	ReadBytes(length int) ([]byte, error)
	// IsFixedWidth checks if the source has a fixed-width rune encoding.
	IsFixedWidth() bool
}
