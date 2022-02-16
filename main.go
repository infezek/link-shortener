package main

import (
	"fmt"
	"net/http"
	"shortener/src/database"
	"shortener/src/routers"
)

func main() {
	fmt.Println("Link Shortener")
	db, err := database.Connect()
	if err != nil {
		fmt.Println("Could not start the database")
	}
	routers := routers.Generate(db)

	http.ListenAndServe(":8000", routers)
}
