package main

import (
	"fmt"

	"github.com/muesli/termenv"

	"github.com/google/uuid"
	"golang.design/x/clipboard"
)

const (
	colorReset = "\033[0m"
	colorBlack = "\033[30m"
	colorGray  = "\033[38;2;128;128;128m"
)

func main() {
	printFn := getPrintFunction()

	id := uuid.NewString()
	fmt.Println(id)

	err := clipboard.Init()
	if err != nil {
		fmt.Println("Could not paste to clipboard. Error: ", err)
		return
	}

	clipboard.Write(clipboard.FmtText, []byte(id))
	printFn("UUID copied to clipboard")
}

func getPrintFunction() func(string) {
	r, g, b := getTerminalColor()
	br := getBrightness(r, g, b)

	if br < 3 {
		return func(s string) {
			fmt.Println(colorGray + s + colorReset)
		}
	} else {
		return func(s string) {
			fmt.Println(colorBlack + s + colorReset)
		}
	}
}

func getTerminalColor() (r int, g int, b int) {
	seq := termenv.BackgroundColor().Sequence(true)

	fmt.Sscanf(seq, "48;2;%d;%d;%dm", &r, &g, &b)
	return
}

func getBrightness(r, g, b int) float64 {
	return 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
}
