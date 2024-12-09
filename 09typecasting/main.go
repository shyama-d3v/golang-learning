package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 42
	fmt.Printf("before typecasting, variable of type: %T \n", a)
	f := float64(a)
	fmt.Println(f) // 42
	fmt.Printf("after typecasting, variable of type: %T \n", f)

	// Conversions between string and int
	var s string = "44"
	fmt.Printf("before typecasting, variable of type: %T \n", s)

	v, _ := strconv.Atoi(s) // convert string to int

	fmt.Println(v) // 44
	fmt.Printf("after typecasting, variable of type: %T \n", v)
	var i int = 44
	fmt.Printf("before typecasting, variable of type: %T \n", i)

	str := strconv.Itoa(i) // convert int to string

	fmt.Println(str) // 44
	fmt.Printf("after typecasting, variable of type: %T \n", str)

}
