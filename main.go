package main

import (
	"fmt"

	"github.com/google/uuid"
	"golang.design/x/clipboard"
)

const (
	colorReset = "\033[0m"
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
	return func(s string) {
		fmt.Println(colorGray + s + colorReset)
	}
}
