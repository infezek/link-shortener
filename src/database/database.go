package database

import (
	"database/sql"
	"fmt"
	"shortener/src/config"

	_ "github.com/bmizerany/pq"
)

func Connect() (*sql.DB, error) {
	fmt.Println(config.StringDatabase())
	db, err := sql.Open("postgres", config.StringDatabase())

	if err != nil {
		fmt.Println(err)
		return nil, error(err)
	}
	return db, nil
}
