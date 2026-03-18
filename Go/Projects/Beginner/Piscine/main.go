package main // PACKAGE MAIN: Defines the package name.
// "main" is the entry point for executable Go programs.

import ( //Imports required standard libraries
	"fmt" // used for printing of texts
	"os"  // used of for command line arguments [??]
)

const size = 9 //const size = 9: Declared a constant for the board size.
// a constant in go is a built-in keyword for assigning values that cannot be changed

func main() { //func main() {} defines where the programs execution will start from.
	if len(os.Args) != size+1 { // checks the number of command-line arguments[they are the 9 strings representing each row of the Sudoku grid. Each string should be exactly 9 characters, using digits '1' to '9'] passed that is
		fmt.Println("Error") //exactly 9 rows and colums otherwise the programs exits
		return
	}

	board := [size][size]int{} // Initializes a 9x9 integer array to hold the puzzle board.

	// Parse and validate input
	for i := 0; i < size; i++ {
		row := os.Args[i+1]
		if len(row) != size {
			fmt.Println("Error")
			return
		} //Parses each row from command-line arguments.
		for j := 0; j < size; j++ { //Validates row length is 9.
			if row[j] == '.' { //Converts '.' to 0 (empty cell).
				board[i][j] = 0 //Converts digit characters '1' to '9' to integers.
			} else if row[j] >= '1' && row[j] <= '9' { //Prints "Error" and exits if invalid character found.
				board[i][j] = int(row[j] - '0')
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	if !isValidBoard(board) {
		fmt.Println("Error: invalid operation.")
		return
	}

	if !solve(&board) {
		fmt.Println("Error")
		return
	}

	printBoard(board)
}

func printBoard(board [size][size]int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(board[i][j])
			if j != size-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func isValidBoard(board [size][size]int) bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			val := board[i][j]
			if val != 0 {
				board[i][j] = 0
				if !isValid(board, i, j, val) {
					return false
				}
				board[i][j] = val
			}
		}
	}
	return true
}

func isValid(board [size][size]int, row, col, num int) bool {
	for i := 0; i < size; i++ {
		if board[row][i] == num {
			return false
		}
		if board[i][col] == num {
			return false
		}
	}
	boxRowStart := (row / 3) * 3
	boxColStart := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[boxRowStart+i][boxColStart+j] == num {
				return false
			}
		}
	}
	return true
}

func solve(board *[size][size]int) bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(*board, i, j, num) {
						board[i][j] = num
						if solve(board) {
							return true
						}
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return true
}
