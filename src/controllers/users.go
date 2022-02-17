package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shortener/src/config"
	"shortener/src/entity"
	repositories "shortener/src/repository"
	"shortener/src/responses"
	"shortener/src/security"
	"time"

	"github.com/golang-jwt/jwt"
)

func SignIn(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			corpoRequisicao, erro := ioutil.ReadAll(r.Body)
			if erro != nil {
				return
			}

			var usuario entity.Users
			if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
				return
			}

			hmacSampleSecret := []byte(config.ProjectSettings().SecretKey)

			repositorios := repositories.UserRepositoryDb{Db: db}
			valor, err := repositorios.FindByEmail(usuario.Email)

			if err != nil {
				return
			}
			fmt.Println(valor)

			now := time.Now()
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub":      "123123",
				"Subject":  "teste",
				"IssuedAt": now.Unix(),
				"exp":      now.Add(time.Second * 60).Unix(),
			})
			tokenString, err := token.SignedString(hmacSampleSecret)

			if err != nil {
				fmt.Println(err)
			}
			responses.Json(w, 200, map[string]string{"token": tokenString})
		},
	)
}

func SignUp(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			request, _ := ioutil.ReadAll(r.Body)

			var user entity.Users

			if err := json.Unmarshal(request, &user); err != nil {
				return
			}
			passwordEncrypt, err := security.EncryptPassword(user.Password)
			user.Password = passwordEncrypt
			if err != nil {
				return
			}

			if err := user.Prepare(); err != nil {
				return
			}
			repositorios := repositories.UserRepositoryDb{Db: db}
			userID, err := repositorios.Insert(user)
			if err != nil {
				return
			}
			fmt.Println(userID)
		},
	)
}
