package cmd

import (
	"fmt"
    "time"

	"github.com/spf13/cobra"
    "todo/internal/todo"
)

var showDoneFlag bool
var showUndoneFlag bool
// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists existing tasks",
	Long: `retrieves tasks from the database and prints them`,
	Run: func(cmd *cobra.Command, args []string) {
        tasks, err := todo.ListTasks(showDoneFlag, showUndoneFlag)
        if err != nil {
            fmt.Println(err)
            return
        }

        if len(tasks) == 0 {
            fmt.Println("No tasks")
            return
        }

        for _, task := range tasks {
            fmt.Printf("%d: %s\n", task.Id, task.Title)
            fmt.Printf("  Done: %t\n", task.Done)
            fmt.Printf("  Created at: %s\n", task.CreatedAt.Format(time.RFC3339))
            fmt.Printf("  Updated at: %s\n", task.UpdatedAt.Format(time.RFC3339))
        }
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
    listCmd.Flags().BoolVarP(&showDoneFlag, "done", "d", false, "Show only done tasks")
    listCmd.Flags().BoolVarP(&showUndoneFlag, "undone", "u", false, "Show only undone tasks")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
