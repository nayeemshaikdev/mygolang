package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our Pizza from 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if(err != nil){
		fmt.Println("Error occurred while processing")
	} else {
		fmt.Println("Rating given")
		numRating := numRating + 1
		fmt.Println(numRating)
	}
}