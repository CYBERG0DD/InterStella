package piscine

func RectParameter(w, h int) int {
	if w < 0 || h < 0 {
		return -1
	}
	return 2 * (w + h)
}
