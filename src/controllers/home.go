package controllers

import (
	"database/sql"
	"net/http"
	"shortener/src/responses"
	"time"
)

var startedAt = time.Now()

func HomeRouter(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			err := db.Ping()
			duration := time.Since(startedAt)

			if duration.Seconds() < 25 {
				responses.Json(w, 400, map[string]string{"message": "time error"})
				return
			}
			if err != nil {
				responses.Json(w, 400, map[string]string{"message": "database error"})
				return
			}
			responses.Json(w, 200, map[string]string{"message": "ok"})
			return
		},
	)
}
