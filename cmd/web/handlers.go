package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error: Internal Service Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error: Internal Service Error", 500)
		return
	}

}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.Error(w, "Error : This ID is invalid", 404)
		return
	}
	fmt.Fprintf(w, "{\"Showing\": \"Snippet Information for snippet %d!\"}", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"Create" :"snippets here"}`))
}
