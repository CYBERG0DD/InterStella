package piscine

func Swap(a *int, b *int) {
	value1 := *a
	value2 := *b
	*a = value2
	*b = value1
}
