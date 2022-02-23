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

func (repo *ShortenerRepositoryDb) RedirectURL(url string) (string, error) {
	sql_statement := "SELECT url_original, visits FROM shorteners where url_shortened = $1"
	repository, err := repo.Db.Query(sql_statement, url)

	if err != nil {
		return "", err
	}
	shortened := struct {
		urlOriginal string
		visits      int16
	}{}

	for repository.Next() {
		if err := repository.Scan(
			&shortened.urlOriginal,
			&shortened.visits,
		); err != nil {
			return "", err
		}
	}
	if shortened.urlOriginal == "" {
		return "", errors.New("url não encontrada")
	}

	sql_statement = "UPDATE shorteners SET visits=$1 where url_shortened = $2;"
	_, err = repo.Db.Exec(sql_statement, shortened.visits+1, url)

	if err != nil {
		return "", nil
	}

	return shortened.urlOriginal, nil
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
	sql_statement := "SELECT id, url_shortened, url_original, user_id, visits from shorteners;"
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
			&shortenerReponse.Visits,
		); err != nil {
			return nil, err
		}
		shortenersReponse = append(shortenersReponse, shortenerReponse)
	}

	return shortenersReponse, nil
}

func (repo *ShortenerRepositoryDb) FindByID(shortenerId string) (entity.Shorteners, error) {
	sql_statement := "SELECT id, url_shortened, url_original, user_id, visits from shorteners WHERE id = $1;"
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
			&shortener.Visits,
		); err != nil {
			fmt.Println(err)
			return entity.Shorteners{}, nil
		}
	}

	if shortener.ID == "" {
		return entity.Shorteners{}, errors.New("url não encontrada")
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
	sql_statement := "SELECT id, url_shortened, url_original, user_id, visits FROM shorteners where user_id = $1"

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
			&repository.Visits,
		); err != nil {
			return []entity.Shorteners{}, nil
		}
		repositories = append(repositories, repository)

	}
	return repositories, nil
}
