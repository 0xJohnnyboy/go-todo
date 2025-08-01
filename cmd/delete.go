package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
    "todo/internal/task"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "removes a task from the list",
	Long: `removes a task from the list and saves the changes`,
	Run: func(cmd *cobra.Command, args []string) {
        err := task.DeleteTask(args[0])
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("task removed")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
