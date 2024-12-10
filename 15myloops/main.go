package main

import "fmt"

func main() {
	// basic for loop syntax
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value of i:", i)
	// }

	// while like syntax
	// count := 0
	// for count < 5 {
	// 	fmt.Println("Count is :", count)
	// 	count++
	// }

	// use range for  iterate slice
	// numbers := []int{10, 20, 30, 40}
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }

	// use range for Loop Over a String
	// text := "hello"
	// for index, char := range text {
	// 	fmt.Printf("Index: %d, Character: %c\n", index, char)
	// }

	// use break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}
