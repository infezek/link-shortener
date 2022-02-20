package router

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	URI                    string
	Method                 string
	Function               func(db *sql.DB) http.HandlerFunc
	RequiresAuthentication bool
}

func Settings(r *mux.Router, db *sql.DB) *mux.Router {

	routers := append(routersShortener, routersSign...)

	for _, router := range routers {
		r.HandleFunc(router.URI,
			router.Function(db),
		).Methods(router.Method)
	}
	return r
}
