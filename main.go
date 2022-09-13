package main

import (
	"log"

	"github.com/matheusrbarbosa/gofin/cmd"
	crossCutting "github.com/matheusrbarbosa/gofin/crosscutting"
	"github.com/matheusrbarbosa/gofin/domain"
)

var (
	appEnvs domain.Env
)

func init() {
	appEnvs = crossCutting.LoadEnvs()
	log.Println(appEnvs.DB_HOST)
}

func main() {
	cmd.Execute()
}
