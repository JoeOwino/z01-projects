package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	a "ascii/ascii_art"
)

const ErrorText = `Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"`

func main() {
	if len(os.Args) <= 1 {
		fmt.Println(ErrorText)
		os.Exit(1)
	}

	// Default arguments
	input := ""
	bannerFile := "standard.txt"
	subString := input

	flgColor := flag.String("color", "reset", "Color")
	flag.Parse()
	if (!strings.Contains(os.Args[1], "--color=") && flag.NFlag() == 1) || *flgColor == "" {
		fmt.Println(ErrorText)
		os.Exit(1)
	}
	color := a.ColorPicker(*flgColor)

	args := flag.Args()    // Non flags arguments
	nArgs := len(args)     // Count of non flag arguments
	nflags := flag.NFlag() // Count of flags

	if nflags == 0 {
		switch nArgs {
		case 1:
			input = args[0]
			subString = input
		case 2:
			input = args[0]
			bannerFile = args[1]
			subString = input
		default:
			fmt.Println(ErrorText)
			os.Exit(1)
		}
	} else {
		switch nArgs {
		case 1:
			input = args[0]
			subString = input
		case 2:
			subString = args[0]
			input = args[1]
		case 3:
			subString = args[0]
			input = args[1]
			bannerFile = args[2]
		default:
			fmt.Println(ErrorText)
			os.Exit(1)
		}
	}

	contents := a.GetFile(bannerFile)
	
	str := a.ProcessInput(contents, input, color, subString, "right")

	print(str)
}
