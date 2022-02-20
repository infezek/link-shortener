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
	shortenerDb, err := repo.Db.Exec(sql_statement, shortener.UrlShortened, shortener.UrlOriginal, shortener.UserId)

	if err != nil {
		fmt.Println("err", err)
		return 0, nil
	}
	id, _ := shortenerDb.LastInsertId()

	return id, nil
}

func (repo *ShortenerRepositoryDb) FindAll() ([]entity.Shorteners, error) {
	sql_statement := "SELECT id, url_shortened, url_original, user_id from shorteners;"
	shortenersDb, err := repo.Db.Query(sql_statement)

	if err != nil {
		fmt.Println(err)
		return []entity.Shorteners{}, err
	}

	var shortenersReponse []entity.Shorteners

	for shortenersDb.Next() {
		var shortenerReponse entity.Shorteners
		if err = shortenersDb.Scan(
			&shortenerReponse.ID,
			&shortenerReponse.UrlShortened,
			&shortenerReponse.UrlOriginal,
			&shortenerReponse.UserId,
		); err != nil {
			return nil, err
		}
		shortenersReponse = append(shortenersReponse, shortenerReponse)
	}

	return shortenersReponse, nil
}
