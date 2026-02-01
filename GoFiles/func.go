func RepeatAlpha(s string) string {
	result := ""
	for _, lap := range s {
		var plus int
		if lap >= 'a' && lap <= 'z' {
			plus = int(lap - 'a' + 1)
		} else if lap >= 'A' && lap <= 'Z' {
			plus = int(lap - 'A' + 1)
		} else {
			result += string(lap)
			continue
		}
		for i := 0; i < plus; i++ {
			result += string(lap)
		}
	}
	return result
}