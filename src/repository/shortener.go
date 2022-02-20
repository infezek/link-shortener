package repositories

import (
	"database/sql"
	"errors"
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

func (repo *ShortenerRepositoryDb) FindByID(shortenerId string) (entity.Shorteners, error) {
	sql_statement := "SELECT id, url_shortened, url_original, user_id from shorteners WHERE id = $1;"
	repositoryShortener, err := repo.Db.Query(sql_statement, shortenerId)
	if err != nil {
		fmt.Println(err)
		return entity.Shorteners{}, err
	}

	var shortener entity.Shorteners

	for repositoryShortener.Next() {
		if err := repositoryShortener.Scan(
			&shortener.ID,
			&shortener.UrlShortened,
			&shortener.UrlOriginal,
			&shortener.UserId,
		); err != nil {
			fmt.Println(err)
			return entity.Shorteners{}, nil
		}
	}

	return shortener, nil
}

func (repo *ShortenerRepositoryDb) DeleteByID(shortenerID string) error {
	sql_statement := "DELETE FROM shorteners WHERE id = $1;"
	shortener, err := repo.Db.Exec(sql_statement, shortenerID)
	if err != nil {
		return err
	}
	isDeleted, _ := shortener.RowsAffected()

	if isDeleted == 1 {
		return nil
	}
	return errors.New("url não encontrada")
}

func (repo *ShortenerRepositoryDb) FindByUserID(userID string) ([]entity.Shorteners, error) {
	sql_statement := "SELECT id, url_shortened, url_original, user_id FROM shorteners where user_id = $1"

	repositoryRows, err := repo.Db.Query(sql_statement, userID)
	if err != nil {
		return []entity.Shorteners{}, err
	}

	var repositories []entity.Shorteners

	for repositoryRows.Next() {
		var repository entity.Shorteners
		if err := repositoryRows.Scan(
			&repository.ID,
			&repository.UrlShortened,
			&repository.UrlOriginal,
			&repository.UserId,
		); err != nil {
			return []entity.Shorteners{}, nil
		}
		repositories = append(repositories, repository)

	}
	return repositories, nil
}
