package main

import (
	"fmt"
	"net/http"
	"shortener/src/database"
	"shortener/src/routers"
)

func main() {
	fmt.Println("Link Shortener")
	routers := routers.Generate()
	_, err := database.Connect()
	if err != nil {
		fmt.Println("Could not start the database")
	}

	http.ListenAndServe(":8000", routers)
}
