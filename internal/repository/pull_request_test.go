package repository

import (
	"testing"
	"time"

	"tirelease/commons/configs"
	"tirelease/commons/database"
	"tirelease/internal/entity"

	"github.com/google/go-github/v41/github"
	"github.com/stretchr/testify/assert"
)

func TestCreatePullRequest(t *testing.T) {
	// Init
	var config = generateConfig2()
	database.Connect(config)

	// Create
	var pr = &entity.PullRequest{
		ID:         100,
		Number:     100,
		State:      "open",
		Title:      "first",
		Repo:       "ff",
		HTMLURL:    "json",
		ClosedAt:   time.Now(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		MergedAt:   time.Now(),
		HeadBranch: "targetBranch",

		Merged:         true,
		Mergeable:      true,
		MergeableState: "OK",

		SourcePullRequestID: 1000,
	}
	pr.Assignee = &github.User{Login: github.String("jcye")}
	err := CreateOrUpdatePullRequest(pr)
	// Assert
	assert.Equal(t, true, err == nil)

	// Select
	var option = &entity.PullRequestOption{
		ID: 100,
	}
	prs, err := SelectPullRequest(option)
	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(*prs) > 0)
}

func generateConfig2() *configs.ConfigYaml {
	var config = &configs.ConfigYaml{}

	config.Mysql.UserName = "cicd_online"
	config.Mysql.PassWord = "wGEXq8a4MeCw6G"
	config.Mysql.Host = "172.16.4.36"
	config.Mysql.Port = "3306"
	config.Mysql.DataBase = "cicd_online"
	config.Mysql.CharSet = "utf8"

	return config
}
