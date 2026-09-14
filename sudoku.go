package piscine

func ParseSudoku(args []string) ([][]int, string) {
	board := [][]int{}

	if len(args) != 9 {
		return nil, "Error"
	}

	for _, arg := range args {
		if len(arg) != 9 {
			return nil, "Error"
		}
		row := []int{}
		for _, r := range arg {
			if r >= '1' && r <= '9' || r == '.' {
				if r == '.' {
					row = append(row, 0)
				} else {
					row = append(row, int(r-48))
				}
			} else {
				return nil, "Error"
			}
		}
		board = append(board, row)
	}

	return board, ""
}

func GetPrintedBoard(board [][]int) string {
	if board == nil {
		return "Error"
	}
	res := ""
	// allowedMin := 0 // for testing print (new parsed board wihout solving)
	allowedMin := 1 // for solved print

	for _, row := range board {
		str := []rune{}
		for i, num := range row {
			isValid := num >= allowedMin && num <= 9
			if !isValid {
				return "Error"
			}
			str = append(str, rune(num+48))

			if i < len(row)-1 {
				str = append(str, ' ')
			}
		}
		str = append(str, '\n')
		res += string(str)
	}

	return res
}

func SolveSudoku(board [][]int) ([][]int, string) {
	if board == nil {
		return nil, "Error"
	}

	solved := copyBoard(board)

	if solve(solved) {
		return solved, ""
	}

	return nil, "Error"
}

// рекурсивный бэктрекинг
func solve(board [][]int) bool {
	row, col, found := findEmpty(board)

	// если пустая ячейка не найдена, то пропускает
	if !found {
		return true
	}

	// когда найдена пустая ячейка, то подставляет цифры
	for num := 1; num <= 9; num++ {
		// проверяет подходит ли цифра
		if isValid(board, row, col, num) {
			board[row][col] = num

			// если цифра подходит
			if solve(board) {
				return true
			}

			board[row][col] = 0 // если цифра не подошла, то откатывается назад
		}
	}

	return false
}

// ищет первую пустую клетку (0)
func findEmpty(board [][]int) (int, int, bool) { // возвращает положение ячейки(i, j) и true/false
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return 0, 0, false
}

// проверяет можно ли поставить num в row,col (т.е повторяется ли num по строке/столбцу/квадрат 3x3)
func isValid(board [][]int, row, col, num int) bool {
	// проверка строки
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}

	// проверка столбца
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}

	// проверка квадрата 3x3
	startRow := (row / 3) * 3 // берет начало (по строке) квадрата 3x3
	startCol := (col / 3) * 3 // берет начало (по столбцу) квадрата 3x3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

// копирует доску
func copyBoard(board [][]int) [][]int {
	copied := make([][]int, len(board))

	for i := range board {
		copied[i] = make([]int, len(board[i]))
		copy(copied[i], board[i])
	}

	return copied
}
