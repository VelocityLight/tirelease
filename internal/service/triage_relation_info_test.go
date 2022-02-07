package service

import (
	"testing"

	"tirelease/commons/git"

	"github.com/stretchr/testify/assert"
)

func TestConsistIssuePrRelationsByIssue(t *testing.T) {
	git.Connect(git.TestToken)
	git.ConnectV4(git.TestToken)

	issuePrRelations, pullRequests, err := ConsistIssuePrRelationsByIssue(git.TestOwner, git.TestRepo, git.TestIssueId)

	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(issuePrRelations) > 0)
	assert.Equal(t, true, len(pullRequests) > 0)
}
