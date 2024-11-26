package main

import (
	"fmt"
	"os"

	"github.com/RudysAcosta/file-manager-ClI/filemanager"
	"github.com/fatih/color"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		fmt.Println("Please provide an argument")
	}

	switch arguments[0] {
	case "interactive":
		interactive()
	case "create":
		filemanager.Create(os.Args[2:])
	case "read":
		filemanager.Read(os.Args[2:])
	default:
		fmt.Println("Invalid option")
	}
}

func interactive() {
	color.Blue("Interactive")
}
