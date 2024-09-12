package pkg

import "github.com/google/go-github/github"


// Provider is the interface to the back end data source
type Provider interface {
	GetProjectByName(name, org, user, repo string) *github.Project
	GetColumnID (projectName, columnName, org, user, repo string) (int64, error)
	ListColumnsForProject(projectName, org, user, repo string) ([]*github.ProjectColumn, error)
	ListIssues(query string, opts github.SearchOptions) (*github.IssuesSearchResult, error)
	ListIssuesForProjectColumn(columnID int64) ([]*github.Issue, error)
	ListProjectsForOrg(orgName string, opts github.ProjectListOptions) ([]*github.Project, *github.Response, error)
	ListProjectsForRepo(repoName string, opts github.ProjectListOptions) ([]*github.Project, *github.Response, error)
	ListProjectsForUser(userName string, opts github.ProjectListOptions) ([]*github.Project, *github.Response, error)
	ListRepos(query string, opts github.SearchOptions) (*github.RepositoriesSearchResult, error)
	MoveUntrackedIssues() error
}
