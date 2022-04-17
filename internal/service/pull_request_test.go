package service

import (
	"testing"

	"tirelease/commons/git"

	"github.com/stretchr/testify/assert"
)

func TestGetPullRequestByNumberFromV3(t *testing.T) {
	git.Connect(git.TestToken)
	git.ConnectV4(git.TestToken)

	pr, err := GetPullRequestByNumberFromV3(git.TestOwner, git.TestRepo, git.TestPullRequestId)

	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, pr != nil)
}

func TestGetPullRequestRefIssuesByRegexFromV4(t *testing.T) {
	git.Connect(git.TestToken)
	git.ConnectV4(git.TestToken)

	pr, err := git.ClientV4.GetPullRequestByID(git.TestPullRequestNodeID)
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, pr != nil)

	issueNumbers, err := GetPullRequestRefIssuesByRegexFromV4(pr)
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(issueNumbers) > 0)
}
