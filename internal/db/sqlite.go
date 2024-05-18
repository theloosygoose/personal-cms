package db

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/mattn/go-sqlite3"
)

type DB struct { *sql.DB }

func NewDB(db *sql.DB) DB {
    return DB {db}
}

func Connect() (*sql.DB, error) {
    connInfo := os.Getenv("DB_SQLITE")

	db, err := sql.Open("sqlite3", connInfo)
	if err != nil {
		log.Fatal("Could not Connect to DB", err)
        return nil, err
	}
	log.Println("Successfully Connected to db!")

    db.Ping()
	return db, nil
}

func CloseConnection(db DB) {
	defer db.Close()
}
