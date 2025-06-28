package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	fmt.Println("Welcome to Web Requests")

	response, err := http.Get(url)
	genericError(err)

	//fmt.Printf("Type of response is : %T \n",response)

	//fmt.Println(response)

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	genericError(err)

	fmt.Println(string(body))
}

func genericError(err error){
	if err!=nil {
		panic(err)
	}
}