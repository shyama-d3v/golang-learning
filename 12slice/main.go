// In Go, a slice is a flexible and more powerful abstraction built
//
// on top of arrays. Unlike arrays, slices are dynamically-sized, making them the
//
// most commonly used data structure for handling collections of data in Go.
package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices in golang")

	var slice []int    // Declares a slice of integers
	fmt.Println(slice) // Output: []

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	length := len(slice1)

	fmt.Println("Length of the slice:", length)
	// Total number of elements the slice's underlying array can hold. use cap() method
	fmt.Println("Capacity of the slice:", cap(slice1))
	// how to add value in slices
	slice1 = append(slice1, 6, 7)
	fmt.Println(slice1)

	// get particular range of data in slice
	slice1 = append(slice1[1:3])
	fmt.Println(slice1)

}
