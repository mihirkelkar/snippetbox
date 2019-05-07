package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//Controller struct for HomeController
type HomeController struct {
	HomeView    *View //View is a custom class containg a template that we defined
	Application *Application
}

//Function that lets you create a new homeController
func NewHomeController(app *Application) *HomeController {
	homeStr := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl"}
	homeView, err := NewView(homeStr...)
	if err != nil {
		app.ErrorLog.Println(err.Error())
	}
	return &HomeController{HomeView: homeView, Application: app}
}

func (h *HomeController) home(w http.ResponseWriter, r *http.Request) {
	h.HomeView.Render(w)
	return
}

type SnippetController struct {
	SnippetView *View
	app         *Application
}

//NewSnippetController : Controller handler for snippets
func NewSnippetController(app *Application) *SnippetController {
	return &SnippetController{SnippetView: nil, app: app}
}

func (s *SnippetController) showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		s.app.notFound(w)
		return
	}
	snippet, err := s.app.ModelService.Snippet.Get(id)
	if err != nil {
		s.app.notFound(w)
		return
	}
	fmt.Fprintf(w, "%v", snippet)
}

func (s *SnippetController) createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"Create" :"snippets here"}`))
}
