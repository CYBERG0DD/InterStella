package piscine

import "strings"

func ToSnakeCase(str string) string {
	if str == "" {
		return ""
	}
	for _, r := range str {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return str
		}
	}

	last := str[len(str)-1]
	if last >= 'A' && last <= 'Z' {
		return str
	}

	for i := 1; i < len(str); i++ {
		if str[i-1] >= 'A' && str[i-1] <= 'Z' && str[i] >= 'A' && str[i] <= 'Z' {
			return str
		}
	}

	var result strings.Builder
	for i := 0; i < len(str); i++ {
		if i > 0 && str[i] >= 'A' && str[i] <= 'Z' {
			result.WriteByte('_')
		}
		if str[i] >= 'A' && str[i] <= 'Z' {
			result.WriteByte(str[i] + 32)
		} else {
			result.WriteByte(str[i])
		}
	}
	return result.String()
}
