package tetris

func Tetris(board, tetromino [][]string) [][]string {
	n := 0
	x := 0
	for x < len(board) {
		y := 0
		isPlaced := false

		for y < len(board) {
			if IsAddable(tetromino[n], board, x, y) {
				AddTetro(board, tetromino[n], x, y)
				n++
				isPlaced = true
				break
			}
			y++
		}
		if !isPlaced {
			x++
		}
	}

	return board
}

func IsAddable(tetro []string, board [][]string, x, y int) bool {
	if len(board)-y < len(tetro) || len(board[0])-x < len(tetro[0]) {
		return false
	}

	for i, ln := range tetro {
		for j := range ln {
			if board[i+y][j+x] != "." {
				return false
			}
		}
	}
	return true
}

func AddTetro(board [][]string, tetro []string, x, y int) [][]string {
	for i, ln := range tetro {
		for j, ch := range ln {
			board[i+y][j+x] = string(ch)
		}
	}
	return board
}

func RemoveTetro(board [][]string, tetro []string, x, y int) [][]string {
	for i, ln := range tetro {
		for j := range ln {
			board[i+y][j+x] = "."
		}
	}
	return board
}
