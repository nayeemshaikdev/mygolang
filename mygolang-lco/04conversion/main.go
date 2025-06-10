package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our Pizza from 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if(err != nil){
		fmt.Println("Error occurred while processing")
	} else {
		fmt.Println("Rating given")
		//input := input + 1
		fmt.Println(input)
	}
}