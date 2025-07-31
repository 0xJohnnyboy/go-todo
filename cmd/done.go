package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
    "todo/internal/todo"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "marks a task as done",
	Long: `marks a task as done and saves the changes`,
	Run: func(cmd *cobra.Command, args []string) {
        err := todo.DoneTask(args[0])
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("task marked as done")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
