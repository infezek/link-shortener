package database

import (
	"database/sql"
	"fmt"
	"shortener/src/config"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.StringDatabase())

	if err != nil {
		return nil, error(err)
	}

	_, err = db.Exec(SqlInit)
	if err != nil {
		fmt.Println("Err database: ", err)
	}

	return db, nil
}
