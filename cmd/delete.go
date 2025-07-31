package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
    "todo/internal/todo"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "removes a task from the list",
	Long: `removes a task from the list and saves the changes`,
	Run: func(cmd *cobra.Command, args []string) {
        err := todo.DeleteTask(args[0])
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("task removed")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
