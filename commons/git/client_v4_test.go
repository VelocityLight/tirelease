package git

import (
	"context"

	"github.com/shurcooL/githubv4"
	"github.com/stretchr/testify/assert"
	"testing"
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
	// ConnectV4(testToken)

	// // Query
	// var query struct {
	// 	Repository struct {
	// 		Issues struct {
	// 			Edges []struct {
	// 				Node   IssueNode
	// 			}
	// 		} `graphql:"issues(number: $number)"`
	// 	} `graphql:"repository(owner: $owner, name: $name)"`
	// }
	// param := map[string]interface{}{
	// 	"owner":  githubv4.String(testOwner),
	// 	"name":   githubv4.String(testRepo),
	// 	"number": githubv4.Int(testIssueId),
	// }
	// err := ClientV4.client.Query(context.Background(), &query, param)

	// // Assert
	// assert.Equal(t, true, err == nil)
}
