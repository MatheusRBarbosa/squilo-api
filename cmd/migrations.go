package cmd

import (
	"log"

	migrator "github.com/matheusrbarbosa/gofin/infra/database"
	"github.com/spf13/cobra"
)

var (
	migrationCommand = &cobra.Command{
		Use:   "migrate [up|down]",
		Short: "Migrate or rollback migrations",
		Long:  "Migrate or rollback migrations",
		Args:  cobra.ExactArgs(1),
		RunE:  migrateExecute,
	}
)

func init() {
	rootCmd.AddCommand(migrationCommand)
}

func migrateExecute(cmd *cobra.Command, args []string) error {
	if args[0] == "up" {
		migrator.Up()
		return nil
	}

	if args[0] == "down" {
		migrator.Down()
		return nil
	}

	log.Fatal("Command flag not supported")
	return nil
}
