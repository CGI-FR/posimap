package data2

type Source interface {
	// ReadRunes reads a specified number of runes from the source.
	// It returns a slice of runes and an error if any occurs.
	ReadRunes(length int) ([]rune, error)
	// ReadBytes reads a specified number of bytes from the source.
	// It returns a slice of bytes and an error if any occurs.
	ReadBytes(length int) ([]byte, error)
}
