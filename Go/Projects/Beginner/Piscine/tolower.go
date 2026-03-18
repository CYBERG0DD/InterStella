package piscine

func ToLower(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'A' && c <= 'Z' {
			b[i] = c + 32 // or c - ('A' - 'a')
		}
	}
	return string(b)
}
