package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.design/x/clipboard"
)

const (
	colorReset = "\033[0m"
	colorGray  = "\033[38;2;128;128;128m"
)

func main() {
	printFn := getPrintFunction()

	var id string

	if len(os.Args) > 1 && os.Args[1] == "mongo" {
		id = primitive.NewObjectID().Hex()
	} else {
		id = uuid.NewString()
	}

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
