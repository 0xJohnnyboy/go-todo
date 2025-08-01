package cmd

import (
	"fmt"
    "time"
    "strings" 

	"github.com/spf13/cobra"
    "todo/internal/logic"
)

var showDoneFlag bool
var showAllFlag bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists existing tasks",
	Long: `retrieves tasks from the database and prints them`,
	Run: func(cmd *cobra.Command, args []string) {
        tasks, err := logic.ListTasks(showDoneFlag, showAllFlag)
        if err != nil {
            fmt.Println(err)
            return
        }

        if len(tasks) == 0 {
            fmt.Println("No tasks")
            return
        }

        var outputHeader string
        switch {
        case showAllFlag:
            outputHeader = fmt.Sprintf("All Tasks (%d)", len(tasks))
        case showDoneFlag:
            outputHeader = fmt.Sprintf("Completed Tasks (%d)", len(tasks))
        default:
            outputHeader = fmt.Sprintf("Todo Tasks (%d)", len(tasks))
        }

        fmt.Println(outputHeader)
        fmt.Println(strings.Repeat("=", len(outputHeader)))
        fmt.Println()

        var doneEmoji string
        for _, task := range tasks {
            header := fmt.Sprintf("%d: %s", task.Id, task.Title)
            fmt.Printf("%d: %s\n", task.Id, task.Title)
            separator := strings.Repeat("-", len(header))
            fmt.Printf("%s\n", separator)
            if task.Done {
                doneEmoji = "✅"
            } else {
                doneEmoji = "❌"
            }
            fmt.Printf("  Done: %s\n", doneEmoji)
            fmt.Printf("  Created at: %s\n", task.CreatedAt.Format(time.RFC3339))
            fmt.Printf("  Updated at: %s\n", task.UpdatedAt.Format(time.RFC3339))
            fmt.Println()
        }
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
    listCmd.Flags().BoolVarP(&showDoneFlag, "done", "d", false, "Show only done tasks")
    listCmd.Flags().BoolVarP(&showAllFlag, "all", "a", false, "Show all tasks")
}
