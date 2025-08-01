package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
    "todo/internal/task"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clears the list",
	Long: `clears the list and saves the changes`,
	Run: func(cmd *cobra.Command, args []string) {
        err := task.ClearTasks()
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("list cleared")
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
