package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"todo/internal/task"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "marks one or more tasks as done",
	Long:  `marks one or more tasks as done and saves the changes`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("missing arguments")
			return
		}
		for _, arg := range args {
			err := task.DoneTask(arg)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("task %s marked as done.\n", arg)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
