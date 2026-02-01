package piscine

func RetainFirstHalf(str string) string {
	length := len(str) //set a variable
	if length == 0 {
		return ""
	}
	if length == 1 {
		return str
	}
	half := length / 2
	return str[:half]
}
