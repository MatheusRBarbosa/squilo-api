package database

import (
	"database/sql"
	"embed"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	crossCutting "github.com/matheusrbarbosa/gofin/crosscutting"
	"github.com/pressly/goose/v3"
)

var (
	//go:embed migrations/*.sql
	embedMigrations  embed.FS
	connectionString string
	migrationsPath   string
)

func init() {
	connectionString = crossCutting.GetConnectionString()
	migrationsPath = "migrations"
}

func Up() {
	db, err := sql.Open("sqlserver", connectionString)
	goose.SetBaseFS(embedMigrations)
	if err != nil {
		log.Fatal(err)
	}

	if err := goose.SetDialect("mssql"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, migrationsPath); err != nil {
		panic(err)
	}

	if err := goose.Version(db, migrationsPath); err != nil {
		panic(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}

func Down() {
	db, err := sql.Open("sqlserver", connectionString)
	goose.SetBaseFS(embedMigrations)
	if err != nil {
		log.Fatal(err)
	}

	if err := goose.SetDialect("mssql"); err != nil {
		panic(err)
	}

	if err := goose.Down(db, migrationsPath); err != nil {
		panic(err)
	}

	if err := goose.Version(db, migrationsPath); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}
