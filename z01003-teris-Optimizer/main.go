package main

import (
	tetris "tetris/tetromino"
)

func main() {
	// tetromino := tetris.Tetromino()
	size := tetris.BoardSize()

	board := tetris.CreateBoard(int(size))

	board = tetris.Tetris(board, 0)

	tetris.PrintBoard(board)

	// tetris.Assemble(arrTetris)
}
