package main

import (
	//"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrice int    `json:"price"`
	Author      *Author
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to Build API")
	r := mux.NewRouter()

	//Seeding
	courses = append(courses, Course{CourseId:"1", CourseName:"React JS", CoursePrice:199, Author: &Author{FullName: "Aydin", Website: "aydin.com"}})

	courses = append(courses, Course{CourseId:"2", CourseName:"Python", CoursePrice:299, Author: &Author{FullName: "Mohammed", Website: "aydin.com"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request){
	
	w.Write([]byte("<h1>Welcome to API by Aydin</h1>"))
}


func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all Courses")

	w.Header().Set("content-type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get One Course based on Id from Request")

	w.Header().Set("content-type", "application/json")

	param := mux.Vars(r)

	for _, course := range  courses {
		if(course.CourseId == param["id"]){
			json.NewEncoder(w).Encode(course)
			return
		}
		
	}
	json.NewEncoder(w).Encode("No courses found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request)  {
	/* Request Payload
	{
		"coursename": "Salesforce",
		"price": 399,
		"Author": {
			"fullname": "Aydin",
			"website": "aydin.com"
		}
	}
	*/
	fmt.Println("Create one course")
	w.Header().Set("content-type", "application/json")
	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	// what about - {}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send valid data")
		return
	}

	// generate unique Id , string
	// append the data to slick i.e., append course to Courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	for _, eachCourse := range courses {
		if eachCourse.CourseId == course.CourseId || eachCourse.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course Id or Course Name already exists")
			return
		}
	}

	courses = append(courses, course)

	json.NewEncoder(w).Encode(courses)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Update one course")

	w.Header().Set("conteny-type", "application/json")
	//Deserialize request into a variable "params" using Mux
	params := mux.Vars(r)

	//var course Course

	//_ = json.NewDecoder(r.Body).Decode(&course)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//TODO: send a response when id is not found
	json.NewEncoder(w).Encode("Id not found in JSON")
	return

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Delete one course")
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]... )
			json.NewEncoder(w).Encode(courses)
			break
		}
	}
	//TODO: send a response when id is not found
	//json.NewEncoder(w).Encode("Id not found in JSON")
	return
}