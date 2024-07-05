package ascii_art

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func colorMap(text, subString string) map[int]string {
	mapText := map[int]string{}
	colorStart := strings.Index(text, subString)
	colorReset := len(text) - 1
	if colorStart != -1 {
		colorReset = colorStart + len(subString) - 1
	}
	
	for i, ch := range text {

		if colorStart == -1 {
			mapText[i] =string(ch) + "NA"
			continue
		}

		if i < colorStart || (i > colorStart && i < colorReset) {
			mapText[i] = string(ch) + "NA"
			continue
		}

		if i == colorStart {
			mapText[i] = string(ch) + "CS"
			colorReset = colorStart + len(subString) - 1
			continue
		}

		if i == colorReset && i < len(text) - 1 {
			mapText[i] = string(ch) + "CR"
			colorStart = strings.Index(text[i + 1:], subString) + i + 1
			continue
		}

		mapText[i] = string(ch) + "NA"

	}


	return mapText
}

const hexErr = `Error: Incorrect color format
usage: --collor=#<color code (0-9 and A-F)>`

const rgbErr = `Error: Incorrect color format
usage: --collor=<rgb(r,g,b) (0 <= r, g & b >= 255)>`

const ansiErr = `Error: Incorrect color format
usage: --collor=ANSI Code (0 <= ANSI >= 255)>`

func ColorPicker(color string) (colorCode string) {
	if color == "" {
		return ""
	}
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
			if len(arrSubSrting) > 1 {
				if ci == 0 || ci == len(text)-len(word) {
					return ci, len(word)
				}
			} else {
				return ci, len(word)
			}
		}
	}

	return ci, len(word)
}
