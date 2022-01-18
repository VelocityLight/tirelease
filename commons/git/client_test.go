package git

import (
	"context"

	"testing"

	"github.com/google/go-github/v41/github"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	// Connect
	Connect("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")

	// List all repositories for the authenticated user
	repos, _, err := Client.Client.Repositories.List(context.Background(), "", nil)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(repos) > 0)
}


func TestGetIssue(t *testing.T) {
	// Connect
	Connect("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")

	// List comments
	issue, _, err := Client.Client.Issues.Get(context.Background(), "pingcap", "tidb", 28078)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, issue != nil)
}

func TestGetIssueComments(t *testing.T) {
	// Connect
	Connect("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")

	// List comments
	comments, _, err := Client.Client.Issues.ListComments(context.Background(), "pingcap", "tidb", 28078, &github.IssueListCommentsOptions{})

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(comments) > 0)
}

func TestGetIssueTimelines(t *testing.T) {
	// Connect
	Connect("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")

	// List comments
	reactions, _, err := Client.Client.Issues.ListIssueTimeline(context.Background(), "pingcap", "tidb", 28078, &github.ListOptions{})

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(reactions) > 0)
}

func TestGetPullRequest(t *testing.T) {
	// Connect
	Connect("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")

	// List comments
	pullRequest, _, err := Client.Client.PullRequests.Get(context.Background(), "pingcap", "tidb", 31345)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, pullRequest != nil)
}

