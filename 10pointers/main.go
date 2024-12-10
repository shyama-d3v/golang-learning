package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class of pointers ")

	//creating the pointer
	var ptr *int
	fmt.Println("value of the pointer is :", ptr)

	//referencing the pointer
	myNumber := 23
	var pointer = &myNumber
	fmt.Println("address of the pointer:", pointer)
	fmt.Println("value of the actual pointer:", *pointer)

	*pointer = *pointer + 2
	fmt.Println("New value is:", myNumber)

}
