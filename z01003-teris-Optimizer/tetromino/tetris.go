package tetris

import "fmt"

func Tetris(board [][]string, n, c int, arrAdded []int) [][]string {
	tetromino := Tetromino()
	size := len(tetromino)

	if n == size {
		return board
	}

	tetro := tetromino[n]
	prevx := arrAdded[len(arrAdded)-2]
	prevy := arrAdded[len(arrAdded)-1]

	for y := 0; y < size; y++ {
		for x := 0; x < len(board); x++ {

			if x == prevx && y == prevy {
				continue
			}

			if IsAddable(tetro, board, x, y) {
				AddTetro(board, tetro, x, y)
				arrAdded := append(arrAdded, x, y)
				return Tetris(board, n+1, c, arrAdded)
			}
		}
	}

	PrintBoard(board)
	fmt.Println(arrAdded)
	fmt.Println("###################")


	if n == 0 {
		board = CreateBoard(size + 1)
		return Tetris(board, n, c, arrAdded)
	}

	n--
	tetro = tetromino[n]

	RemoveTetro(board, tetro, prevx, prevy)
	arrAdded = arrAdded[:len(arrAdded)-2]

	PrintBoard(board)
	fmt.Println(arrAdded)
	fmt.Println(prevx, prevy)
	fmt.Println("___________")

	if c == 10 {
		return board
	} 
	c++

	return (Tetris(board, n, c, arrAdded))
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
