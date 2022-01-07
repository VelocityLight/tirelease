package repository

import (
	"testing"
	"time"

	"tirelease/internal/entity"
	"tirelease/commons/configs"
	"tirelease/commons/database"

	"github.com/stretchr/testify/assert"
	"github.com/google/go-github/v41/github"
)

func TestCreateIssue(t *testing.T) {
	// Init
	var config = generateConfig()
	database.Connect(config)

	// Create
	var issue = &entity.Issue{
		ID: 100,
		Number: 100,
		State: "open",
		Title: "first",
		Repo: "ff",
		HTMLURL: "json",
		ClosedAt: time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	issue.Assignee = &github.User{Login: github.String("jcye")}
	err := CreateOrUpdateIssue(issue)
	// Assert
	assert.Equal(t, true, err == nil)

	// Select
	var option = &entity.IssueOption{
		ID: 100,
	}
	issues, err := SelectIssue(option)
	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(*issues) > 0)
}

func generateConfig() *configs.ConfigYaml {
	var config = &configs.ConfigYaml{}

	config.Mysql.UserName = "cicd_online"
	config.Mysql.PassWord = "wGEXq8a4MeCw6G"
	config.Mysql.Host = "172.16.4.36"
	config.Mysql.Port = "3306"
	config.Mysql.DataBase = "cicd_online"
	config.Mysql.CharSet = "utf8"
	
	return config
}
