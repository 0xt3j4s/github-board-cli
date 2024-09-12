package cmd

import (
	"fmt"
	"os"

	// "github.com/0xt3j4s/github-board-cli/cmd"
	"github.com/0xt3j4s/github-board-cli/pkg"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "github-board-cli",
	Short: "A CLI tool for managing GitHub project boards",
	Long:  `github-board-cli is a command line interface for managing GitHub project boards.`,
	Run: func(cmd *cobra.Command, args []string) {
		// If no subcommand is provided, print the help
		cmd.Help()
	},
}

func Execute() {
	cmd := NewRootCommand()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func NewRootCommand() *cobra.Command {
	client := pkg.NewClient()

	rootCmd.AddCommand(
		NewMoveUntrackedCommand(&client),
	)
}