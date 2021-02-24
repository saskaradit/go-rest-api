package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Description"`
	Content string `json:"Content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Title 1", Desc: "First", Content: "Hello World"},
	}
	fmt.Println("Endpoint Hit: All articles is here")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST article")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page Hello World")
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", allArticles).Methods(("GET"))
	router.HandleFunc("/articles", testPostArticles).Methods(("POST"))
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequest()
}
