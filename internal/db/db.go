package db

import (
	"database/sql"
)

// CreateTable ensures the URLs table exists
func CreateTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS urls (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    short_url TEXT NOT NULL,
	    origin_url TEXT NOT NULL
	);`
	_, err := db.Exec(query)
	return err
}

// StoreURL inserts a new short URL and the original URL into the database
func StoreURL(db *sql.DB, shortURL, originURL string) error {
	query := `INSERT INTO urls (short_url, origin_url) VALUES (?, ?)`
	_, err := db.Exec(query, shortURL, originURL)
	return err
}

// GetOriginalURL fetches the original URL by the short URL
func GetOriginalURL(db *sql.DB, shortURL string) (string, error) {
	var originURL string
	query := `SELECT origin_url FROM urls WHERE short_url = ? LIMIT 1`
	err := db.QueryRow(query, shortURL).Scan(&originURL)
	if err != nil {
		return "", err
	}
	return originURL, nil
}
