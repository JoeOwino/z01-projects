package ascii_art

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetFile reads from the file specified by filename and returns its contents
func GetFile(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(0)
	}

	if len(file) == 0 {
		fmt.Println("Error: The banner file is empty")
		os.Exit(1)
	}

	myfile := string(file)
	var contents []string

	if filename == "thinkertoy.txt" {
		contents = strings.Split(myfile, "\r\n")
	} else {
		contents = strings.Split(myfile, "\n")
	}

	return contents
}

const hexErr = `Error: Incorrect color format
usage: --collor=#<color code (0-9 and A-F)>`

const rgbErr = `Error: Incorrect color format
usage: --collor=#<r,g,b (0 <= r, g & b >= 255)>`

const ansiErr = `Error: Incorrect color format
usage: --collor=ANSI Code (0 <= ANSI >= 255)>`

func ColorPicker(color string) (colorCode string) {
	colorChat := map[string]string{
		"reset":   "\u001b[39m",
		"red":     "\u001b[31m",
		"green":   "\u001b[32m",
		"magenta": "\u001b[35m",
	}

	colorCode, ok := colorChat[color]

	if ok {
		return colorCode
	}

	if ansi, err := strconv.Atoi(color); err == nil {
		if ansi < 0 || ansi > 255 {
			fmt.Println(ansiErr)
			os.Exit(1)
		}
		return fmt.Sprintf("\033[38;5;%dm", ansi)
	}

	r, g, b := -1, -1, -1

	if color[0] == '#' {
		if len(color) < 2 || len(color) > 7 {
			fmt.Println(hexErr)
			os.Exit(1)
		}

		hex := color

		if len(color) < 7 {
			for i := len(color); i < 7; i++ {
				hex += "0"
			}
		}

		for i := 1; i < 6; i += 2 {
			rgb, err := strconv.ParseInt(hex[i:i+2], 16, 32)
			if err == nil && rgb < 256 && rgb > -1 {
				switch i {
				case 1:
					r = int(rgb)
				case 3:
					g = int(rgb)
				case 5:
					b = int(rgb)
				}
			} else {
				fmt.Println(hexErr)
				os.Exit(1)
			}

		}
	}

	if strings.HasPrefix(strings.ToLower(color), "rgb") {
		if color[3] != '(' || color[len(color)-1] != ')' {
			fmt.Println(rgbErr)
			os.Exit(1)
		}
		rgbCode := strings.Split(strings.ReplaceAll(color[4:len(color)-1], " ", ""), ",")
		if len(rgbCode) != 3 {
			fmt.Println(rgbErr)
			os.Exit(1)
		}

		for i, v := range rgbCode {
			rgb, err := strconv.Atoi(v)
			if err == nil && rgb < 256 && rgb > -1 {
				switch i {
				case 0:
					r = rgb
				case 1:
					g = rgb
				case 2:
					b = rgb
				}
			} else {
				fmt.Println(rgbErr)
				os.Exit(1)	
			}
		}
	}

	if r == -1 || g == -1 || b == -1 {
		fmt.Println("Error: Color Not Found")
		os.Exit(1)
	}

	colorCode = fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)

	return colorCode
}

func getCi(substring, text string) (ci, n int) {
	arrSubSrting := strings.Split(substring, "\\n")
	word := ""

	for _, word = range arrSubSrting {
		ci = strings.Index(text, word)
		if ci != -1 {
			return ci, len(word)
		}
	}

	return ci, len(word)
}

// ProcessInput accepts the contents of the ASCII art file and the input string,
// and processes the input to display the corresponding ASCII art
func ProcessInput(contents []string, input, color, subString string) (strArt string) {
	count := 0
	strInput := strings.ReplaceAll(input, "\n", "\\n")
	strInput = strings.ReplaceAll(strInput, "\\t", "    ")
	newInput := strings.Split(strInput, "\\n")

	start := -1
	n := 0
	ci := 0

	for _, arg := range newInput {
		if arg == "" {
			count++
			if count < len(newInput) {
				strArt += "\n"
				continue
			} else {
				continue
			}
		}

		for i := 1; i <= 8; i++ {
			if subString != "" && subString != input {
				start, n = getCi(subString, arg)
			}

			if subString == input {
				start = 0
			}

			for j, ch := range arg {
				if ch > 126 {
					fmt.Println("The text contains an unprintable character", ch)
					os.Exit(0)
				}

				index := int(ch-32)*9 + i

				if index >= 0 && index < len(contents) {
					if start == j {
						strArt += color
					}

					strArt += (contents[index])

					if start != -1 && start+n-1 == j && j < len(arg)-1 && subString != "" && subString != input {
						strArt += ColorPicker("reset")
						ci, _ = getCi(subString, arg[j+1:]) // strings.Index(arg[j+1:], subString) + j + 1
						start = ci + j + 1
					}

				}
			}
			strArt += ColorPicker("reset") + "\n"
		}
	}

	return strArt
}
