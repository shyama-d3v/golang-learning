package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var filename string
	data := make([]byte, 100)
	fmt.Println("Enter the file name")
	fmt.Scanln(&filename)
	if len(filename) == 0 {
		log.Fatal("Please provide file Name")
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open the file: %v", err)
	}
	defer func() {
		file.Close()
		fmt.Println("file handle is closed.")
	}()

	count, _ := file.Read(data)
	fmt.Println(string(data[:count]))

}
