package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mihirkelkar/snippetbox/pkg/models"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP Network Port")
	dsn := flag.String("dsn", "web:test@tcp(db:3306)/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	fmt.Println(*dsn)

	//Set Loggers for Info and Error
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//Lets define an http server object instead of sticking to the normal.
	//listen and serve routine. This way, we can make our server object also
	//log errors to our standard error logger.

	db, err := OpenDB(*dsn)
	if err != nil {
		errLog.Fatal(err.Error())
	}
	defer db.Close()

	newModelService := models.NewModelService(db)

	app := Application{ErrorLog: errLog, InfoLog: infoLog, ModelService: newModelService}

	srv := http.Server{
		Addr:     *addr,
		Handler:  app.routes(), //This sets up all the end points and paths
		ErrorLog: errLog,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errLog.Fatal(err)
}

//OpenDB : opens a connection to the database.

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
