package main

import (
	"encoding/json"
	"fmt"
	//"io"
	//"net/http"
)

func main() {
	fmt.Println("Welcome to build json")

	//encodeJson()
	decodeJson()
}

type Course struct {
	Name string 	`json:"coursename"`
	Price int		
	Platform string	`json:"website"`
	Password string	`json:"-"`
	Tags []string	`json:"tags,omitempty"`
}

func encodeJson(){
	//url := "http://localhost:8000/post"
	lcoCoursed := []Course {
		{"React js", 1299, "lcodev", "abc123", []string {"web-dev", "javascript"}},
		{"Python", 1399, "lcodev", "bck123", []string {"full-stack", "back-end"}},
		{"Salesforce", 1899, "lcodev", "cde123", nil},
		{"Java", 1399, "lcodev", "def123", []string {"web-dev", "full-stack"}},
	}

	finalJson, err := json.MarshalIndent(lcoCoursed, "", "\t")

	genericError(err)

	fmt.Printf("final json is :  %s\n",finalJson)

	//response, err := http.Post(url, "application/json", finalJson)
	//genericError(err)
	//defer response.Body.Close()

	//content, err := io.ReadAll(response.Body)
	//genericError(err)

}

func decodeJson()  {
	var jsonFromTheWeb = (`
		{
			"coursename": "Python",
			"Price": 1399,
			"website": "lcodev",
			"tags": ["full-stack", "back-end"]
        }
	`)
	fmt.Println("json from web is : ",jsonFromTheWeb)

	validJson := json.Valid([]byte(jsonFromTheWeb))
	fmt.Println("is json valid : ",validJson)

	var contentPlaceholder Course

	if validJson {
		json.Unmarshal([]byte(jsonFromTheWeb), &contentPlaceholder)
	}
	fmt.Println("content Placeholder is : ",contentPlaceholder)

	var mapPlaceholder map[string]interface {}
	json.Unmarshal([]byte(jsonFromTheWeb), &mapPlaceholder)

	fmt.Println("map placeholder is : ",mapPlaceholder)

	for k, v := range mapPlaceholder {
		fmt.Printf("key is %v and value is %v and Type is %T\n",k, v, v)
	}



}

func genericError(err error)  {
	if err != nil {
		panic(err)
	}
}