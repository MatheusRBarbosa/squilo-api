package crossCutting

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/matheusrbarbosa/gofin/domain/models"
)

func LoadEnvs() models.Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envs := models.Env{
		APP_ENV:     os.Getenv("APP_ENV"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
	}

	return envs
}

func GetConnectionString() string {
	envs := LoadEnvs()
	return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database-%s",
		envs.DB_USER, envs.DB_PASSWORD, envs.DB_HOST, envs.DB_PORT, envs.DB_NAME)
}
