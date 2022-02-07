package git

import (
	"testing"
	"time"

	"github.com/google/go-github/v41/github"
	"github.com/stretchr/testify/assert"
)

//=======================================================================Repository
func TestGetRepository(t *testing.T) {
	// Connect
	Connect(TestToken)

	// List all repositories for the authenticated user
	repos, _, err := Client.GetRepositories()

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(repos) > 0)
}

//=======================================================================Issue
func TestGetIssue(t *testing.T) {
	// Connect
	Connect(TestToken)

	// Query
	issue, _, err := Client.GetIssueByNumber(TestOwner, TestRepo, TestIssueId)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, issue != nil)
}

func TestGetIssuesByTimeRange(t *testing.T) {
	// Connect
	Connect(TestToken)

	// Query
	day, _ := time.ParseDuration("-24h")
	option := &github.IssueListByRepoOptions{
		Since: time.Now().Add(15 * day),
	}
	issues, _, err := Client.GetIssuesByTimeRange(TestOwner, TestRepo, option)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(issues) > 0)
}

func TestGetIssueComments(t *testing.T) {
	// Connect
	Connect(TestToken)

	// List comments
	comments, _, err := Client.GetIssueCommentsByIssueNumber(TestOwner, TestRepo, TestIssueId)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(comments) > 0)
}

func TestGetIssueTimelines(t *testing.T) {
	// Connect
	Connect(TestToken)

	// List comments
	reactions, _, err := Client.GetIssueTimelinesByIssueNumber(TestOwner, TestRepo, TestIssueId)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(reactions) > 0)
}

func TestGetIssueEvents(t *testing.T) {
	// Connect
	Connect(TestToken)

	// List comments
	events, _, err := Client.GetIssueEventsByIssueNumber(TestOwner, TestRepo, TestIssueId)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(events) > 0)
}

//=======================================================================PullRequest
func TestGetPullRequest(t *testing.T) {
	// Connect
	Connect(TestToken)

	// List comments
	pullRequest, _, err := Client.GetPullRequestByNumber(TestOwner, TestRepo, TestPullRequestId)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, pullRequest != nil)
}
