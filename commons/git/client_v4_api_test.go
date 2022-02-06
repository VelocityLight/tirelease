package git

import (
	"context"
	"testing"

	"github.com/shurcooL/githubv4"
	"github.com/stretchr/testify/assert"
)

//=======================================================================Login
func TestGetLoginV4(t *testing.T) {
	// Connect
	ConnectV4(testToken)

	// Query
	var query struct {
		Viewer struct {
			Login     githubv4.String
			CreatedAt githubv4.DateTime
		}
	}
	err := ClientV4.client.Query(context.Background(), &query, nil)

	// Assert
	assert.Equal(t, true, err == nil)
}

//=======================================================================Issue
func TestGetIssueV4(t *testing.T) {
	// Connect
	ConnectV4(testToken)

	// Query
	issue, err := ClientV4.GetIssueByNumber(testOwner, testRepo, testIssueId)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, issue != nil)
}

//=======================================================================Pr
func TestGetPullRequestV4(t *testing.T) {
	// Connect
	ConnectV4(testToken)

	// Query
	pr, err := ClientV4.GetPullRequestsByNumber(testOwner, testRepo, testPullRequestId)

	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, pr != nil)
}
