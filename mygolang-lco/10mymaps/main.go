package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Maps in Go")
	
	language := make(map[string]string)
	language["JS"] = "Javascript"
	language["PY"] = "Python"
	language["RB"] = "Ruby"
	fmt.Println("Languages are : ",language)
	fmt.Println("JS short is : ", language["JS"])
	fmt.Println("RB short is : ", language["RB"])
	fmt.Println("PY short is : ", language["PY"])

	delete(language, "RB")
	fmt.Println("Languages are : ",language)

	for key, value := range language{
		fmt.Printf("key is %v and value is %v\n", key, value)
	}
}