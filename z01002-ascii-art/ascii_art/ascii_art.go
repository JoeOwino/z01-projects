package ascii_art

import (
	"fmt"
	"os"
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

	if strings.Contains(myfile, "\r") {
		contents = strings.Split(myfile, "\r\n")
	} else {
		contents = strings.Split(myfile, "\n")
	}

	return contents
}

// ProcessInput accepts the contents of the ASCII art file and the input string,
// and processes the input to display the corresponding ASCII art
func ProcessInput(contents []string, input, color, subString, align string) (strArt string) {
	count := 0
	strInput := strings.ReplaceAll(input, "\n", "\\n")
	subString = strings.ReplaceAll(subString, "\n", "\\n")
	strInput = strings.ReplaceAll(strInput, "\\t", "    ")
	subString = strings.ReplaceAll(subString, "\\t", "    ")

	coloMap := colorMap(strInput, subString)
	newInput := strings.Split(strInput, "\\n")
	ci := 0

	for n, arg := range newInput {
		ci += n * 2
		if arg == "" {
			count++
			if count < len(newInput) {
				strArt += "\n"
			}
			ci += 2
			continue
		}

		for i := 1; i <= 8; i++ {
			strLine := ""
			lenColor := 0
			isLeading := true
			traillingSpace := ""

			if subString != "" && subString != input && strings.Contains(input, subString) {
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
						strLine += color
						lenColor += len(color)
					}

					if align != "justify" {
						strLine += (contents[index])
					} else if ch == ' ' {
						if isLeading {
							strLine += (contents[index])
						} else {
							if j < len(arg)-1 && arg[j+1] != ' ' {
								strLine += "j"
								traillingSpace = ""
							} else {
								traillingSpace += (contents[index])
							}
						}
					} else if ch != ' ' {
						strLine += (contents[index])
						isLeading = false
					}

					if strings.Contains(input, subString) && start+n-1 == j && j < len(arg)-1 && subString != "" && subString != input {
						strLine += ColorPicker("reset")
						lenColor += len(ColorPicker("reset"))
						ci, _ = getCi(subString, arg[j+1:])
						start = ci + j + 1
					}

				}
			}
			strLine += traillingSpace
			if color != "" {
				strLine += ColorPicker("reset")
				lenColor += len(ColorPicker("reset"))
			}
			if align == "right" || align == "center" || align == "justify" {
				strLine = Align(strLine, align, lenColor)
			}
			strArt += strLine + "\n"
		}
	}

	return strArt
}
