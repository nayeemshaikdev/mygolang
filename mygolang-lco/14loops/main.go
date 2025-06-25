package main

import "fmt"

func main()  {
	fmt.Println("Welcome to loops")

	//days := []string {"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	//fmt.Println(days)

	// for i:=0; i < len(days); i++{
	// 	fmt.Println(days[i])
	// }

	//for i := range days {
	//	fmt.Println(days[i])
	//}

	//for index, day := range days {
	//	fmt.Printf("index is %v and day is %v\n", index, day)
	//}

	//for _, day := range days {
	//	fmt.Printf("index is _ and day is %v\n",day)
	//}

	rougeValue :=1

	for rougeValue < 10{
		if rougeValue ==7 {
			//rougeValue++
			goto lco
		}
		fmt.Println("Value is : ",rougeValue)
		rougeValue++
		fmt.Println("Value is 2: ",rougeValue)
	}
	lco:
		gotoCheck()
}                                                      



func gotoCheck()  {
	fmt.Println("from gotoCheck function")
}