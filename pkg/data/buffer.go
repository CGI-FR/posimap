package data

type Buffer []rune

func (b Buffer) Read(start, length int) Buffer {
	if start >= len(b) || start < 0 {
		return b[0:0]
	}

	if length == 0 {
		return b[start:]
	}

	if start+length > len(b) {
		return b[start:]
	}

	return b[start : start+length]
}

func (b Buffer) Write(start, length int, value string) error {
	if start >= len(b) || start < 0 {
		return nil
	}

	done := length

	for idx, r := range value {
		b[start+idx] = r

		if done--; done == 0 {
			break
		}
	}

	for idx := range done {
		b[start+done+idx] = ' '
	}

	return nil
}

func (b Buffer) String() string {
	return string(b)
}
