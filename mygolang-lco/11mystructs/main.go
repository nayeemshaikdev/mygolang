package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs")
	myDetails := myStruct{"Aydin", "3", "aydin@gmail.com"}
	fmt.Println(myDetails)
	fmt.Printf("My Name is : %v , My Age is %v , My Email is %v  \n",myDetails.Name, myDetails.Age, myDetails.Email)
	fmt.Printf("My complete details are : %+v",myDetails)
}

type myStruct struct{
	Name string
	Age string
	Email string
}