package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shortener/src/entity"
	"shortener/src/responses"
)

type Shortener struct {
	UrlOriginal string
	UserId      string
}

func GetShortener(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			body, erro := ioutil.ReadAll(r.Body)
			if erro != nil {
				fmt.Println("1")
				return
			}

			shortener := struct {
				UrlOriginal string
				UserId      string
			}{}

			if erro = json.Unmarshal(body, &shortener); erro != nil {
				fmt.Println(erro)
				return
			}

			shortenerEntity := entity.Shorteners{
				UrlOriginal: shortener.UrlOriginal,
				UserId:      shortener.UserId,
			}
			shortenerEntity, err := shortenerEntity.Validate()

			if err != nil {
				responses.Json(w, 400, map[string]string{"message": err.Error()})
				return
			}
			fmt.Println(shortenerEntity)
			responses.Json(w, 200, map[string]string{"message": "ok"})
			return
		},
	)
}
