package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP Network Port")
	flag.Parse()

	//Set Loggers for Info and Error
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/snippet", showSnippet)
	r.HandleFunc("/snippet/create", createSnippet).Methods("POST")

	//Serving static assets like logos, images, css etc.
	staticHandler := http.FileServer(http.Dir("./ui/static/"))
	withoutHeader := http.StripPrefix("/static", staticHandler)
	r.PathPrefix("/static/").Handler(withoutHeader)
	infoLog.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, r)
	if err != nil {
		errLog.Fatal(err.Error())
	}
}
