package service

import (
	"testing"

	"tirelease/commons/database"
	"tirelease/commons/git"
	"tirelease/internal/entity"

	"github.com/stretchr/testify/assert"
)

func TestCronRefreshIssuesV4(t *testing.T) {
	// init
	git.Connect(git.TestToken)
	git.ConnectV4(git.TestToken)
	database.Connect(generateConfig())
	repo := &entity.Repo{
		Owner: git.TestOwner2,
		Repo:  git.TestRepo2,
	}
	repos := []entity.Repo{*repo}

	// detail
	err := CronRefreshIssuesV4(&repos)
	assert.Equal(t, true, err == nil)
}

func TestWebhookRefreshIssueV4(t *testing.T) {
	// init
	git.Connect(git.TestToken)
	git.ConnectV4(git.TestToken)
	database.Connect(generateConfig())

	// detail
	issue, _, err := git.Client.GetIssueByNumber(git.TestOwner2, git.TestRepo2, git.TestIssueId2)
	assert.Equal(t, true, err == nil)
	err = WebhookRefreshIssueV4(issue)
	assert.Equal(t, true, err == nil)
}
