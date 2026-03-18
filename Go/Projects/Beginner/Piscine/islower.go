package piscine

func IsLower(s string) bool {
	count := 0
	for _, val := range s {
		if !(val >= 'a' && val <= 'z') {
			count++
		}
	}
	if count >= 1 {
		return false
	}
	return true
}
