package piscine

func IsAlpha(s string) bool {
	spring := 0
	for _, val := range s {
		if !((val >= 'a' && val <= 'z') || (val >= 'A' && val <= 'Z') || (val >= '0' && val <= '9')) {
			spring++
		}
	}
	if spring >= 1 {
		return false
	}
	return true
}
