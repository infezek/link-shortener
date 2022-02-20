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

	return fmt.Sprintf(`
	host=localhost 
	port=5432 
	user=%s
	password=%s 
	dbname=%s 
	sslmode=disable`,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}

type Settings struct {
	SecretKey string
}

func ProjectSettings() Settings {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	config := Settings{
		os.Getenv("SECRET_KEY"),
	}
	return config
}
