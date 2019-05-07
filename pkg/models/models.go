package models

import (
	"database/sql"

	"github.com/mihirkelkar/snippetbox/pkg/mysql"
)

type ModelService struct {
	Snippet mysql.SnippetDB
}

//Return an object of the several types of models we can have and
//provides a common interface of interacting with them.
func NewModelService(db *sql.DB) *ModelService {
	return &ModelService{
		Snippet: mysql.NewSnippetDB(db),
	}
}
