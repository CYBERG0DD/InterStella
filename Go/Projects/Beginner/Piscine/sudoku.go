package main // PACKAGE MAIN: Defines the package name.
// "main" is the entry point for executable Go programs.

import (
	"fmt" // used for printing of texts to the console or terminal
	"os"  // used for command line arguments
)

// Board represents a 9x9 Sudoku grid
type Board [9][9]byte

// isValid checks if placing num at board[row][col] is valid
func (b *Board) isValid(row, col int, num byte) bool {
	n := int(num - '1')

	// Check row
	for c := 0; c < 9; c++ {
		if b[row][c] != '.' && int(b[row][c]-'1') == n {
			return false
		}
	}

	// Checks the column
	for r := 0; r < 9; r++ {
		if b[r][col] != '.' && int(b[r][col]-'1') == n {
			return false
		}
	}

	// Check if its actully 3x3 box
	boxRow, boxCol := row/3*3, col/3*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[boxRow+i][boxCol+j] != '.' && int(b[boxRow+i][boxCol+j]-'1') == n {
				return false
			}
		}
	}
	return true
}

// solve uses backtracking to find unique solution
func (b *Board) solve() bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if b[r][c] == '.' {
				for num := byte('1'); num <= '9'; num++ {
					if b.isValid(r, c, num) {
						b[r][c] = num
						if b.solve() {
							return true
						}
						b[r][c] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

// countSolutions verifies exactly one solution exists
func (b *Board) countSolutions() int {
	count := 0
	var solveHelper func() bool

	solveHelper = func() bool {
		for r := 0; r < 9; r++ {
			for c := 0; c < 9; c++ {
				if b[r][c] == '.' {
					for num := byte('1'); num <= '9'; num++ {
						if b.isValid(r, c, num) {
							b[r][c] = num
							if solveHelper() {
								return true
							}
							b[r][c] = '.'
						}
					}
					return false // returns false and exits the function
				}
			}
		}
		return true // found a solution
	}

	// First solution
	if solveHelper() {
		count++
		// Reset and find second solution
		b.resetEmpty()
		if solveHelper() {
			count++
		}
	}
	return count
}

// resetEmpty sets all empty cells back to '.'
func (b *Board) resetEmpty() {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if b[r][c] == '1' || b[r][c] == '2' || b[r][c] == '3' ||
				b[r][c] == '4' || b[r][c] == '5' || b[r][c] == '6' ||
				b[r][c] == '7' || b[r][c] == '8' || b[r][c] == '9' {
				b[r][c] = '.'
			}
		}
	}
}

func main() {
	// Read exactly 9 arguments
	args := os.Args[1:]
	if len(args) != 9 {
		fmt.Println("Error")
		return
	}

	// Parse board
	var board Board
	for i, arg := range args {
		if len(arg) != 9 {
			fmt.Println("Error")
			return
		}
		for j, ch := range arg {
			if (ch < '1' || ch > '9') && ch != '.' {
				fmt.Println("Error")
				return
			}
			board[i][j] = ch
		}
	}

	// Check unique solution exists
	if board.countSolutions() != 1 {
		fmt.Println("Error")
		return
	}

	// Print solved board
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			fmt.Printf("%c", board[r][c])
			if c < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
