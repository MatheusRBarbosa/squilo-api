package cmd

import (
	server "github.com/matheusrbarbosa/gofin/api"
	"github.com/spf13/cobra"
)

var (
	apiCommand = &cobra.Command{
		Use:   "api",
		Short: "Initializes HTTP server",
		Long:  "Initializes the codebase to serve HTTP API server",
		RunE:  apiExecute,
	}
)

func init() {
	rootCmd.AddCommand(apiCommand)
}

func apiExecute(_ *cobra.Command, _ []string) error {
	server.StartHttpServer()
	return nil
}
