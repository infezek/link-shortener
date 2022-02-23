package controllers

import (
	"database/sql"
	"net/http"
	"shortener/src/responses"
)

func HomeRouter(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			responses.Json(w, 200, map[string]string{"message": "ok"})
		},
	)
}
