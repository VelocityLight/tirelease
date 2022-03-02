package service

import (
	"strings"
	"time"

	"tirelease/commons/git"
	"tirelease/internal/entity"

	"github.com/google/go-github/v41/github"
	"github.com/shurcooL/githubv4"
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
	var page = 1
	var pageSize = 100
	option := &github.IssueListByRepoOptions{
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: pageSize,
		},
		Sort: "updated",
	}
	if nil != time {
		option.Since = *time
	}

	issues := make([]*entity.Issue, 0)
	for {
		gitIssues, _, err := git.Client.GetIssuesByOption(owner, repo, option)
		if nil != err {
			return nil, err
		}
		for _, issue := range gitIssues {
			if nil == issue.PullRequestLinks { // V3 considers every pull request an issue, so api return both issues and pull requests in the response
				issues = append(issues, ConsistIssueFromV3(issue))
			}
		}
		page++
		option.ListOptions.Page = page

		if len(gitIssues) < pageSize {
			break
		}
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

// ConsistIssueFromV4
// TODO: v4 implement by tony at 2022/02/14
func ConsistIssueFromV4(issueFiled *git.IssueField) *entity.Issue {
	labels := &[]github.Label{}
	for _, labelNode := range issueFiled.Labels.Nodes {
		label := github.Label{
			Name: github.String(string(labelNode.Name)),
		}
		*labels = append(*labels, label)
	}
	assignees := &[]github.User{}
	for _, userNode := range issueFiled.Assignees.Nodes {
		user := github.User{
			Login:     (*string)(&userNode.Login),
			CreatedAt: (*github.Timestamp)(&userNode.CreatedAt),
		}
		*assignees = append(*assignees, user)
	}
	closedByPrID := ""
	if issueFiled.State == githubv4.IssueStateClosed {
		for _, edge := range issueFiled.TimelineItems.Edges {
			closer := edge.Node.ClosedEvent.Closer.PullRequest
			if closer.Number != 0 {
				closedByPrID = closer.ID.(string)
			}
		}
	}

	return &entity.Issue{
		IssueID: issueFiled.ID.(string),
		Number:  int(issueFiled.Number),
		State:   string(issueFiled.State),
		Title:   string(issueFiled.Title),
		Owner:   string(issueFiled.Repository.Owner.Login),
		Repo:    string(issueFiled.Repository.Name),
		HTMLURL: string(issueFiled.Url),

		CreatedAt: issueFiled.CreatedAt.Time,
		UpdatedAt: issueFiled.UpdatedAt.Time,

		Labels:    labels,
		Assignees: assignees,

		ClosedByPullRequestID: closedByPrID,
	}
}
