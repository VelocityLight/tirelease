package git

import (
	"context"

	"testing"

	"github.com/google/go-github/v41/github"
	"github.com/stretchr/testify/assert"
)

var testToken string = "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"
var testIssueId int = 28078
var testPullRequestId int = 31287
var testOwner string = "pingcap"
var testRepo string = "tidb"

//=======================================================================Repository
func TestGetRepository(t *testing.T) {
	// Connect
	Connect(testToken)

	// List all repositories for the authenticated user
	repos, _, err := Client.Client.Repositories.List(context.Background(), "", nil)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(repos) > 0)
}

//=======================================================================Issue
func TestGetIssue(t *testing.T) {
	// Connect
	Connect(testToken)

	// List comments
	issue, _, err := Client.Client.Issues.Get(context.Background(), testOwner, testRepo, testIssueId)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, issue != nil)
}

func TestGetIssueComments(t *testing.T) {
	// Connect
	Connect(testToken)

	// List comments
	comments, _, err := Client.Client.Issues.ListComments(context.Background(), testOwner, testRepo, testIssueId, &github.IssueListCommentsOptions{})

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(comments) > 0)
}

func TestGetIssueTimelines(t *testing.T) {
	// Connect
	Connect(testToken)

	// List comments
	reactions, _, err := Client.Client.Issues.ListIssueTimeline(context.Background(), testOwner, testRepo, testIssueId, &github.ListOptions{})

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(reactions) > 0)
}

func TestGetIssueEvents(t *testing.T) {
	// Connect
	Connect(testToken)

	// List comments
	events, _, err := Client.Client.Issues.ListIssueEvents(context.Background(), testOwner, testRepo, testIssueId, &github.ListOptions{})

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(events) > 0)
}

//=======================================================================PullRequest
func TestGetPullRequest(t *testing.T) {
	// Connect
	Connect(testToken)

	// List comments
	pullRequest, _, err := Client.Client.PullRequests.Get(context.Background(), testOwner, testRepo, testPullRequestId)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, pullRequest != nil)
}
