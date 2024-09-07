package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "GitHubBoardCLI",
	Short: "A CLI tool to manage GitHub project Kanban boards",
	Long:  `A CLI tool that can move missing issues and PRs to the GitHub project Kanban board.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to GitHubBoardCLI!")
		fmt.Println("Use 'GitHubBoardCLI --help' to see available commands.")
	},
}

var moveCmd = &cobra.Command{
	Use:   "move [issue/PR number] [column]",
	Short: "Move an issue or PR to a specific column",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Moving issue/PR %s to column '%s'\n", args[0], args[1])
		// Implement the actual moving logic here
	},
}

func init() {
	rootCmd.AddCommand(moveCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}