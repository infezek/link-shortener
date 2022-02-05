package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

func Settings(r *mux.Router) *mux.Router {
	routers := routersSign

	for _, router := range routers {
		r.HandleFunc(router.URI,
			router.Function,
		).Methods(router.Method)
	}
	return r
}
