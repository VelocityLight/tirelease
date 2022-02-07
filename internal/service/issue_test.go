package service

import (
	"testing"
	"time"

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

func TestGetIssuesByTimeFromV3(t *testing.T) {
	git.Connect(git.TestToken)
	git.ConnectV4(git.TestToken)

	day, _ := time.ParseDuration("-24h")
	time := time.Now().Add(15 * day)
	issues, err := GetIssuesByTimeFromV3(git.TestOwner, git.TestRepo, &time)

	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(issues) > 0)
}
