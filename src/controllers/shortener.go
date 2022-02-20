package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"shortener/src/entity"
	repositories "shortener/src/repository"
	"shortener/src/responses"
	"shortener/src/security"

	"github.com/gorilla/mux"
)

type Shortener struct {
	UrlOriginal string
	UserId      string
}

func GetAllShortener(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			repositories := repositories.ShortenerRepositoryDb{Db: db}
			shortenersRepository, err := repositories.FindAll()
			if err != nil {
				log.Fatal("Erro", err)
				return
			}

			responses.Json(w, 200, shortenersRepository)
			return
		},
	)
}

func GetByIDShortener(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			parametros := mux.Vars(r)
			shortenerID := parametros["shortenerID"]

			repository := repositories.ShortenerRepositoryDb{Db: db}
			shortenersRepository, err := repository.FindByID(shortenerID)
			if err != nil {
				responses.Json(w, 400, map[string]string{"message": "não foi possivel encontrar essa url"})
				return
			}
			responses.Json(w, 200, shortenersRepository)
			return
		},
	)
}

func CreateShortener(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			body, erro := ioutil.ReadAll(r.Body)
			if erro != nil {
				return
			}

			decodeJwt := security.DecodeToken(r)

			shortener := struct {
				UrlOriginal string
				UserId      string
			}{}

			if erro = json.Unmarshal(body, &shortener); erro != nil {
				return
			}

			shortenerEntity := entity.Shorteners{
				UrlOriginal: shortener.UrlOriginal,
				UserId:      decodeJwt.Sub,
			}
			shortenerEntity, err := shortenerEntity.Validate()

			shortenerFormated := entity.Shorteners{
				UrlShortened: shortenerEntity.UrlShortened,
				UrlOriginal:  shortener.UrlOriginal,
				UserId:       decodeJwt.Sub,
			}

			repositorios := repositories.ShortenerRepositoryDb{Db: db}
			repositorios.Insert(shortenerFormated)

			if err != nil {
				responses.Json(w, 400, map[string]string{"message": err.Error()})
				return
			}
			responses.Json(w, 200, map[string]string{"message": "ok"})
			return
		},
	)
}

func DeleteByID(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			param := mux.Vars(r)
			shortenerID := param["shortenerID"]

			repository := repositories.ShortenerRepositoryDb{Db: db}
			err := repository.DeleteByID(shortenerID)

			if err != nil {
				responses.Json(w, 400, map[string]string{"message": "não foi possivel encontrar a url"})
				return
			}
			responses.Json(w, 200, map[string]string{"message": "url foi deletada"})
			return
		},
	)
}

func FindByUserID(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			bodyToken := security.DecodeToken(r)
			repositoryDB := repositories.ShortenerRepositoryDb{Db: db}
			repository, err := repositoryDB.FindByUserID(bodyToken.Sub)

			if err != nil {
				responses.Json(w, 400, err)
				return
			}

			responses.Json(w, 200, repository)
			return
		},
	)
}
