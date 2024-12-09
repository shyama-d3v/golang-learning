package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")
	presentTime := time.Now()
	fmt.Println("Current time:", presentTime)
	fmt.Println("Formatted time:", presentTime.Format("01-02-2006 15:04:05"))

	future := time.Now().Add(2 * time.Hour)   // Add 2 hours
	past := time.Now().Add(-30 * time.Minute) // Subtract 30 minutes
	fmt.Println("2 hours later:", future)
	fmt.Println("30 minutes earlier:", past)

	fmt.Println("Year:", presentTime.Year())
	fmt.Println("Month:", presentTime.Month())
	fmt.Println("Day:", presentTime.Day())
	fmt.Println("Hour:", presentTime.Hour())
	fmt.Println("Minute:", presentTime.Minute())
	fmt.Println("Second:", presentTime.Second())

	cratedDate := time.Date(2024, time.August, 10, 20, 23, 0, 0, time.UTC)
	fmt.Println(cratedDate)
}
