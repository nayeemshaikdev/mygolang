package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://jsonplaceholder.typicode.com/posts/1/?user=Aydin"

func main() {
	fmt.Println("Welcome to URLs")

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	partOfUrl := &url.URL{
		Scheme: "https",
		Host: "jsonplaceholder.typicode.com",
		Path: "/posts/1/",
		RawQuery: "user=Aydin",
	}
	fmt.Println("part of url is : ", partOfUrl.String())
}