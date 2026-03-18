package piscine

func HashCode(s string) string {
	length := len(s)
	result := []rune{}

	for _, ch := range s {
		hashed := (int(ch) + length) % 127
		if hashed < 32 {
			hashed += 33
		}
		result = append(result, rune(hashed))
	}

	return string(result)
}
