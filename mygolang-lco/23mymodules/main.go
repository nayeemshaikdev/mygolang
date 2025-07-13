package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Mod in golang")

	r := mux.NewRouter()
    r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter(){
	fmt.Println("This is greeter")
}

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Shaik Aydin Mohammed</h1>"))
	
}
