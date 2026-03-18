package piscine

func RepeatAlpha(s string) string {
	result := ""
	for _, ch := range s {
		var count int
		if ch >= 'a' && ch <= 'z' {
			count = int(ch - 'a' + 1)
		} else if ch >= 'A' && ch <= 'Z' {
			count = int(ch - 'A' + 1)
		} else {
			result += string(ch)
			continue
		}
		for i := 0; i < count; i++ {
			result += string(ch)
		}
	}
	return result
}
