package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shortener/src/entity"
	repositories "shortener/src/repository"
)

func SignIn(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

		},
	)
}

func SignUp(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			request, _ := ioutil.ReadAll(r.Body)

			var usuario entity.Users

			if err := json.Unmarshal(request, &usuario); err != nil {
				return
			}
			if err := usuario.Prepare(); err != nil {
				return
			}
			repositorios := repositories.UserRepositoryDb{Db: db}
			usuarioID, err := repositorios.Insert(usuario)
			if err != nil {
				return
			}
			fmt.Println(usuarioID)
		},
	)
}
