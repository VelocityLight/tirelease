package git

import (
	"testing"

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

	// List comments

	issue, _, err := Client.GetIssueByNumber(TestOwner, TestRepo, TestIssueId)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, issue != nil)
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
