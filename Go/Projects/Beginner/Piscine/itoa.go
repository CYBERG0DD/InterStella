package piscine

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	var result string
	for n > 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n /= 10
	}
	return sign + result
}
