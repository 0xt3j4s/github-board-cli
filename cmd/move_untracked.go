package cmd

import (
	"fmt"
	// "io"
	// "os"


	ghBoardCLI "github.com/0xt3j4s/github-board-cli/pkg"

	// "github.com/google/go-github/v39/github"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

type moveUntracked struct {
	Args []string
	Org string
	Repo string
	User string
	ProjectName string
	ColumnName string
}


func NewMoveUntrackedCommand (client ghBoardCLI.Provider) *cobra.Command {
	var opts moveUntracked

	cmd := &cobra.Command {
		Use: "move-untracked",
		Short: "Move the untracked issues and PRs to the desired column in the required project for ${GITHUB_OWNER}",
		RunE: func (cmd *cobra.Command, args []string) error {
			return MoveUntracked(client, opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&opts.Org, "org", "", "The organisation to get projects for")
	flags.StringVar(&opts.Repo, "repo", "", "Move Untracked issues and PRs for a specific repo or for all (default)")
	flags.StringVar(&opts.User, "user", "", "The user to get projects for")
	flags.StringVar(&opts.ProjectName, "project-name", "", "Project in which we need to move the untracked issues and PRs")
	flags.StringVar(&opts.ColumnName, "column-name", "", "Target column to move the untracked issues and PRs")

	return cmd
}

func MoveUntracked (client ghBoardCLI.Provider, userOpts moveUntracked) error {
	
	// get project by name
	// from project ID get column ID using column Name
	// call another function to move the untracked items


	var project *github.Project	= client.GetProjectByName(userOpts.ProjectName, userOpts.Org, userOpts.User, userOpts.Repo)

	var columnId int64 
	columnId, err := client.GetColumnID(userOpts.ProjectName, userOpts.ColumnName, userOpts.Org, userOpts.User, userOpts.Repo)
	if (err != nil) {
		fmt.Println("error fetching the column")
	}

	fmt.Println("ColumnId: ", columnId)
	fmt.Println("project: ", project.URL);

	if err := client.MoveUntrackedIssues(); err != nil {
		return fmt.Errorf("error moving untracked issues: %v", err)
	}

	// if err := moveUntrackedPRs(client, userOpts, columnID); err != nil {
	// 	return fmt.Errorf("error moving untracked PRs: %v", err)
	// }
	
	return nil
}


