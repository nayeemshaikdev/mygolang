package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to slices")

	var fruitList = [] string{"Apple", "Tomato", "Peach"}

	fmt.Println(fruitList)

	fruitList = append(fruitList, "Litchi")

	fmt.Println(fruitList)

	//fruitList = append(fruitList[:1])
	highScores := [] int{23, 43, 11, 78, 9, 5, 99}
	sort.Ints(highScores)
	fmt.Println(highScores)

	courses := []string {"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	index := 3

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}