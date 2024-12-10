// In Go, a slice is a flexible and more powerful abstraction built
//
// on top of arrays. Unlike arrays, slices are dynamically-sized, making them the
//
// most commonly used data structure for handling collections of data in Go.
package main

import (
	"fmt"
	"sort"
)

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

	hightScores := make([]int, 4)

	hightScores[0] = 111

	hightScores[1] = 112

	hightScores[2] = 113

	hightScores[3] = 114
	// hightScores[4] = 115 error occure
	hightScores = append(hightScores, 116, 115)
	fmt.Println(hightScores)

	sort.Ints(hightScores)
	fmt.Println("after sorting", hightScores)
	fmt.Println("is highScore sorted:", sort.IntsAreSorted(hightScores))

	//how to remove value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "ruby"}
	fmt.Println("course are: ", courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
