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

func GetByIdShortener(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			parametros := mux.Vars(r)
			shortenerID := parametros["shortenerID"]

			repository := repositories.ShortenerRepositoryDb{Db: db}
			shortenersRepository, err := repository.FindById(shortenerID)
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
