package crossCutting

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/matheusrbarbosa/gofin/domain/models"
)

var AppEnvs *models.Env

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
		JWT_SECRET:  os.Getenv("JWT_SECRET"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
	}

	AppEnvs = &envs
}

func GetConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		AppEnvs.DB_USER, AppEnvs.DB_PASSWORD, AppEnvs.DB_HOST, AppEnvs.DB_PORT, AppEnvs.DB_NAME)
}
