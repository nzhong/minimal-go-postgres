package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nzhong/minimal-go-postgres/step3/database"
)

// AllBooks ... /api/v1/books
func AllBooks(res http.ResponseWriter, req *http.Request) {
	bks, err := database.AllBooks()
	if err != nil {
		http.Error(res, http.StatusText(500), 500)
		return
	}
	//for _, bk := range bks {
	//	fmt.Fprintf(res, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	//}
	jsonData, err := json.Marshal(bks)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonData)
}

// GetBook ... /api/v1/book?idbn=...
func GetBook(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	isbn := params.Get("isbn")

	bk, err := database.GetBook(isbn)
	if err != nil {
		http.Error(res, http.StatusText(500), 500)
		return
	}

	// fmt.Fprintf(res, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	jsonData, err := json.Marshal(bk)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonData)
}
