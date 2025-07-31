package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "a todo list CLI",
	Long: `this is a todo list CLI based on cobra, including API endpoints and a sqlite database`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}


