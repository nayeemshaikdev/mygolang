package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	//"time"
)

var wg sync.WaitGroup

func main(){
	fmt.Println("Welcome to Goroutines")
	websiteList := [] string {
		"https://google.com",
		"https://fb.com",
		"https://go.dev",
		"https://github.com",
	}

	for _, web := range websiteList {
		//time.Sleep(1000 * time.Millisecond)
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}

func getStatusCode(endpoint string)  {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d is the status code for website %s\n",res.StatusCode, endpoint)
}