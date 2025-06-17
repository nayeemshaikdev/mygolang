package main

import(
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Welcome to my time")

	presentTime :=time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2022, time.July, 05, 5, 30, 0, 0, time.UTC)

	fmt.Println("Aydin Birthday : ", createdDate)
	fmt.Println("Aydin Birthday : ", createdDate.Format("01-02-2006 Monday"))
}o