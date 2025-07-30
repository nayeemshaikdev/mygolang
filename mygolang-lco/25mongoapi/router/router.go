package router

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/nayeemshaikdev/mongoapi/controller"
)

func Router() *mux.Router{
	fmt.Println("This is Router file")

	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/DeleteAllMoview", controller.DeleteAllMovies).Methods("DELETE")

	return router
}