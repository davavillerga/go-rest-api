package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Movie struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Movies []Movie

func allMovies(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{Title: "The Lord of The Rings: The Return of The King", Desc: "The 3rd part of the LOTR trilogy", Content: "https://storage.googleapis.com/LOTR-ROTK.video"},
		Movie{Title: "Reservoir Dogs", Desc: "A violent masterpiece by Quentin Tarantino", Content: "https://storage.googleapis.com/Reservoir-Dogs.video"},
		Movie{Title: "Matrix Revolutions", Desc: "The 3rd part of the Matrix trilogy", Content: "https://storage.googleapis.com/Matrix-revolutions.video"},
	}
	fmt.Println("Endpoint Hit: All Movies Endpoint")
	json.NewEncoder(w).Encode(movies)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/movies", allMovies)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
