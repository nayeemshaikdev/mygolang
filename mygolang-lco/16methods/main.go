package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs")
	myDetails := myStruct{"Aydin", "3", "aydin@gmail.com", true}
	fmt.Println(myDetails)
	fmt.Printf("My Name is : %v , My Age is %v , My Email is %v  \n",myDetails.Name, myDetails.Age, myDetails.Email)
	fmt.Printf("My complete details are : %+v\n",myDetails)
	myDetails.GetMethod("rajiya@aydin.com")
	fmt.Printf("222My complete details are : %+v\n",myDetails)
}

type myStruct struct{
	Name string
	Age string
	Email string
	Status bool
}

func (s myStruct) GetMethod(eml string) {
	s.Email = eml
	fmt.Println("New Email is : ",s.Email)
}