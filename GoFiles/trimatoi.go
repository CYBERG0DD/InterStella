package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := false
	started := false

	for _, i := range s {

		if i == '-' && !started {
			sign = true
		}
		if i >= '0' && i <= '9' {
			sign = true
			result = result*10 + int(i-'0')
		}

	}
	if sign {
		return -result
	}
	return result
}
