package piscine

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}
	i := 0
	for i < len(s) {
		for i < len(s) && s[i] == ' ' {
			i++
		}
		if i == len(s) {
			break
		}
		first := s[i]
		if first >= 'a' && first <= 'z' {
			return false
		}
		for i < len(s) && s[i] != ' ' {
			i++
		}
	}
	return true
}
