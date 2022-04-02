package dbmanager

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
}

func (sqlite SQLite) GetDBConnection() *sql.DB {
	database, myerror := sql.Open("sqlite3", "./test.db")
	if myerror != nil {
		log.Fatal(myerror)
	}
	return database
}

type DBImplementation struct {
	Conn IDatabaseConnection
}

