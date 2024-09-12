package pkg

import (
	"context"
	"fmt"
	"strings"
	// "time"

	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
	"github.com/spf13/viper"
)

type Client struct {
	GHClient *github.Client
	Ctx context.Context
	User string
	Org string
	Token string
}

func NewClient() Client {
	token := viper.GetString("GITHUB_TOKEN")
	owner := viper.GetString("GITHUB_OWNER")
	org := viper.GetString("GITHUB_ORG")

	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	http := oauth2.NewClient(ctx, tokenSource)
	ghClient := github.NewClient(http)

	client := Client{
		GHClient:      ghClient,
		Ctx: ctx,
		User:    owner,
		Org:     org,
		Token:   token,
	}
	
	return client
}

func (c *Client) ListProjectsForOrg(orgName string, opts github.ProjectListOptions) ([]*github.Project, *github.Response, error) {
	result, res, err := c.GHClient.Organizations.ListProjects(c.Ctx, orgName, &opts)
	if err != nil {
		return nil, nil, err
	}

	return result, res, err
}

// ListProjectsForRepo will return the projects defined against a repo
func (c *Client) ListProjectsForRepo(repoName string, opts github.ProjectListOptions) ([]*github.Project, *github.Response, error) {
	parts := strings.Split(repoName, "/")

	var user, repo string
	if len(parts) < 2 {
		user = c.User
		repo = parts[0]
	} else {
		user = parts[0]
		repo = parts[1]
	}

	result, res, err := c.GHClient.Repositories.ListProjects(c.Ctx, user, repo, &opts)

	if err != nil {
		return nil, nil, err
	}

	return result, res, err
}


func (c *Client) ListProjectsForUser (userName string, opts github.ProjectListOptions) ([] *github.Project, *github.Response, error) {
	result, res, err := c.GHClient.Users.ListProjects(c.Ctx, userName, &opts)
	if err != nil {
		return nil, nil, err
	}

	return result, res, err
}

// GetProjectByName will return a single project given a name
func (c *Client) GetProjectByName(name, org, user, repo string) *github.Project {
	var projects []*github.Project

	if org != "" {
		opts := github.ProjectListOptions{}
		projects, _, _ = c.ListProjectsForOrg(org, opts)
	} else if user != "" {
		opts := github.ProjectListOptions{}
		projects, _, _ = c.ListProjectsForUser(user, opts)
	} else if repo != "" {
		opts := github.ProjectListOptions{}
		projects, _, _ = c.ListProjectsForRepo(repo, opts)
	}

	for _, project := range projects {
		if project.GetName() == name {
			return project
		}
	}

	return nil
}

// ListColumnsForProject will return columns for a project board
func (c *Client) ListColumnsForProject(projectName, org, user, repo string) ([]*github.ProjectColumn, error) {
	project := c.GetProjectByName(projectName, org, user, repo)
	opts := github.ListOptions{}
	columns, _, err := c.GHClient.Projects.ListProjectColumns(c.Ctx, project.GetID(), &opts)
	if err != nil {
		return nil, err
	}

	return columns, nil
}

func (c *Client) GetColumnID (projectName, columnName, org, user, repo string) (int64, error) {
	columns, err := c.ListColumnsForProject(projectName, org, user, repo)
	if err != nil {
		return 0, err
	}

	for _, column := range columns {
		if *column.Name == columnName {
			return *column.ID, nil
		}
	}

	return 0, fmt.Errorf("column '%s' not found in the project", columnName)
}

func (c *Client) MoveUntrackedIssues() error {
	client := c.GHClient
	issues, _, err := client.Issues.ListByOrg(c.Ctx, c.Org, &github.IssueListOptions{State: "open"})
	if err != nil {
		return err
	}

	for _, issue := range issues {
		fmt.Println("Issue: ", issue.Title)
	}

	// for _, issue := range issues {
	// 	if err := moveItemIfUntracked(c.Ctx, client, issue.GetID(), columnID); err != nil {
	// 		fmt.Printf("Error moving issue #%d: %v\n", issue.GetNumber(), err)
	// 	} else {
	// 		fmt.Printf("Moved issue #%d to the project\n", issue.GetNumber())
	// 	}
	// 	time.Sleep(1 * time.Second) // Delay to avoid hitting rate limits
	// }

	return nil
}


// func moveUntrackedPRs(ctx context.Context, client *github.Client, owner, repo string, columnID int64) error {
// 	prs, _, err := client.PullRequests.List(ctx, owner, repo, &github.PullRequestListOptions{State: "open"})
// 	if err != nil {
// 		return err
// 	}

// 	for _, pr := range prs {
// 		if err := moveItemIfUntracked(ctx, client, pr.GetID(), columnID); err != nil {
// 			fmt.Printf("Error moving PR #%d: %v\n", pr.GetNumber(), err)
// 		} else {
// 			fmt.Printf("Moved PR #%d to the project\n", pr.GetNumber())
// 		}
// 		time.Sleep(1 * time.Second) // Delay to avoid hitting rate limits
// 	}

// 	return nil
// }

// func moveItemIfUntracked(ctx context.Context, client *github.Client, contentID, columnID int64) error {
// 	cards, _, err := client.Projects.ListProjectCards(ctx, columnID, &github.ProjectCardListOptions{})
// 	if err != nil {
// 		return fmt.Errorf("error checking if item is tracked: %v", err)
// 	}

// 	for _, card := range cards {
// 		cardContentID, err := getContentIDFromURL(card.GetContentURL())
// 		if err != nil {
// 			fmt.Printf("Warning: Could not parse content ID for a card: %v\n", err)
// 			continue
// 		}
// 		if cardContentID == contentID {
// 			return nil // Item is already tracked, skip it
// 		}
// 	}

// 	_, _, err = client.Projects.CreateProjectCard(ctx, columnID, &github.ProjectCardOptions{
// 		ContentID:   contentID,
// 		ContentType: "Issue",
// 	})
// 	if err != nil {
// 		return fmt.Errorf("error adding item to project: %v", err)
// 	}

// 	return nil
// }