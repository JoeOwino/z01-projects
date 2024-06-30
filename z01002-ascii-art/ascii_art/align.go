package ascii_art

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getTerminalSize() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	dim := strings.Fields(string(out))

	w, _ := strconv.Atoi(dim[1])

	return w
}

func Align(str, align string) string {
	if len(str) >= getTerminalSize() {
		return str
	}

	padding := getTerminalSize() - len(str)

	if align == "center" {
		padding = (padding / 2) //- (len(str)/2)
	}

	// format := fmt.Sprintf("%%%ds", padding)
	// strAlign += fmt.Sprintf(format, str)
	return strings.Repeat(" ", padding) + str
}

func Justify(txt string, bannerConts []string) (strJust string) {
	spaces := len(strings.Fields(strings.TrimSpace(txt))) - 1

	if spaces < 1 {
		return txt
	}

	txtLen := 0
	inSpace := 0
	strLead := ""
	strTrail := ""
	strTxt := ""
	isLeading := true
	for i, ch := range txt {
		if ch == ' ' {
			if isLeading {
				txtLen += len(bannerConts[int(ch-32)*9+1])
				strLead += " "
			} else {
				inSpace += len(bannerConts[int(ch-32)*9+1])
				strTrail += " "	
			}
		} else {
			if isLeading {
				txtLen += len(bannerConts[int(ch-32)*9+1])
				isLeading = false
			} else {
				txtLen += len(bannerConts[int(ch-32)*9+1]) + len(bannerConts[1])
				inSpace = 0
				strTrail = ""
			}
			if i != 0 && txt[i - 1] == ' ' && !isLeading {
				strTxt  += " "
			}
			strTxt += string(ch)
		}
	}

	txtLen += inSpace
	
	txtLen -= spaces * len(bannerConts[1])
	if txtLen >= getTerminalSize() {
		fmt.Println(txt)
		return txt
	}

	justSpaces := ((getTerminalSize() - txtLen) / spaces) / len(bannerConts[1])
	padding := strings.Repeat(" ", justSpaces)

	for _, ch := range strTxt {
		if ch == ' ' {
			strJust += padding
		} else {
			strJust += string(ch)
		}
	}

	strJust = strLead + strJust + strTrail

	return strJust
}
