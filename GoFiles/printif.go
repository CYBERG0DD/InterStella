package piscine

func PrintIf(str string) string {

	if str == "" {
		return "G\n"
	}
	if len(str) > 3 {
		return "G\n"
	}
	return "Invalid Input"
}
