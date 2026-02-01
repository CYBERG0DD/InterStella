package piscine

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}
	divisibleBy2 := n%2 == 0
	divisibleBy3 := n%3 == 0

	if divisibleBy2 && divisibleBy3 {
		return "fish and chips"
	} else if divisibleBy2 {
		return "fish"
	} else if divisibleBy3 {
		return "chips"
	} else {
		return "error: non divisible"
	}
}
