package router

import (
	"database/sql"
	"net/http"
	"shortener/src/middleware"

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
		if router.RequiresAuthentication {
			r.HandleFunc(router.URI,
				middleware.Auth(
					router.Function(db),
				),
			).Methods(router.Method)
		} else {
			r.HandleFunc(router.URI,
				router.Function(db),
			).Methods(router.Method)
		}
	}
	return r
}
