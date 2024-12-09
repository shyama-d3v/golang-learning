package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wecome := "welcome to user input"
	fmt.Println(wecome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")
	//comma ok || comma ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("variable of type: %T \n", input)
}
