package main

import (
	"fmt"
	"net/http"
	"shortener/src/routers"
)

func main() {
	fmt.Println("Link Shortener")
	routers := routers.Generate()
	http.ListenAndServe(":8000", routers)
}
