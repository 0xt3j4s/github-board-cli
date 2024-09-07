package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "github-board-cli",
	Short: "CLI tool for managing GitHub project boards",
	Long: `github-board-cli is a command line interface for managing GitHub project boards.
It allows you to move issues and pull requests to different columns on your project board.`,
	Run: func(cmd *cobra.Command, args []string) {
		// If no subcommand is provided, print the help
		cmd.Help()
	},
}

var moveCmd = &cobra.Command{
	Use:   "move [issue/PR number] [column]",
	Short: "Move an issue or PR to a specific column",
	Long: `Move an issue or pull request to a specific column on your GitHub project board.
For example:
  github-board-cli move 123 "In Progress"
This will move issue or PR #123 to the "In Progress" column.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Moving issue/PR %s to column '%s'\n", args[0], args[1])
		// Implement the actual moving logic here
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all columns in the project board",
	Long: `List all columns currently present in your GitHub project board.
This command will display the names of all columns in your board.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing all columns in the project board...")
		// Implement the listing logic here
	},
}

func init() {
	rootCmd.AddCommand(moveCmd)
	rootCmd.AddCommand(listCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}