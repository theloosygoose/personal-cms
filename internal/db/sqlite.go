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
	return db, nil
}

func CloseConnection(db *sql.DB) {
	defer db.Close()
}

func Init(db *sql.DB) error {
    query := `CREATE TABLE IF NOT EXISTS articles
    (id INTEGER PRIMARY KEY,
    title VARCHAR(100),
    body BLOB,
    creation_date DATE,
    modify_date DATETIME);`

    results, err := db.Exec(query)

    if err != nil {
        log.Println("Error in Creating Table")
        return err
    }

    log.Println("Table created successfully", results)

    return nil
}
