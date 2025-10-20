package piscine

func NRune(s string, n int) rune {
	if n < 0 || n >= len(s)+1 {
		return 0
	}
	for ch, r := range s {
		if ch == n-1 {
			return r
		}
	}
	return 0
}
