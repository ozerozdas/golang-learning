// Firstly, run: go get -u github.com/gorilla/mux
package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux" // import Gorilla mux
)

// Required page parameters
type page struct {
	Title string
}

func main() {
	router := mux.NewRouter()

	// Route Management
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/{title}", list).Methods("GET")

	http.ListenAndServe("127.0.0.1:8080", router) // localhost:8080
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("layout.html")) // Template definition
	data := page{Title: "My TODO list"}                       // Data to be passed to the template
	tmpl.Execute(w, data)                                     // Execute the template
}

func list(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)    // Get the variables from the request
	title := vars["title"] // Get the title from the URL

	tmpl := template.Must(template.ParseFiles("layout.html")) // Template definition
	data := page{Title: title}                                // Data to be passed to the template
	tmpl.Execute(w, data)                                     // Execute the template
}
