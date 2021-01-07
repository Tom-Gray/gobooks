package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	driver       = "sqlite3"
	fileLocation = "./book.db"
)

type dbSettings struct {
	driver       string
	fileLocation string
}

func createDbSettings(driver string, fileLocation string) *dbSettings {
	db := dbSettings{driver: driver}
	db.fileLocation = fileLocation
	return &db
}

func initDatabase(dbSettings dbSettings) error {
	database, err := sql.Open(dbSettings.driver, dbSettings.fileLocation)
	createDbStatement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, title TEXT, author TEXT, rating INTEGER)")
	createDbStatement.Exec()
	if err != nil {
		return err
	}
	return nil
}
