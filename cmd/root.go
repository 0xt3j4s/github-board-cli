package cmd

import (
	// "fmt"
	"fmt"
	"os"
	"strings"

	"github.com/0xt3j4s/github-board-cli/pkg"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func Execute() {
	initConfig()

	cmd := NewRootCommand()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// NewRootCommand will return the application
func NewRootCommand() *cobra.Command {

	var rootCmd = &cobra.Command{
		Use:   "github-board-cli",
		Short: "A CLI tool for managing GitHub project boards",
		Long:  `github-board-cli is a command line interface for managing GitHub project boards.`,
		Run: func(cmd *cobra.Command, args []string) {
			// If no subcommand is provided, print the help
			cmd.Help()
		},
	}

	// Configuration settings.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/tejas/ghcli.yaml)")

	
	client := pkg.NewClient()

	rootCmd.AddCommand(
		NewMoveUntrackedCommand(&client),
	)

	return rootCmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		path := strings.Join([]string{home, ""}, "/")
		viper.AddConfigPath(path)
		viper.SetConfigName("ghcli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if fmt.Sprintf("%T", err) == "ConfigParseError" {
		fmt.Fprintf(os.Stderr, "Failed to load config: %s\n", err)
	}
}

