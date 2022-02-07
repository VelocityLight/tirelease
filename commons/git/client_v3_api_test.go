package git

import (
	"testing"

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
	repos, _, err := Client.GetRepositories()

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(repos) > 0)
}

//=======================================================================Issue
func TestGetIssue(t *testing.T) {
	// Connect
	Connect(testToken)

	// List comments

	issue, _, err := Client.GetIssueByNumber(testOwner, testRepo, testIssueId)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, issue != nil)
}

func TestGetIssueComments(t *testing.T) {
	// Connect
	Connect(testToken)

	// List comments
	comments, _, err := Client.GetIssueCommentsByIssueNumber(testOwner, testRepo, testIssueId)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(comments) > 0)
}

func TestGetIssueTimelines(t *testing.T) {
	// Connect
	Connect(testToken)

	// List comments
	reactions, _, err := Client.GetIssueTimelinesByIssueNumber(testOwner, testRepo, testIssueId)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(reactions) > 0)
}

func TestGetIssueEvents(t *testing.T) {
	// Connect
	Connect(testToken)

	// List comments
	events, _, err := Client.GetIssueEventsByIssueNumber(testOwner, testRepo, testIssueId)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(events) > 0)
}

//=======================================================================PullRequest
func TestGetPullRequest(t *testing.T) {
	// Connect
	Connect(testToken)

	// List comments
	pullRequest, _, err := Client.GetPullRequestByNumber(testOwner, testRepo, testPullRequestId)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, pullRequest != nil)
}
