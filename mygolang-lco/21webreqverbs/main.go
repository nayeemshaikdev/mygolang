package main

import (
	"fmt"
	"strings"
	"io"
	"net/http"
	"net/url"
)

const globalUrl string = "http://localhost:8000/get"

func main(){
	fmt.Println("welcome to Web Request Verbs")
	//myRequest()
	//PerformPostJsonRequest()
	PerformPostformRequest()
}

func myRequest()  {
	
	response, err := http.Get(globalUrl)

	genericError(err)

	defer response.Body.Close()

	fmt.Println("Status code is : ", response.StatusCode)
	fmt.Println("content length is : ", response.ContentLength)

	content, err := io.ReadAll(response.Body)
	genericError(err)

	fmt.Println("Content is : ", string(content))

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

func PerformPostJsonRequest()  {
	const myUrl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"name": "Aydin",
			"DOB": "050722"
		}
	`)
	//fmt.Println("request body is : ",requestBody)
	response, err := http.Post(myUrl, "application/json", requestBody)
	genericError(err)	
	defer response.Body.Close()
	var responseString strings.Builder
	content, err := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("byteCount is : ",byteCount)
	genericError(err)
	fmt.Println(responseString.String())

	//fmt.Println(string(content))
}

func PerformPostformRequest(){
	const myUrl = "http://localhost:8000/postform"

	data := url.Values{}

	data.Add("firstname", "Aydin")
	data.Add("lastname", "Shaik")
	data.Add("email", "aydinmohammed@aydin.com")

	response, err := http.PostForm(myUrl, data)

	genericError(err)

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	var responseString strings.Builder

	_, err = responseString.Write(content)

	fmt.Println("content is : ", responseString.String())

}