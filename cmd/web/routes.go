package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *Application) routes() *mux.Router {

	hmCtrl := NewHomeController(app)
	spCtrl := NewSnippetController(app)

	r := mux.NewRouter()
	r.HandleFunc("/", hmCtrl.home).Methods("GET")
	r.HandleFunc("/snippet", spCtrl.showSnippet)
	r.HandleFunc("/snippet/create", spCtrl.createSnippet).Methods("POST")

	//Serving static assets like logos, images, css etc.
	staticHandler := http.FileServer(http.Dir("./ui/static/"))
	withoutHeader := http.StripPrefix("/static", staticHandler)
	r.PathPrefix("/static/").Handler(withoutHeader)
	return r
}
