package main //I have to move all this stuff out of main if I wnant to share it with another frontend (like a rest api )

import (
	"bookstore/pkg/util"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type book struct {
	title  string
	author string
	rating int
}

func newBook(title string, author string, rating int) *book {
	//whats the best way to test for empty\incorrect values?
	b := book{title: title}
	b.author = author
	b.rating = rating
	return &b
}

func addBook() error {
	//create new book object
	bookEntry := newBook(
		CLI.Add.Title,
		CLI.Add.Author,
		CLI.Add.Rating,
	)
	f := "I am a string"
	g := util.AppendButt(f)
	fmt.Println(g)
	dbConfig := createDbSettings(driver, fileLocation)
	//do something with the book object (print, add to database etc)
	fmt.Println("Adding book entry: ", bookEntry)
	err := writeEntry(*dbConfig, bookEntry)
	if err != nil {
		return err
	}
	return nil

}

//interact with database
func writeEntry(db dbSettings, book *book) error {
	database, err := sql.Open(db.driver, db.fileLocation)
	if err != nil {
		return err
	}
	createDbStatement, err := database.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, title TEXT, author TEXT, rating INTEGER)")
	if err != nil {
		return err
	}
	createDbStatement.Exec()
	fmt.Println("Writing this book:", book)
	statement, _ := database.Prepare("INSERT INTO books (title, author, rating) VALUES (?, ?, ?)")
	statement.Exec(book.title, book.author, book.rating)

	return nil
}

func listAllEntries(db *dbSettings) error {
	database, _ := sql.Open(db.driver, db.fileLocation)
	rows, _ := database.Query("SELECT id, title, author, rating FROM books")
	var id int
	var title string
	var author string
	var rating int

	for rows.Next() {
		rows.Scan(&id, &title, &author, &rating)
		fmt.Printf("Book No. %v: Title: %v \n Author: %v \n Rating: %v \n", id, title, author, rating)
	}
	return nil
}
