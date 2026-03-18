package piscine

func StrRev(s string) string {
	reverse := ""
	for _, char := range s {
		reverse = string(char) + reverse
	}
	return reverse
}
