package middleware

import (
	"net/http"
	"shortener/src/responses"
	"shortener/src/security"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := security.ValidateToken(r)
		if err != nil {
			responses.Json(w, 401, map[string]string{"message": "Token invalido"})
			return
		}

		next(w, r)
	}
}
