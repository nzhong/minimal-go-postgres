package database

// Book ...
type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

// AllBooks ...
func AllBooks() ([]Book, error) {

	rows, err := db.Query("SELECT isbn, title, author, price FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		var isbn string
		var title string
		var author string
		var price float32

		err := rows.Scan(&isbn, &title, &author, &price)
		if err != nil {
			return nil, err
		}
		book := Book{Isbn: isbn, Title: title, Author: author, Price: price}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

// GetBook ...
func GetBook(isbn string) (Book, error) {
	book := Book{}

	var title string
	var author string
	var price float32

	db.Query("SELECT isbn, title, author, price FROM books")
	err := db.QueryRow(`SELECT title, author, price FROM books where isbn = $1`, isbn).Scan(&title, &author, &price)
	if err == nil {
		book = Book{Isbn: isbn, Title: title, Author: author, Price: price}
	}

	return book, err
}
