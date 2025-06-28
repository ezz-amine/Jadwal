package core

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ezz-amine/Jadwal/pkg/sqlc"
	_ "modernc.org/sqlite"
)

func getDBPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintf("Could not get home directory: %v", err))
	}
	return filepath.Join(home, ".local", "share", ".jadwal.db")
}

func OpenDB() (*sql.DB, error) {
	dbPath := getDBPath()

	// Ensure the directory exists
	if err := os.MkdirAll(filepath.Dir(dbPath), 0o755); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Enable WAL mode for better concurrency
	if _, err := db.Exec("PRAGMA journal_mode=WAL;"); err != nil {
		return nil, fmt.Errorf("failed to enable WAL mode: %w", err)
	}

	// Enable foreign key constraints
	if _, err := db.Exec("PRAGMA foreign_keys=ON;"); err != nil {
		return nil, fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	return db, nil
}

func GetQueries() (*sqlc.Queries, error) {
	db, err := OpenDB()
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}
	return sqlc.New(db), nil
}
