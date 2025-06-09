package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok || err err

	rating, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating", rating)
	fmt.Printf("Type of this input %T \n",rating)
	fmt.Println("error occurred for rating", err)
}