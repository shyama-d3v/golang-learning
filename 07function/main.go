package main

import "fmt"

// Function With Return Value
// func addNumbers(a int, b int) int {
// 	return a + b
// }

// func main() {

// 	var input1, input2 string

// 	fmt.Println("Enter the first number: ")
// 	fmt.Scanln(&input1)

// 	fmt.Println("Enter the second number: ")
// 	fmt.Scanln(&input2)

// 	num1, err1 := strconv.Atoi(input1)
// 	num2, err2 := strconv.Atoi(input2)

// 	if err1 != nil || err2 != nil {
// 		fmt.Println("Invalid input. Please enter valid integers.")
// 		return
// 	}

// 	result := addNumbers(num1, num2)

// 	fmt.Println("The sum is:", result)

// }

// ------------------------------------------------------------
// Function Without Parameters and Return Values

// func greet() {
// 	fmt.Println("Hello, welcome to Go!")
// }
// func main() {
// 	greet()
// }

// ----------------------------------------------------------------------------------
// Define and call an anonymous function
// func main() {
// 	func() {
// 		fmt.Println("This is an anonymous function!")
// 	}()
// }

// ----------------------------------------------------------------------------------
// Function with parameters

func greetUser(name string, age int) {
	fmt.Println("Hello,", name)
	fmt.Printf("You are %d years old.\n", age)
}

func main() {

	greetUser("Alice", 25)
	greetUser("Bob", 30)
}

// ----------------------------------------------------------------------------------
