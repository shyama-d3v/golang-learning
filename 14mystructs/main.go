package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("struct in golang")
	shyama := User{"Shyama", "shyama.patro@gmail.com", true, 27}

	fmt.Printf("Shyama's profile details is:%+v\n", shyama)
	fmt.Printf("Name is %v and Email is %v ", shyama.Name, shyama.Name)
}
