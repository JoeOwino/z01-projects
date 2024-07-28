package main

import (
	tetris "tetris/tetromino"
)

func main() {
	// tetromino := tetris.Tetromino()
	size := tetris.BoardSize()

	board := tetris.CreateBoard(int(size))

	board = tetris.Tetris(board, 0, 0, []int{-1, -1})

	//tetris.PrintBoard(board)

	//println("___________________________")

	//tetris.PrintBoard(tetris.RemoveTetro(board, tetris.Tetromino()[4], 4, 1))

	// tetris.Assemble(arrTetris)
}
