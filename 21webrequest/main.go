package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://dogapi.dog/api/v2/breeds"

func main() {
	fmt.Println("Welcome to web request tutorial")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("type of the response %T\n", response)
	defer response.Body.Close()

	databyte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println(content)

}
