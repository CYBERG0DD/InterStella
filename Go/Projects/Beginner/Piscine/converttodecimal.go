package main

import (
	"fmt"
	"strconv"
)

func ConvertToDecimal(num string, base int) (int64, error) {
	result, err := strconv.ParseInt(num, base, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func main() {

	hex, _ := ConvertToDecimal("1E", 16)
	oct, _ := ConvertToDecimal("FFFF", 8)
	bin, _ := ConvertToDecimal("10", 2)

	fmt.Println(hex, oct, bin)
}
