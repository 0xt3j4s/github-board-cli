package pkg

import "github.com/google/go-github/v39/github"


// Provider is the interface to the back end data source
type Provider interface {
	GetProjectByName(name, org, user, repo string) *github.Project
	GetColumnID (projectName, columnName, org, user, repo string) (int64, error)
	ListColumnsForProject(projectName, org, user, repo string) ([]*github.ProjectColumn, error)
	ListProjectsForOrg(orgName string, opts github.ProjectListOptions) ([]*github.Project, *github.Response, error)
	ListProjectsForRepo(repoName string, opts github.ProjectListOptions) ([]*github.Project, *github.Response, error)
	ListProjectsForUser(userName string, opts github.ProjectListOptions) ([]*github.Project, *github.Response, error)
}
