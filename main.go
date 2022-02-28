package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"shortener/src/database"
	"shortener/src/routers"

	"github.com/joho/godotenv"
)

func main() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	var port string = "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	fmt.Println("Link Shortener")
	db, err := database.Connect()
	if err != nil {
		fmt.Println("Could not start the database")
	}
	routers := routers.Generate(db)

	http.ListenAndServe(":"+port, routers)
}
