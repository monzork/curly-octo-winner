package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	firstName string `json: "firstName"`
	lastName  string `json: "lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "48", Title: "Movie One", Director: &Director{firstName: "Jhon", lastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "4234", Title: "Movie Two", Director: &Director{firstName: "Paul", lastName: "Carusso"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
