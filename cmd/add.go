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
        err := todo.AddTask(title, doneFlag)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("task added")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
    addCmd.Flags().BoolVarP(&doneFlag, "done", "d", false, "Adds the task as already done")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
