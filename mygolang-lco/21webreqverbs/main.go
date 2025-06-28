package main

import (
	"fmt"
	"strings"
	"io"
	"net/http"
)

const url string = "http://localhost:8000/get"

func main(){
	fmt.Println("welcome to Web Request Verbs")
	myRequest()
}

func myRequest()  {
	
	response, err := http.Get(url)

	genericError(err)

	defer response.Body.Close()

	fmt.Println("Status code is : ", response.StatusCode)
	fmt.Println("content length is : ", response.ContentLength)

	content, err := io.ReadAll(response.Body)
	genericError(err)

	//fmt.Println("Content is : ", string(content))

	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is : ",byteCount)
	fmt.Println(responseString.String())


}

func genericError(err error)  {
	if err != nil{
		panic(err)
	}
}