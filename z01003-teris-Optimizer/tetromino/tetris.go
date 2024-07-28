package tetris

import (
	"fmt"
)

func Tetris(board [][]string, n, c int, placedMap map[int][]int) [][]string {
	tetromino := Tetromino()
	size := len(tetromino)

	if n == size {
		return board
	}

	tetro := tetromino[n]
	arrAdded := placedMap[n]

	fmt.Println(n)
	fmt.Println("___________")

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {

			fmt.Println(tetro, ":", x, y, IsAddable(tetro, board, x, y, arrAdded))

			if IsAddable(tetro, board, x, y, arrAdded) {
				AddTetro(board, tetro, x, y)
				placedMap[n] = append(arrAdded, x, y)

				// PrintMap(placedMap)
				// fmt.Println("___________")

				return Tetris(board, n+1, c, placedMap)
			}
		}
	}

	// PrintBoard(board)
	// fmt.Println(arrAdded)
	// fmt.Println("###################")

	if n == 0 {
		board = CreateBoard(size + 1)
		return Tetris(board, n, c,  map[int][]int{-1 : {-1, -1}})
	}

	n--
	tetro = tetromino[n]
	arrAdded = placedMap[n]
	prevx, prevy := arrAdded[len(arrAdded)-2], arrAdded[len(arrAdded)-1]
	RemoveTetro(board, tetro, prevx, prevy)

	// PrintBoard(board)
	// fmt.Println(arrAdded)
	// fmt.Println(prevx, prevy)
	// fmt.Println("___________")

	// if c == 10 {
	// 	return board
	// }
	// c++

	return (Tetris(board, n, c, placedMap))
}

func IsAddable(tetro []string, board [][]string, x, y int, arrAdded []int) bool {
	if len(board)-y < len(tetro) || len(board[0])-x < len(tetro[0]) {
		return false
	}

	if len(arrAdded) >= 2 && arrAdded[0] == x && arrAdded[1] == y {
		//fmt.Println(arrAdded, ":", tetro, arrAdded[0] == x && arrAdded[1] == y)
		return false
	}

	for i := 2; i < len(arrAdded)-2; i = i + 2 {
		if arrAdded[i] == x && arrAdded[i+1] == y {
			return false
		}
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

func PrintMap(tmap map[int][]int) {
	for k, ln := range tmap {
		fmt.Println(k, ":", ln)
	}
}
