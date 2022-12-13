package crossCutting

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var envs map[string]string
var connectionString string

func init() {
	loadEnvs()
	connectionString = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		GetEnv("DB_USER"), GetEnv("DB_PASSWORD"), GetEnv("DB_HOST"), GetEnv("DB_PORT"), GetEnv("DB_NAME"))
}

func GetEnv(name string) string {
	return envs[name]
}

func GetConnectionString() string {
	return connectionString
}

func loadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envs = map[string]string{
		"APP_ENV":     os.Getenv("APP_ENV"),
		"JWT_SECRET":  os.Getenv("JWT_SECRET"),
		"DB_HOST":     os.Getenv("DB_HOST"),
		"DB_PORT":     os.Getenv("DB_PORT"),
		"DB_NAME":     os.Getenv("DB_NAME"),
		"DB_USER":     os.Getenv("DB_USER"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
	}
}
