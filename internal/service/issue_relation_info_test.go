package service

import (
	"testing"

	"tirelease/commons/git"

	"github.com/stretchr/testify/assert"
)

func TestConsistIssuePrRelationsByIssue(t *testing.T) {
	git.Connect(git.TestToken)
	git.ConnectV4(git.TestToken)

	triageRelationInfo, err := GetIssueRelationInfoByIssueNumber(git.TestOwner, git.TestRepo, git.TestIssueId)

	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(triageRelationInfo.IssuePrRelations) > 0)
	assert.Equal(t, true, len(triageRelationInfo.PullRequests) > 0)
	assert.Equal(t, true, len(triageRelationInfo.IssuePrRelations) == len(triageRelationInfo.PullRequests))
}
