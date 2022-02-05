package routers

import (
	"shortener/src/routers/router"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return router.Settings(r)
}
