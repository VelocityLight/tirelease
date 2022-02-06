package git

import (
	"context"

	"github.com/google/go-github/v41/github"
)

// ============================================================================ Repository
func (client *GithubInfo) GetRepositories() ([]*github.Repository, *github.Response, error) {
	return client.Client.Repositories.List(context.Background(), "", nil)
}

// ============================================================================ Issue
func (client *GithubInfo) GetIssueByNumber(owner, name string, number int) (*github.Issue, *github.Response, error) {
	return client.Client.Issues.Get(context.Background(), owner, name, number)
}

func (client *GithubInfo) GetIssueCommentsByIssueNumber(owner, name string, number int) ([]*github.IssueComment, *github.Response, error) {
	return client.Client.Issues.ListComments(context.Background(), owner, name, number, &github.IssueListCommentsOptions{})
}

func (client *GithubInfo) GetIssueTimelinesByIssueNumber(owner, name string, number int) ([]*github.Timeline, *github.Response, error) {
	return client.Client.Issues.ListIssueTimeline(context.Background(), owner, name, number, nil)
}

func (client *GithubInfo) GetIssueEventsByIssueNumber(owner, name string, number int) ([]*github.IssueEvent, *github.Response, error) {
	return client.Client.Issues.ListIssueEvents(context.Background(), owner, name, number, nil)
}

// ============================================================================ PullRequest
func (client *GithubInfo) GetPullRequestByNumber(owner, name string, number int) (*github.PullRequest, *github.Response, error) {
	return client.Client.PullRequests.Get(context.Background(), owner, name, number)
}
