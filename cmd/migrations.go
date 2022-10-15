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

	createMigrationCommand = &cobra.Command{
		Use:   "create migration [name]",
		Short: "Create migration",
		Long:  "Create SQL migration file template",
		Args:  cobra.ExactArgs(2),
		RunE:  createMigrationExecute,
	}
)

func init() {
	rootCmd.AddCommand(migrationCommand)
	rootCmd.AddCommand(createMigrationCommand)
}

func migrateExecute(_ *cobra.Command, args []string) error {
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

func createMigrationExecute(_ *cobra.Command, args []string) error {
	if args[0] == "migration" {
		migrator.Create(args[1])
		return nil
	}

	log.Fatal("Command flag not supported")
	return nil
}
