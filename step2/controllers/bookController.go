package controllers

import (
	"fmt"
	"net/http"

	"github.com/nzhong/minimal-go-postgres/step2/database"
)

// AllBooks ... /api/v1/books
func AllBooks(res http.ResponseWriter, req *http.Request) {

	bks, err := database.AllBooks()
	if err != nil {
		http.Error(res, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(res, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}

}
