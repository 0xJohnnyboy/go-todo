package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
    "todo/internal/logic"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "marks a task as done",
	Long: `marks a task as done and saves the changes`,
	Run: func(cmd *cobra.Command, args []string) {
        err := logic.DoneTask(args[0])
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("task marked as done")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
