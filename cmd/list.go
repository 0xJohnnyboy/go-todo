package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
    "todo/internal/todo"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists existing tasks",
	Long: `retrieves tasks from the database and prints them`,
	Run: func(cmd *cobra.Command, args []string) {
        err := todo.ListTasks()
        if err != nil {
            fmt.Println(err)
            return
        }
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
