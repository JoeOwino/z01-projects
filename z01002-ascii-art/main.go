package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	a "ascii/ascii_art"
)

const strUsage = `Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"`

func main() {
	if len(os.Args) <= 1 {
		fmt.Println(strUsage)
		os.Exit(1)
	}

	// Default arguments
	input := ""
	bannerFile := "standard.txt"
	subString := input

	flgColor := flag.String("color", "reset", "Color")
	flag.Parse()

	if (!strings.Contains(os.Args[1], "--color=") && flag.NFlag() == 1) || *flgColor == "" {
		fmt.Println(strUsage)
		os.Exit(1)
	}
	color := a.ColorPicker(*flgColor)


	args := flag.Args()
	nArgs := len(args)

	switch nArgs {
	case 1:
		input = args[0]
		subString = input
	case 2:
		if flag.NFlag() == 0 {
			input = args[0]
			bannerFile = args[1]
			subString = input
		} else {
			subString = args[0]
			input = args[1]
		}
	case 3:
		if flag.NFlag() == 1 {
			subString = args[0]
			input = args[1]
			bannerFile = args[2]
		} else {
			fmt.Println(strUsage)
			os.Exit(1)
		}
	default:
		fmt.Println(strUsage)
		os.Exit(1)
	}

	contents := a.GetFile(bannerFile)

	fmt.Print(a.ProcessInput(contents, input, color, subString))
}
