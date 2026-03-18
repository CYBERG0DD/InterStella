func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	var result string
	step := 1
	if from > to {
		step = -1
	}

	for i := from; ; i += step {
		if i < 10 {
			result += "0" + fmt.Sprint(i)
		} else {
			result += fmt.Sprint(i)
		}
		if i == to {
			break
		}
		result += ", "
	}
	result += "\n"
	return result
}
