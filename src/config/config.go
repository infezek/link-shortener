package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func StringDatabase() string {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	return fmt.Sprintf("postgresql://%s:%s@localhost:5432/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}