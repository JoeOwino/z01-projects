package tetris

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func Tetromino() [][]string {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [FILENAME]\n\nEX: go run . sample.txt")
		os.Exit(1)
	}

	fileName := os.Args[1]
	arrTetris := ReadFile(fileName)

	return Assemble(arrTetris)
}

func BoardSize() (size int) {
	tetromino := Tetromino()
	n := float64(len(tetromino))
	size = int(math.Ceil(math.Sqrt(n * 4)))

	return size
}

func ReadFile(file string) []string {
	tetromino, err := (os.ReadFile(file))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	arrTetro := strings.Split(string(tetromino), "\n\n")
	if len(arrTetro) > 26 || len(tetromino) < 1 {
		fmt.Println("Invalid Tetromino File")
		os.Exit(0)
	}

	return arrTetro
}

func Assemble(tetromino []string) (arrTetris [][]string) {
	l := 'A'

	for _, square := range tetromino {
		if square == "" {
			continue
		}

		if strings.Count(square, "#") != 4 || (strings.Count(square, ".")+strings.Count(square, "\n")+strings.Count(square, "#") != len(square)) {
			fmt.Println("Invalid Tetromino")
			os.Exit(0)
		}

		square := strings.ReplaceAll(square, "#", string(l))
		arrSquare := strings.Split(square, "\n")
		if !isValidTetro(arrSquare) {
			fmt.Println("Invalid Tetromino")
			os.Exit(0)
		}

		arrSquare = trimRows(arrSquare)
		arrSquare = trimColumns(arrSquare)

		//(arrSquare)

		arrTetris = append(arrTetris, arrSquare)

		l++
	}

	return arrTetris
}

func isValidTetro(tetro []string) bool {
	if len(tetro) != 4 {
		return false
	}

	connCount := 0
	for i, ln := range tetro {
		if len(ln) != 4 {
			return false
		}

		for j, ch := range ln {
			if ch == '.' {
				continue
			}

			if i > 0 && ch == rune(tetro[i-1][j]) {
				connCount++
			}

			if i < len(ln)-1 && ch == rune(tetro[i+1][j]) {
				connCount++
			}

			if j > 0 && ch == rune(tetro[i][j-1]) {
				connCount++
			}

			if j < len(ln)-1 && ch == rune(tetro[i][j+1]) {
				connCount++
			}
		}
	}

	return connCount == 6 || connCount == 8
}

func trimColumns(tetro []string) []string {
	trimedTetro := []string{}

	for _, ln := range tetro {
		if strings.Count(ln, ".") != len(ln) {
			trimedTetro = append(trimedTetro, ln)
		}
	}

	return trimedTetro
}

func trimRows(tetro []string) []string {
	trimedTetro := tetro
	i := 0

	for i < len(trimedTetro[1]) {
		if tetro[0][i] == '.' && tetro[1][i] == '.' && tetro[2][i] == '.' && tetro[3][i] == '.' {
			for j := 0 ; j < 4; j++ {
				trimedTetro[j] = trimedTetro[j][:i] + trimedTetro[j][i+1:]
			}
		} else {
			i++
		}
	}

	return trimedTetro
}
