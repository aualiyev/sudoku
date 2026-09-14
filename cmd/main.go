package main

import (
	"fmt"
	"os"

	sudoku "piscine/internal"
)

func main() {
	args := os.Args[1:]

	board, _ := sudoku.ParseSudoku(args)

	solved, _ := sudoku.SolveSudoku(board)

	printedBoard := sudoku.GetPrintedBoard(solved)

	fmt.Println(printedBoard)
}
