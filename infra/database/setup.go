package database

import (
	crossCutting "github.com/matheusrbarbosa/squilo/crosscutting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbContext *gorm.DB

func Context() *gorm.DB {
	return dbContext
}

func ConnectDatabase() {
	connectionString := crossCutting.GetConnectionString()

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	dbContext = db
}
