package crossCutting

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/matheusrbarbosa/gofin/domain"
)

func LoadEnvs() domain.Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envs := domain.Env{
		APP_ENV:     os.Getenv("APP_ENV"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
	}

	return envs
}
