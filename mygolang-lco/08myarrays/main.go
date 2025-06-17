package main

import "fmt"

func main()  {
	fmt.Println("Welcome to Arrays")

	var fruitList =  [4]string{"Apple", "Banana", "", "Peach"}

	fmt.Println("Fruit list is : ",fruitList)
	fmt.Println("Fruit list is : ",len(fruitList))
}