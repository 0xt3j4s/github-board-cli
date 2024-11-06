package pkg

import (
	"context"
	// "fmt"
	"fmt"
	"strings"
	// "time"

	"github.com/google/go-github/v39/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
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
	fmt.Println("result ",result)
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
	fmt.Print("result: ",result)
	if err != nil {
		return nil, nil, err
	}

	return result, res, err
}

// GetProjectByName will return a single project given a name
func (c *Client) GetProjectByName(name, org, user, repo string) *github.Project {
	var projects []*github.Project
	fmt.Print("org: \n", org)
	fmt.Print("name: \n", name)
	fmt.Print("user: \n", user)
	fmt.Print("repo: \n", repo)

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
	fmt.Println("project info: ", project.OwnerURL)
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

	fmt.Print("columns: ", columns)
	for _, column := range columns {
		if *column.Name == columnName {
			return *column.ID, nil
		}
	}

	return 0, fmt.Errorf("column '%s' not found in the project", columnName)
}
