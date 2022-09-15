package database

import (
	crossCutting "github.com/matheusrbarbosa/gofin/crosscutting"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DbContext *gorm.DB

func ConnectDatabase() {
	connectionString := crossCutting.GetConnectionString()

	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DbContext = db
}
