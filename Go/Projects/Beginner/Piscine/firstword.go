package piscine

func FirstWord(s string) string {
	var i int
	for i = 0; i < len(s); i++ {
		if s[i] == ' ' {
			break
		}
	}
	return s[:i] + "\n"
}
