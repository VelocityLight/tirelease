package service

import (
	"strings"
	"time"

	"tirelease/commons/git"
	"tirelease/internal/entity"

	"github.com/google/go-github/v41/github"
)

// GetIssueByNumberFromV3
func GetIssueByNumberFromV3(owner, repo string, number int) (*entity.Issue, error) {
	issue, _, err := git.Client.GetIssueByNumber(owner, repo, number)
	if nil != err {
		return nil, err
	}
	return ConsistIssueFromV3(issue), nil
}

// GetIssuesByTimeFromV3
func GetIssuesByTimeFromV3(owner, repo string, time *time.Time) ([]*entity.Issue, error) {
	option := &github.IssueListByRepoOptions{
		Since: *time,
	}
	gitIssues, _, err := git.Client.GetIssuesByTimeRange(owner, repo, option)
	if nil != err {
		return nil, err
	}

	issues := make([]*entity.Issue, 0)
	for _, issue := range gitIssues {
		issues = append(issues, ConsistIssueFromV3(issue))
	}
	return issues, nil
}

// ConsistIssueFromV3
func ConsistIssueFromV3(issue *github.Issue) *entity.Issue {
	labels := &[]github.Label{}
	for _, node := range issue.Labels {
		*labels = append(*labels, *node)
	}
	assignees := &[]github.User{}
	for _, node := range issue.Assignees {
		*assignees = append(*assignees, *node)
	}
	url := strings.Split(*issue.RepositoryURL, "/")
	owner := url[len(url)-2]
	repo := url[len(url)-1]

	return &entity.Issue{
		IssueID: *issue.NodeID,
		Number:  *issue.Number,
		State:   *issue.State,
		Title:   *issue.Title,
		Owner:   owner,
		Repo:    repo,
		HTMLURL: *issue.HTMLURL,

		CreatedAt: *issue.CreatedAt,
		UpdatedAt: *issue.UpdatedAt,
		ClosedAt:  issue.ClosedAt,

		Labels:    labels,
		Assignee:  issue.Assignee,
		Assignees: assignees,
	}
}
