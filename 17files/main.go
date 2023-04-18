package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("files")
	const filepath string = "./file.txt"
	file, err := os.Create(filepath)
	checkNilErr(err)
	writeToFile("Hello from main", file)
	readFromFile(filepath)
}

func writeToFile(contents string, file *os.File) {
	fmt.Println("contents to be written", contents)

	_, err := io.WriteString(file, contents)
	checkNilErr(err)
	file.Close()
}

func readFromFile(path string) {
	contents, err := os.ReadFile(path)
	checkNilErr(err)
	fmt.Println("contents of file", string(contents))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
