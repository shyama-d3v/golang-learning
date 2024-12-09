package main

import "fmt"

func main() {
	const a = 50
	fmt.Println(a)

	// declearing group of constant
	const (
		retryLimit = 4
		httpMethod = "GET"
	)

	fmt.Println(retryLimit)
	fmt.Println(httpMethod)

	const n = "Sam"
	var name = n
	// %T : Displays the type of a value
	// %v : Displays the value of a variable in its default format
	fmt.Printf("type %T value %v", name, name)

}
