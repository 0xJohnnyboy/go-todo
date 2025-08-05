package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	u "todo/internal/user"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var userFlag string
var pwdFlag string

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register a new user",
	Long:  `Register prompts the user to enter username and password to create a new user for the app`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		username := userFlag
		if username == "" {
			fmt.Print("Username: ")
			username, _ := reader.ReadString('\n')
			username = strings.TrimSpace(username)
		} 

		if err := u.ValidateUsername(username); err != nil {
			fmt.Println(err)
			return 
		}

		password := pwdFlag
		if password == "" {
			fmt.Print("Password: ")
			passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
			fmt.Println()

			if err != nil {
				fmt.Println(err)
			}
			password = strings.TrimSpace(string(passwordBytes))
		} 
		if err := u.ValidatePassword(password); err != nil {
			fmt.Println(err)
			return
		}

		err := u.Register(username, password)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("User created")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringVarP(&userFlag, "username", "u", "", "username")
	registerCmd.Flags().StringVarP(&pwdFlag, "password", "p", "", "password")
}
