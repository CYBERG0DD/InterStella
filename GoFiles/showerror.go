package main

import (
	"fmt"
	"strconv"
)

func ParseNumber(s string, base int) (int64, error) {
	result, err := strconv.ParseInt(s, base, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func main() {

	fmt.Println(ParseNumber("FF", 16))
	fmt.Println(ParseNumber("10101", 2))
	fmt.Println(ParseNumber("EFF", 8))
	fmt.Println(ParseNumber("37", 1))
	fmt.Println(ParseNumber("99", 2))
	fmt.Println(ParseNumber("99", 1))
}
