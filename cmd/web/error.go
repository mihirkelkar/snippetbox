package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/mihirkelkar/snippetbox/pkg/models"
)

type Application struct {
	ErrorLog     *log.Logger
	InfoLog      *log.Logger
	ModelService *models.ModelService
}

func (app *Application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s/n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
