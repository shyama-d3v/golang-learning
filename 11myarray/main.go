package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")
	var fruitList [3]string
	fruitList[0] = "apple"
	fruitList[2] = "banana"
	fmt.Println(fruitList)
	fmt.Println(fruitList[0], fruitList[1])
	fmt.Println(len(fruitList))

	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("arr[%d] = %d\n", i, numbers[i])
	}

	//coping the array
	a := [3]int{1, 2, 3}
	b := a
	fmt.Println(b)
}
