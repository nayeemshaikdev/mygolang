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

	input, _ := reader.ReadString('\n')
	fmt.Println("Thankd for rating", input)
	fmt.Printf("Type of this input %T",input)
}