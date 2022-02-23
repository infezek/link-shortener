package routers

import (
	"database/sql"
	"shortener/src/routers/router"

	"github.com/gorilla/mux"
)

func Generate(db *sql.DB) *mux.Router {

	r := mux.NewRouter()
	return router.Settings(r, db)
}
