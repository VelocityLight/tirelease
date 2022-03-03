package service

import (
	"testing"

	"tirelease/commons/database"
	"tirelease/commons/git"
	"tirelease/internal/repository"

	"github.com/stretchr/testify/assert"
)

func TestComposeIssuePrRelationsByIssue(t *testing.T) {
	// Init
	git.Connect(git.TestToken)
	git.ConnectV4(git.TestToken)
	database.Connect(generateConfig())

	// Test
	triageRelationInfo, err := GetIssueRelationInfoByIssueNumber(git.TestOwner, git.TestRepo, git.TestIssueId)
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(*triageRelationInfo.IssuePrRelations) > 0)
	assert.Equal(t, true, len(*triageRelationInfo.PullRequests) > 0)
	assert.Equal(t, true, len(*triageRelationInfo.IssuePrRelations) == len(*triageRelationInfo.PullRequests))

	// Save (If Needed)
	err = repository.SaveIssueRelationInfo(triageRelationInfo)
	assert.Equal(t, true, err == nil)
}
