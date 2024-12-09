package main

import (
	"fmt"
	"time"
)

func main() {
	// if statements --- example 1
	num := 10
	if num < 15 {
		fmt.Println("Number is less than 15")
	}

	// if statements --- example 2
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// if...else statements  example 1
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// if...else statements  example 2
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// switch...case statements
	today := time.Now()
	// fmt.Println("................", today)
	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
	case 9:
		fmt.Println("Today is 9th. Buy some wine.")
	case 15:
		fmt.Println("Today is 15th. Visit a doctor.")
	case 25:
		fmt.Println("Today is 25th. Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}

}
