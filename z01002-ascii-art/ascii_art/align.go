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

func Align(str, align string) (strAlign string) {
	if align == "left" {
		return str
	}

	wid := getTerminalSize()

	arrStr := strings.Split(str, "\n")
	format := fmt.Sprintf("%%%ds\n", wid + 10)
	//padding := strings.Repeat(" ", wid) //- len(arrStr[0]))

	//fmt.Println(wid)

	if align == "center" {
		wid /= 2
	}

	for _, stnc := range arrStr {
		// fmt.Println(i, stnc)
		strAlign += fmt.Sprintf(format, stnc)
		//fmt.Println(padding, stnc)
		// strAlign += format + stnc + "\n"
	}
	//fmt.Printf(format, "Hello")

	// fmt.Println(wid)

	// right := fmt.Sprintf("%%%ds\n", wid)
	// fmt.Printf(right, str)

	// centWid := wid / 2
	// cent := fmt.Sprintf("%%%ds\n", centWid)
	// fmt.Printf(cent, str)

	// arrStr := strings.Fields(str)
	// space := (int(wid) - len(arrStr[0]) + 1)/(len(arrStr) -1)
	// just := fmt.Sprintf("%%%ds", 0)
	// // fmt.Println(int(wid))
	// // fmt.Println(len(arrStr)-1)
	// // fmt.Println(space)

	// for _, word := range arrStr {
	// 	strJust += fmt.Sprintf(just, word)
	// 	just = fmt.Sprintf("%%%ds", space)
	// }
	// println()

	return strAlign
}
