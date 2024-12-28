package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"log"
	"os"
)

type DB struct {
	db *sql.DB
}

func NewDB() (*DB, error) {
	env := os.Getenv("ENV")
	var dbConnectionString string

	if env == "PROD" {
		tursoUrl := os.Getenv("TURSO_URL")
		tursoToken := os.Getenv("TURSO_TOKEN")
		dbConnectionString = fmt.Sprintf("%s?authToken=%s", tursoUrl, tursoToken)
	} else {
		localDbUrl := os.Getenv("LOCAL_DB_URL")
		if localDbUrl == "" {
			log.Fatal("Missing DB connection details.")
		}
		dbConnectionString = localDbUrl
	}

	if dbConnectionString == "" {
		log.Fatal("Missing DB connection details.")
	}

	db, err := sql.Open("libsql", dbConnectionString)
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}
