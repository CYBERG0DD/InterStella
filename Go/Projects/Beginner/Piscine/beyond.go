package main

import (
	"fmt"
)

type Candidate struct {
	physicalAptitudes int
	noFamily          bool
}

func AssignDestination(candidate Candidate) string {
	if candidate.physicalAptitudes < 80 {
		return "earth"
	}
	if candidate.noFamily {
		return "mars"
	}
	return "moon"
}

func LogDestination(candidate Candidate) {
	destination := AssignDestination(candidate)
	fmt.Printf("Candidate can access: %s\n", destination)
}

func main() {
	// Test cases
	marsCandidate := Candidate{physicalAptitudes: 90, noFamily: true}
	moonCandidate := Candidate{physicalAptitudes: 85, noFamily: false}
	earthCandidate := Candidate{physicalAptitudes: 70, noFamily: true}

	fmt.Println("=== Mars Candidate (90, no family) ===")
	LogDestination(marsCandidate)

	fmt.Println("=== Moon Candidate (85, has family) ===")
	LogDestination(moonCandidate)

	fmt.Println("=== Earth Candidate (70, no family) ===")
	LogDestination(earthCandidate)
}

//To infinity and beyond!
