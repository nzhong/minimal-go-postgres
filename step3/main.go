package main

import (
	"net/http"

	"github.com/nzhong/minimal-go-postgres/step3/controllers"
	"github.com/nzhong/minimal-go-postgres/step3/database"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()

	// STATIC
	staticFileServer := http.FileServer(http.Dir("./static/"))
	staticFileHandler := http.StripPrefix("/", staticFileServer)
	r.Handle("/", staticFileHandler).Methods("GET")

	// NON-STATIC
	r.HandleFunc("/api/v1/books", controllers.AllBooks).Methods("GET")
	r.HandleFunc("/api/v1/book", controllers.GetBook).Methods("GET")
	return r
}

func main() {
	database.InitDB()

	r := newRouter()
	http.ListenAndServe(":8080", r)
}
