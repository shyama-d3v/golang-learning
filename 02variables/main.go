package main

import "fmt"

const LoginToken = "fggg"

func main() {
	var username string = "Shyama"
	fmt.Println(username)
	fmt.Printf("variable of type: %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("variable of type: %T \n", isLoggedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable of type: %T \n", smallVal)

	var smallFloatValue float32 = 255.5766066
	fmt.Println(smallFloatValue)
	fmt.Printf("variable of type: %T \n", smallFloatValue)

	// default values and some alias
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable of type: %T \n", anotherVariable)

	var anotherStringVariable string
	fmt.Println("String is :", anotherStringVariable)
	fmt.Printf("variable of type: %T \n", anotherStringVariable)

	//implicit type
	var website = "learnCodeOnline"
	fmt.Println(website)

	// no var style
	numberOfUser := 30000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}
