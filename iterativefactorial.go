package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	film := 1
	for b := 1; b <= nb; b++ {
		film *= b
	}
	return film
}
