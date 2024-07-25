package main

import (
	"fmt"
	"math"
	"os"

	tetris "tetris/tetromino"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [FILENAME]\n\nEX: go run . sample.txt")
		return
	}
	fileName := os.Args[1]
	arrTetris := tetris.ReadFile(fileName)

	tetromino := tetris.Assemble(arrTetris)
	n := float64(len(tetromino))
	size := math.Ceil(math.Sqrt(n * 4))

	board := tetris.CreateBoard(int(size))

	board = tetris.Tetris(board, tetromino)


	tetris.PrintBoard(board)

	//tetris.Assemble(arrTetris)
	
}
