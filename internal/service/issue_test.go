package service

import (
	"testing"

	"tirelease/commons/git"

	"github.com/stretchr/testify/assert"
)

func TestGetIssueByNumberFromV3(t *testing.T) {
	git.Connect(git.TestToken)
	git.ConnectV4(git.TestToken)

	issue, err := GetIssueByNumberFromV3(git.TestOwner, git.TestRepo, git.TestIssueId)

	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, issue != nil)
}
