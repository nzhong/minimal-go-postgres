package main

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()

	// Declare static file directory
	staticFileDirectory := http.Dir("./static/")
	staticFileServer := http.FileServer(staticFileDirectory)
	// Create file handler. Although the static files are placed inside `./static/` folder
	// in our local directory, it is served at the root (i.e., http://localhost:8080/)
	// when browsed in a browser. Hence, we need `http.StripPrefix` function to change the
	// file serve path.
	staticFileHandler := http.StripPrefix("/", staticFileServer)
	r.Handle("/", staticFileHandler).Methods("GET")

	return r
}

func main() {
	connString := `user=postgres 
				   password=welcome
				   host=localhost
				   port=5432
				   dbname=hello 
				   sslmode=disable`
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	store = &dbStore{db: db}

	r := newRouter()
	http.ListenAndServe(":8080", r)
}
