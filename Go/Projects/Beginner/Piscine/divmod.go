package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b // Integer division quotient
	*mod = a % b // Remainder (modulus)
}
