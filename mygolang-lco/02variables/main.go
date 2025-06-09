package main

import "fmt"

func main()  {
	fmt.Println("This is Variable")
	var username string = "Aydin"
	fmt.Println(username)
	fmt.Printf("Variable of type %T \n", username)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable of type %T \n", anotherVariable) 
}