package filemanager

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}

	if !os.IsNotExist(err) {
		log.Fatalf("Error checking the file: %v", err)
	}

	return false
}

func Create(args []string) {
	fs := flag.NewFlagSet("create", flag.ExitOnError)
	file := fs.String("file", "file.txt", "File to create")
	content := fs.String("content", "Hello, World!", "Content of the file")

	err := fs.Parse(args)
	if err != nil {
		fmt.Printf("Error parsing flags: %v\n", err)
	}

	filePath := "files/" + *file

	err = CreateFile(filePath, *content)
	if err != nil {
		fmt.Println(err)
		return
	}

	color.Green("File create success!")
}

func CreateFile(filePath, content string) error {

	if FileExists(filePath) {
		return fmt.Errorf("error: the file already exists and you cannot create it")
	}

	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	_, err = fmt.Fprint(f, content)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	return nil
}
