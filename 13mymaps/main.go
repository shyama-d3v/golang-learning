package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)
	// fmt.Println(languages)

	// add data in maps
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["GO"] = "Golang"
	languages["PY"] = "Python"
	fmt.Println("list of all the languages:", languages)
	fmt.Println("JS stands for :", languages["JS"])

	//delete data in map
	delete(languages, "RB")

	fmt.Println("list of all the languages:", languages)

	// looping

	for _, value := range languages {
		fmt.Printf("value of language are : %v \n", value)
	}

}
