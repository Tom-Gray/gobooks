package main

import (
	"fmt"
)

type book struct {
	title  string
	author string
	rating int
}

func newBook(title string, author string, rating int) *book {
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

	//do something with the book object (print, add to database etc)
	fmt.Println("Adding book entry: ", bookEntry)
	err := writeEntry(bookEntry)
	if err != nil {
		return err
	}
	return nil

}

func writeEntry(book *book) error {
	fmt.Println("Writing this book:", book)
	return nil
}
