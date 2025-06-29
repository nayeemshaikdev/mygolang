package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("This is Switch Case")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println(diceNumber)

	switch diceNumber {
		case 1:
			fmt.Println("Dice value is 1 and you can open")		
		case 2:
			fmt.Println("You can move 2 spot")		
		case 3:
			fmt.Println("You can move 3 spot")	
			fallthrough	//to execute next case
		case 4:
			fmt.Println("You can move 4 spot")		
		case 5:
			fmt.Println("You can move 5 spot")		
			fallthrough //to execute next case
		case 6:
			fmt.Println("You can move 6 spot")	
	}
}