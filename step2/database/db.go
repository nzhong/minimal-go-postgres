package database

import (
	"database/sql"
)

var db *sql.DB

// InitDB ...
func InitDB() {
	connString := `user=postgres 
				   password=welcome
				   host=localhost
				   port=5432
				   dbname=hello 
				   sslmode=disable`
	db, _ = sql.Open("postgres", connString)
	err := db.Ping()
	if err != nil {
		panic(err)
	}
}
