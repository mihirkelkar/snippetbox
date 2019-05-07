package main

import (
	"html/template"
	"net/http"
)

type View struct {
	Template *template.Template
}

func NewView(files ...string) (*View, error) {
	ts, err := template.ParseFiles(files...)
	if err != nil {
		return nil, err
	}
	return &View{Template: ts}, nil
}

func (v *View) Render(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")
	err := v.Template.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error : 500 internal error from the render", 500)
	}
}
