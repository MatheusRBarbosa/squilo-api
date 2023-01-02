package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "squilo",
	Short: "squilo base command",
	Long:  "squilo base command",
}

func Execute() {
	defer func() {
		err := recover()
		if err != nil {
			log.Fatal("Unexpected error:", err)
		}
	}()

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal("Error on Execute:", err)
	}
}
