package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
    "todo/internal/todo"
)

var portFlag int
var quietFlag bool
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs todo API server",
	Long: `Runs todo API server on default port 9876. 
    Port can be changed with the --port flag
    Also a quiet mode can be enabled with the --quiet flag to avoid printing logs in the console`,
	Run: func(cmd *cobra.Command, args []string) {
		if quietFlag {
			fmt.Println("Not implemented yet")
			return
		}
        todo.ServeApi(portFlag)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
    serveCmd.Flags().IntVarP(&portFlag, "port", "p", 9876, "Port to run the server on")
    serveCmd.Flags().BoolVarP(&quietFlag, "quiet", "q", false, "Quiet mode")
}
