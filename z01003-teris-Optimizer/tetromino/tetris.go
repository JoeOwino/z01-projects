package tetris

func Tetris(board [][]string, n, x, y int) [][]string {
	if n == 8 {
		return board
	}

	tetromino := Tetromino()
	tetro := tetromino[n]

	for y < len(board) {
		arrxy := []int{}
		for x < len(board) {
			if IsAddable(tetro, board, x, y) {
				AddTetro(board, tetro, x, y)
				if n == len(tetromino)-1 {
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
	size := len(board) + 1

	for n < len(tetromino) {
		if isPlaced {
			arrAddedTetros = append(arrAddedTetros, arrxy)
			n++
			Tetris(board, n, 0, 0)
		} else {
			if n == 0 {
				size++
				board = CreateBoard(size)
				Tetris(board, n, 0, 0)
			} else {
				n--
				prevAded := arrAddedTetros[len(arrAddedTetros)-1]
				i := prevAded[0]
				j := prevAded[1]
				RemoveTetro(board, tetromino[n], i, j)
				arrAddedTetros = arrAddedTetros[:len(arrAddedTetros)-1]
				if j < size-1 {
					j++
					i = 0
				} else {
					j = 0
					i++
				}

				Tetris(board, n, i, j)
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
