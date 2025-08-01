package cmd

import (
	"fmt"
    "strings"
	"github.com/spf13/cobra"
    "todo/internal/todo"
)

var doneFlag bool

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a task to the list",
	Long: `adds a task to the list as undone`,
	Run: func(cmd *cobra.Command, args []string) {
        title := strings.Join(args, " ")
        task , err := todo.AddTask(title, doneFlag)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("task added with Id:%d\n", task.Id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
    addCmd.Flags().BoolVarP(&doneFlag, "done", "d", false, "Adds the task as already done")
}
