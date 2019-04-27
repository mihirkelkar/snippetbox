package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/snippet", showSnippet)
	r.HandleFunc("/snippet/create", createSnippet).Methods("POST")

	//Serving static assets like logos, images, css etc.
	staticHandler := http.FileServer(http.Dir("./ui/static/"))
	withoutHeader := http.StripPrefix("/static", staticHandler)
	r.PathPrefix("/static/").Handler(withoutHeader)

	err := http.ListenAndServe("localhost:4000", r)
	if err != nil {
		log.Fatal(err)
	}
}
