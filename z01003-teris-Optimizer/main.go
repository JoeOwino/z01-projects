package main

import (
	"fmt"
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

	arrTetro := tetris.Assemble(arrTetris)

	l := 'A'

	for _, s := range arrTetro {
		//fmt.Println(string(l))

		for _, ln := range s {
			fmt.Println(ln)
		} 

		l++
		fmt.Println("---------------")
	}
}
