package main

import (
	"log"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Add struct {
		Title  string `help:"Title of Book"`
		Author string `help:"Who wrote it"`
		Rating int    `help:"How many stars?"`
	} `cmd help:"Add a book"`
}

func main() {
	if err := mainCore(); err != nil {
		log.Fatal(err)
	}
}

func mainCore() error {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "add":
		return addBook()

	default:
		panic(ctx.Command())
	}
}
