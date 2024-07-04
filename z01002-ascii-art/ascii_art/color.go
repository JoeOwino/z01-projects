package ascii_art

import (
	"strings"
)

func colorMap(text, subString string) map[int]string {
	mapText := map[int]string{}
	colorStart := strings.Index(text, subString)

	for i, ch := range text {
		if colorStart == -1 {
			mapText[i] =string(ch) + "NA"
			continue
		}

		colorRest := colorStart + len(subString) - 1

		if i < colorStart || (i > colorStart && i < colorRest) {
			mapText[i] = string(ch) + "NA"
			continue
		}

		if i == colorStart {
			mapText[i] = string(ch) + "CS"
			continue
		}

		if i == colorRest {
			mapText[i] = string(ch) + "CR"
			if i < len(text) - 1 {
				colorStart = strings.Index(text[i + 1:], subString)
			}

		}

	}

	return mapText
}
