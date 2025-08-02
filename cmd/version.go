package cmd

import (
	"fmt"
    v "todo/internal/version"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "shows app version",
	Long:  `shows app version`,
	Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Version: %s\nCommit: %s\nBuilt at: %s\n",
			v.Version,
			v.Commit,
			v.BuildTime,
		)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
