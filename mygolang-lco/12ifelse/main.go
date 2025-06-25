package main

import "fmt"

func main()  {
	fmt.Println("This is If Else")

	result := 23

	if result < 10 {
		fmt.Println("Irregular user")
	} else if result > 10 {
		fmt.Println("Frequent user")
	} else {
		fmt.Print("user loggedIn 10 times")
	}

	if 9%2 == 0 {
		fmt.Println("Numbe is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 9; num >= 10 {
		fmt.Println("greater")
	} else {
		fmt.Println("lesser")
	}
}