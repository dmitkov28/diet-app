package data

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type DB struct {
	db *sql.DB
}

func NewDB() (*DB, error) {
	// db, err := sql.Open("sqlite3", "./db.db")
	tursoUrl := os.Getenv("TURSO_URL")
	tursoToken := os.Getenv("TURSO_TOKEN")
	dbConnectionString := fmt.Sprintf("%s?authToken=%s", tursoUrl, tursoToken)

	db, err := sql.Open("libsql", dbConnectionString)
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}
