package routers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", hello)
	return r
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Link Shortener")
}
