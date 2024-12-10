package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to handle file in golang")
	content := "This need to go into the file"

	file, err := os.Create("./newFile.txt")
	if err != nil {
		panic(err)
	}
	length, _ := io.WriteString(file, content)
	fmt.Println("length is:", length)
	defer file.Close()
	readFile("./newFile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is:\n", string(databyte))
}
