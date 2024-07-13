package ascii_art

import (
	"strings"
)

func bannerMap(bannerContent []string) map[string]rune {
	bMap := map[string]rune{}
	i := 1

	for ch := ' '; ch <= '~'; ch++ {
		bMap[strings.Join(bannerContent[i:i+8], "\n")] = ch
		i += 9
	}

	return bMap
}

func ReverseAscii() (strRev string) {
	strArt := GetFile("./output/output.txt")
	bContent := GetFile("./.banners/standard.txt")
	bMap := bannerMap(bContent)
	arrCh := make([]string, 8)

	for n := 0; n < len(strArt) - 1; n += 8 {
		for j := 0; j < len(strArt[n]); j++ {
			for i, ln := range strArt[n:n+8] {
				if ln != "" {
					arrCh[i] += string(ln[j])
				}
				if ch, exists := bMap[strings.Join(arrCh, "\n")]; exists {
					strRev += string(ch)
					arrCh = make([]string, 8)
				}
			}
		}
		if n < len(strArt) - 1 {
			strRev += "\n"
		}
	}

	return strRev
}
