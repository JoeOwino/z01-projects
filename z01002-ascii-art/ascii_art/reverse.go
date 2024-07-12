package ascii_art

import (
	"fmt"
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

func PrintMap() {
	bContent := GetFile("./.banners/standard.txt")
	//strContent := strings.Join(bContent,"\n")
	bMap := bannerMap(bContent)

	for key, val := range bMap {
		fmt.Println(string(val))
		fmt.Println("-------------")
		for _, ch := range key {
			fmt.Print(string(ch))
		}
	}
}