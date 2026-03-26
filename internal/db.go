package internal

import (
	"database/sql"
	"fmt"
	"os"
)

func OpenDB() (*sql.DB, error) {
	path := os.Getenv("WTT_DB")
	if path == "" {
		return nil, fmt.Errorf("invalid path to db")
	}

	db, err := sql.Open("sqlite", path)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("cant connect to db: %w", err)
	}

	return db, nil
}
