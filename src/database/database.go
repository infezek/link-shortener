package database

import (
	"database/sql"
	"shortener/src/config"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.StringDatabase())

	if err != nil {
		return nil, error(err)
	}
	return db, nil
}
