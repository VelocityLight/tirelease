package repository

import (
	"testing"
	"time"

	"tirelease/commons/database"
	"tirelease/internal/entity"

	"github.com/google/go-github/v41/github"
	"github.com/stretchr/testify/assert"
)

func TestPullRequest(t *testing.T) {
	// Init
	var config = generateConfig()
	database.Connect(config)

	// Create
	var pr = &entity.PullRequest{
		PullRequestID: "100",
		Number:        100,
		State:         "open",
		Title:         "first",
		Repo:          "ff",
		HTMLURL:       "json",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		HeadBranch:    "targetBranch",

		Merged:         true,
		Mergeable:      true,
		MergeableState: "OK",

		SourcePullRequestID: "1000",
	}
	pr.Assignee = &github.User{Login: github.String("jcye")}
	err := CreateOrUpdatePullRequest(pr)
	// Assert
	assert.Equal(t, true, err == nil)

	// Select
	var option = &entity.PullRequestOption{
		PullRequestID: "100",
	}
	prs, err := SelectPullRequest(option)
	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(*prs) > 0)
}
