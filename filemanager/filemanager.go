package filemanager

import (
	"bufio"
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
		return
	}

	filePath := "files/" + *file

	err = CreateFile(filePath, *content)
	if err != nil {
		fmt.Println(err)
		return
	}

	color.Green("File create success!")
}

func Read(args []string) {
	fs := flag.NewFlagSet("read", flag.ExitOnError)
	file := fs.String("file", "file.txt", "File to read")

	err := fs.Parse(args)
	if err != nil {
		fmt.Printf("Error parsing flags: %v\n", err)
		return
	}

	filePath := "files/" + *file

	err = ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error read the file: %v\n", err)
	}
}

func Update(args []string) {
	fs := flag.NewFlagSet("update", flag.ExitOnError)
	file := fs.String("file", "file.txt", "File to update")
	append := fs.String("append", "Bye, World! ", "Content of the file")

	err := fs.Parse(args)
	if err != nil {
		fmt.Printf("Error parsing flags: %v\n", err)
		return
	}

	filePath := "files/" + *file

	err = UpdateFile(filePath, *append)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	color.Green("File update success!")
}

func List() {
	path := "files"

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if len(files) == 0 {
		fmt.Println("Dir empty")
		return
	}

	for _, file := range files {
		err := showFileDetail(&file)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	}
}

func showFileDetail(file *os.DirEntry) error {

	info, err := (*file).Info()
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Name:", (*file).Name())
	fmt.Printf("Size: %d KB\n", info.Size())
	fmt.Println("Update:", info.ModTime().Format("2006-01-02 03:04 PM"))

	return nil
}

func Delete(args []string) {
	fs := flag.NewFlagSet("update", flag.ExitOnError)
	file := fs.String("file", "file.txt", "File to update")

	err := fs.Parse(args)
	if err != nil {
		fmt.Printf("Error parsing flags: %v\n", err)
		return
	}

	filePath := "files/" + *file

	err = DeleteFile(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	color.Green("File remove success!")
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

func ReadFile(filePath string) error {
	if !FileExists(filePath) {
		return fmt.Errorf("error: the file doesn't exists and you cannot read it")
	}

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func UpdateFile(filePath, content string) error {
	if !FileExists(filePath) {
		return fmt.Errorf("error: the file doesn't exists and you cannot update it")
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write([]byte("\n" + content)); err != nil {
		return err
	}

	return nil
}

func DeleteFile(filePath string) error {
	if !FileExists(filePath) {
		return fmt.Errorf("error: the file doesn't exists and you cannot delete it")
	}

	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}
