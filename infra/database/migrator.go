package database

import (
	"embed"
	"log"
)

var (
	embedMigrations embed.FS
)

func Up() {
	// log.SetFlags(0)
	// db, err := sql.Open("mssql", "embed_example.sql")
	// goose.SetBaseFS(embedMigrations)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// goose.SetDialect("mssql")

	// if err := goose.Up(db, "migrations"); err != nil { //
	// 	panic(err)
	// }
	// if err := goose.Version(db, "migrations"); err != nil {
	// 	log.Fatal(err)
	// }
	log.Print("Migrating...")
}

func Down() {
	log.Print("Rolling back...")
}
