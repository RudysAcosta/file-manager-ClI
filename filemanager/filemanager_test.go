package filemanager

import (
	"flag"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	filePath := "testfile.txt"
	f, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	f.Close()
	defer os.Remove(filePath)

	// Test FileExists
	if !FileExists(filePath) {
		t.Errorf("Expected file %s to exist", filePath)
	}
}

func TestCreateFile(t *testing.T) {
	os.Args = []string{"cmd", "--file=testfile.txt", "--content=Hello, Test!"}

	file := flag.String("file", "file.txt", "File to create")
	content := flag.String("content", "Hello, Test!", "Content of the file")
	flag.Parse()

	setupTestDirectory()

	filePath := "files/" + *file

	// delete the file when de function finish
	defer os.Remove(filePath)

	CreateFile(filePath, *content)

	// Verificar que el archivo fue creado
	if !FileExists(filePath) {
		t.Errorf("expected file to be created: %s", *file)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read the created file: %v", err)
	}

	if string(data) != *content {
		t.Errorf("expected file content '%s', got '%s'", *content, string(data))
	}
}

func TestReadFile(t *testing.T) {
	os.Args = []string{"cmd", "--file=testfile.txt"}

	file := flag.String("file", "file.txt", "File to read")
	flag.Parse()

	setupTestDirectory()

	filePath := "files/" + *file

	// delete the file when de function finish
	defer os.Remove(filePath)

	content := "Hello, Test!"
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	ReadFile(filePath)
}

func TestUpdateFile(t *testing.T) {
	os.Args = []string{"cmd", "--file=testfile.txt", "--append=Bye, World!"}

	file := flag.String("file", "file.txt", "File to update")
	append := flag.String("append", "Bye, World!", "Content of the file")
	flag.Parse()

	setupTestDirectory()

	filePath := "files/" + *file

	// delete the file when de function finish
	defer os.Remove(filePath)

	content := "Hello, Test!"
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	UpdateFile(filePath, *append)

	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read the created file: %v", err)
	}

	expected := content + *append
	if string(data) != expected {
		t.Errorf("expected file content '%s', got '%s'", expected, string(data))
	}
}

func setupTestDirectory() {
	if _, err := os.Stat("files"); os.IsNotExist(err) {
		os.Mkdir("files", 0755)
	}
}
