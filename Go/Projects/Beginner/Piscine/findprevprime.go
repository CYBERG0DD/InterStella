func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}

	// Check nb itself first
	if isPrime(nb) {
		return nb
	}

	// Go down until find prime
	for i := nb - 1; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}
	return 0
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	// Check odd divisors up to sqrt(n)
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
