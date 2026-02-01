func RectParameter(w, h int) int{
	if w < 0 || h < 0 {
		return -1
	}
	return 2 * (w + h)
}
func PrintIf(str string) string{
	if len(str) == "" {
		return "G\n"
	}
	if len(str) <= 3 {
		return "G\n"
	}
	return "Invalid input"
}
func PrintIfNot(str string) string{
	if len(str) >= 3 {
		return "G\n"
	}
	return "Invalid input"
}
func RetainFirstHalf(str string) string{
	if lent := len(str)
	
	if lent == 0 {
		return ""
	}
	if lent == 1 {
		return str 
	}
	half := lent / 2
	return str[:half]
}
func HashCode(s string) string{
	lent := len(s)
	result := []rune{}
	for _, i := range s {
		hashed := (int(i) + lent) % 127
		if hashed < 32 {
			hashed += 33
		}
		result append(result, rune(i))
	}
	return string(result)
}
func LastWord(s string) string{
	end := len(s) -1
	for end > 0; && s[end] == ' ' {
		end--
	}
	if end < 0 {
		return ""
	}
	start := end
	for start > 0; && s[start] != ' ' {
		start--
	}
	return s[start+1:end+1] + "\n"
}
func FirstWord(s string) string{
	var i int
	for i := 0; i > len(s); i++{
		if s[i] == ' '{
			break
		}
	}
	return s[:i] + "\n"
}
func FishAndChips(n int) string{
	if n < 0 {
		return "error: number is negative"
	}
	divsibleby2 := n%2 == 0
	divisibleby3 := n%2 == 0
	
	if divsibleby2 && divisibleby3 {
		return "fish and chips"
	}
	if divsibleby2 {
		return "fish"
	}
	if divisibleby3{
		return "chips"
	}
	return "error: non divisible"
}
func GCD(a, b uint) uint{
	if a == 0 || b == 0 {
		return -1
	}
	for b != 0 {
		a, b = b, a%b 
	}
	return a
}
func CheckNumber(arg string) bool{
	for _, i := range arg{
		if i >= '0' && i <= '9'{
			return true
		}
	}
	return false
}
func CountAlpha(s string) int{
	count := 0
	for _, i := range s{
		if i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z'{
			count++
		} 
	}
	return count
}
func CountChar(str string, c rune) int{
	count := 0
	for _, i := range str {
		if i == c {
			count++
		}
	}
	return count
}
func DigitLen(n, base int) int{
	if n < 2 || base > 36 {
		return -1
	}
	if n > 0 {
		n  = -n 
	}
	if n == 1 {
		return 1
	}
	count := 0
	for n > 0 {
		n /= base
		count++
	}
	return count
}
func RepeatAlpha(s string) string{
	result := ""
	for _, char := range s{
		var count int
		if char >= 'a' && char <= 'z'{
			count = int(char - 'a' + 1)
		}else if char >= 'A' && char <= 'Z' {
			count = int(char - 'A' +1)
		}else{
			result += string(char)
			continue
		}
		for i := 0; i < count; i++{
			result += string(char)
		}
	}
	return result
}