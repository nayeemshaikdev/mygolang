package main

import "fmt"

func main() {
	//fmt.Println("Welcome to Defer")
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Hello")
	myDefer()
}

func myDefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}