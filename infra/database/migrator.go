package database

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"

	goose "github.com/pressly/goose/v3"
)

var embedMigrations embed.FS

func main() {
	arg := os.Args[0]
	fmt.Print(arg)
}

func up() {
	log.SetFlags(0)
	db, err := sql.Open("mssql", "embed_example.sql")
	goose.SetBaseFS(embedMigrations)
	if err != nil {
		log.Fatal(err)
	}

	goose.SetDialect("mssql")

	if err := goose.Up(db, "migrations"); err != nil { //
		panic(err)
	}
	if err := goose.Version(db, "migrations"); err != nil {
		log.Fatal(err)
	}
}

func ping() {
	fmt.Printf("pongando!")
}
