package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza App!")
	fmt.Println("Please rate our pizza between 1 and 5.")

	// Create a buffered reader to read user input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your rating: ")
	input, err := reader.ReadString('\n') // Read input till the newline character
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Trim whitespace and parse the input as a float
	rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		// fmt.Println(err)
		fmt.Println("Invalid input. Please enter a numeric value between 1 and 5.")
		return
	}

	// Validate the rating range
	if rating < 1 || rating > 5 {
		fmt.Println("Rating out of range! Please rate between 1 and 5.")
		return
	}

	// Provide feedback
	fmt.Printf("Thanks for rating our pizza: %.1f\n", rating)
	fmt.Printf("Here’s your rating incremented by 1: %.1f\n", rating+1)
}
