package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"io/ioutil"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func main() {
	fmt.Println("Rest API with Go")
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests() // call to router
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true) // Create a new instance mux router

	myRouter.HandleFunc("/", handle).Methods("GET")                       // setting router rule
	myRouter.HandleFunc("/article", allArticles).Methods("GET")           // endpoint for article list
	myRouter.HandleFunc("/article/{id}", articleById).Methods("GET")      // endpoint for article by id
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")     // endpoint for creating new article
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE") // endpoint for deleting article by id

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", myRouter)) // listen and serve on port 8080
}

func handle(w http.ResponseWriter, r *http.Request) {
	var response [1]string
	response[0] = "Welcome to Rest API with Go"

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
	fmt.Println("Endpoint Hit: homePage")
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(Articles)
}

func articleById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleArticle")
	params := mux.Vars(r) // get params
	for _, item := range Articles {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&Article{})
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewArticle")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article

	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteNewArticle")
	var response [1]string

	vars := mux.Vars(r)
	id := vars["id"]

	for index, item := range Articles {
		if item.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			response[0] = "Article Deleted"
			break
		}
	}

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
