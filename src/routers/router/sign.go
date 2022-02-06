package router

import (
	"database/sql"
	"fmt"
	"net/http"
)

var routersSign = []Router{
	{
		URI:    "/signin",
		Method: http.MethodGet,
		Function: func(db *sql.DB) http.HandlerFunc {
			return http.HandlerFunc(
				func(w http.ResponseWriter, req *http.Request) {
					fmt.Println(&db)
					fmt.Fprintf(w, "Sign In")
				},
			)
		},
		RequiresAuthentication: false,
	},
	{

		URI:    "/signup",
		Method: http.MethodGet,
		Function: func(db *sql.DB) http.HandlerFunc {
			return http.HandlerFunc(
				func(w http.ResponseWriter, req *http.Request) {
					fmt.Println(&db)
					fmt.Fprintf(w, "Sign Up")
				},
			)
		},
		RequiresAuthentication: false,
	},
}
