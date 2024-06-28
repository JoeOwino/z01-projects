package ascii_art

import (
	//"fmt"
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
	if align == "left" {
		return str
	}

	padding := getTerminalSize() - len(str)

	if align == "center" {
		padding = (padding / 2) + (len(str)/2) 
	}

	//format := fmt.Sprintf("%%%ds", padding)
	//strAlign += fmt.Sprintf(format, str)
	return strings.Repeat(" ", padding) + str

}

func justify() string {
	
	return ""
}