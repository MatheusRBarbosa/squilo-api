package crossCutting

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/matheusrbarbosa/gofin/domain/models"
)

var appEnvs *models.Env

func init() {
	LoadEnvs()
}

func LoadEnvs() {
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

	appEnvs = &envs
}

func GetConnectionString() string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database-%s",
		appEnvs.DB_USER, appEnvs.DB_PASSWORD, appEnvs.DB_HOST, appEnvs.DB_PORT, appEnvs.DB_NAME)
}
