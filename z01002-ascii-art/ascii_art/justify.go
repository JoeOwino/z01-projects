package ascii_art

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

type winsize struct {
    Row    uint16
    Col    uint16
    Xpixel uint16
    Ypixel uint16
}

func Justify(str string) (strJust string) {
    ws := &winsize{}
    retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(syscall.Stdin),
        uintptr(syscall.TIOCGWINSZ),
        uintptr(unsafe.Pointer(ws)))

    if int(retCode) == -1 {
        panic(errno)
    }

    wid := (uint(ws.Col))
	// right := fmt.Sprintf("%%%ds\n", wid)
	// fmt.Printf(right, str)

	// centWid := wid/2
	// cent := fmt.Sprintf("%%%ds\n", centWid)
	// fmt.Printf(cent, str)

	arrStr := strings.Fields(str)
	space := (int(wid) - len(arrStr[0]) + 1)/(len(arrStr) -1)
	just := fmt.Sprintf("%%%ds", 0)
	// fmt.Println(int(wid))
	// fmt.Println(len(arrStr)-1)
	// fmt.Println(space)

	for _, word := range arrStr {
		strJust += fmt.Sprintf(just, word)
		just = fmt.Sprintf("%%%ds", space)
	}
	//println()

	return strJust

}