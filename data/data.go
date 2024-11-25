package data

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func NewDB() (*DB, error) {
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}
