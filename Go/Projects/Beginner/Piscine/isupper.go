package piscine

func IsUpper(s string) bool {
	count := 0
	for _, val := range s {
		if !(val >= 'A' && val <= 'Z') {
			count++
		}
	}
	if count >= 1 {
		return false
	}
	return true
}
