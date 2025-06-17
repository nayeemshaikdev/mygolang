package main

import "fmt"

func main()  {
	fmt.Println("Welcome to Pointers")

	//var myNumber *int

	myNumber :=23

	fmt.Println("myNumber is : ",myNumber)
	fmt.Println("myNumber is : ",&myNumber)

	var ptr = &myNumber
	fmt.Println("ptr is : ",ptr)
	fmt.Println("*ptr is : ",*ptr)
	*ptr += 2
	fmt.Println("myNumber is : ",myNumber)
	
}