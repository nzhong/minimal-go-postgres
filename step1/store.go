package main

import (
	"database/sql"
)

// Store ...
type Store interface {
	//	CreatePerson(person *Person) error
	//	GetPerson() ([]*Person, error)
}

// `dbStore` struct implements the `Store` interface. Variable `db` takes the pointer
// to the SQL database connection object.
type dbStore struct {
	db *sql.DB
}

// Create a global `store` variable of type `Store` interface. It will be initialized
// in `func main()`.
var store Store