package tetris

import (
	"fmt"
	"strings"
)

func CreateBoard(size int) [][]string {
	board := make([][]string, size)
	
	for i := 0; i < size; i++ {
		board[i] = make([]string, size)
		for j := 0; j < size; j++ {
			board[i][j] = "."
		}
	}
	return board
}

func PrintBoard(board [][] string) {
	for _, ln := range board {
		fmt.Println(strings.Join(ln, ""))
	}
}