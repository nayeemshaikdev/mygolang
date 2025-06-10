package main

import (
	"fmt"
	//"math/rand"
	"crypto/rand"
	"math/big"
)

func main()  {
	fmt.Println("Welcome to My Maths")

	//using math/rand
	//randomNum := rand.Intn(5)
	//fmt.Println(randomNum)

	//using crypto/rand

	myRandNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandNum)
}