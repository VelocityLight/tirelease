package service

import (
	"testing"

	"tirelease/commons/database"
	"tirelease/commons/git"

	"github.com/stretchr/testify/assert"
)

func TestWebhookRefreshPullRequestV3(t *testing.T) {
	// init
	git.Connect(git.TestToken)
	git.ConnectV4(git.TestToken)
	database.Connect(generateConfig())

	// detail
	pr, _, err := git.Client.GetPullRequestByNumber(git.TestOwner2, git.TestRepo2, git.TestPullRequestId2)
	assert.Equal(t, true, err == nil)
	err = WebhookRefreshPullRequestV3(pr)
	assert.Equal(t, true, err == nil)
}
