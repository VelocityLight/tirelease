package service

import (
	"testing"

	"tirelease/commons/database"
	"tirelease/commons/git"

	"tirelease/internal/dto"

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
	err = SaveIssueRelationInfo(triageRelationInfo)
	assert.Equal(t, true, err == nil)
}

func TestSelectIssueRelationInfo(t *testing.T) {
	// Init
	var config = generateConfig()
	database.Connect(config)

	// Select
	option := &dto.IssueRelationInfoQuery{
		Owner: git.TestOwner,
		Repo:  git.TestRepo,
	}
	issueRelationInfos, err := SelectIssueRelationInfo(option)
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(*issueRelationInfos) > 0)
}
