package storage

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	inf := "storage.sqlite.New"
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", inf, err)
	}

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS photoes(
			id INTEGER PRIMARY KEY,
			photo BLOB NOT NULL,
			person TEXT,
			date DATE NOT NULL,
			latitude FLOAT NOT NULL,
			longitude FLOAT NOT NULL,
			);
		CREATE INDEX IF NOT EXISTS idx_alias ON urls(alias)
	`)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", inf, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s : %w", inf, err)
	}

	return &Storage{db: db}, nil

}

func (s *Storage) Save(photo []byte, person string, date time.Time, latitude float32, longitude float32) error {
	inf := "storage.sqlite.SavePhoto"

	db, err := s.db.Prepare(`
		INSERT INTO photoes(photo, person, date, latitude, longitude) VALUES(?,?,?,?,?)
	`)
	if err != nil {
		return fmt.Errorf("%s : %w", inf, err)
	}

	_, err = db.Exec(photo, person, date, latitude, longitude)
	if err != nil {
		return fmt.Errorf("%s : %w", inf, err)
	}
	return nil
}
