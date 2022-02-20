package repositories

import (
	"database/sql"
	"fmt"
	"shortener/src/entity"
)

type ShortenerRepositoryDb struct {
	Db *sql.DB
}

func (repo *ShortenerRepositoryDb) Insert(shortener entity.Shorteners) (int64, error) {
	sql_statement := "INSERT INTO shorteners (url_shortened, url_original, user_id) VALUES ($1, $2, $3);"
	userDb, err := repo.Db.Exec(sql_statement, shortener.UrlShortened, shortener.UrlOriginal, shortener.UserId)

	if err != nil {
		fmt.Println("err", err)
		return 0, nil
	}
	id, _ := userDb.LastInsertId()

	return id, nil

}
