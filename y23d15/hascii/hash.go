package hascii

func Hash(data string) int {
	hash := 0

	for i := 0; i < len(data); i++ {
		hash += int(data[i])
		hash *= 17
		hash %= 256
	}

	return hash
}
