package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"todo/internal/task"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "show stats about tasks",
	Long:  `show a summary of tasks, their status, number of tasks per day average, and completion rate`,
	Run: func(cmd *cobra.Command, args []string) {
		stats, err := task.GetStats()
		if err != nil {
			fmt.Println(err)
		}

		completionRate := float64(stats.Done) / float64(stats.Total) * 100

        fmt.Printf(`
Todo stats
----------
Tasks: %d
Done: %d
Completion rate: %.2f%%
Average per day: %.2f
`, stats.Total, stats.Done, completionRate, stats.PerDayAvg)
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
