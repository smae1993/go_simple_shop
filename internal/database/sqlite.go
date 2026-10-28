package database

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

func NewSQLiteConnection(dbPath, migrationPath string) (*sql.DB, error) {

	dsn := "file:" + dbPath + "?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)"
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	schema, err := os.ReadFile(migrationPath)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(string(schema)); err != nil {
		return nil, err
	}

	return db, nil
}
