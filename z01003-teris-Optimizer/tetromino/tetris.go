package tetris

func Tetris(board [][]string, n int) [][]string {
	tetromino := Tetromino()
	tetro := tetromino[n]

	y := 0
	for y < len(board) {
		x := 0
		arrxy := []int{}
		for x < len(board) {
			if IsAddable(tetro, board, x, y) {
				AddTetro(board, tetro, x, y)
				if n == len(tetromino) - 1 {
					return board
				}
				arrxy = append(arrxy, x, y)
				Iterator(board, true, arrxy, n)
			}
			x++
		}
		y++
	}

	return board
}

func Iterator(board [][]string, isPlaced bool, arrxy []int, n int) {
	tetromino := Tetromino()
	arrAddedTetros := [][]int{}

	for n < len(tetromino) {
		if isPlaced {
			arrAddedTetros = append(arrAddedTetros, arrxy)
			n++
			Tetris(board, n)
		} else {
			if n == 0 {
				size := len(board)
				size++
				board = CreateBoard(size)
				Tetris(board, n)
			} else {
				n--
				prevAded := arrAddedTetros[len(arrAddedTetros)-1]
				RemoveTetro(board, tetromino[n], prevAded[0], prevAded[1])
			}
		}
	}
}

// func Tetris(board, tetromino [][]string) [][]string {
// 	n := 0
// 	x := 0
// 	for x < len(board) {
// 		y := 0
// 		isPlaced := false

// 		for y < len(board) {
// 			if IsAddable(tetromino[n], board, x, y) {
// 				AddTetro(board, tetromino[n], x, y)
// 				n++
// 				isPlaced = true
// 				break
// 			}
// 			y++
// 		}
// 		if !isPlaced {
// 			x++
// 		}
// 	}

// 	return board
// }

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
