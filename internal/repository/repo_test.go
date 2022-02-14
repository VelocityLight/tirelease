package repository

import (
	"testing"

	"tirelease/commons/database"
	"tirelease/internal/entity"

	"github.com/stretchr/testify/assert"
)

func TestRepo(t *testing.T) {
	// Init
	var config = generateConfig()
	database.Connect(config)

	// Create
	var repo = &entity.Repo{
		Owner:    "Velocity",
		Repo:     "tirelease",
		FullName: "Velocity/tirelease",
	}
	err := CreateRepo(repo)
	// Assert
	assert.Equal(t, true, err == nil)

	// Update
	var desc = "This is a test repo"
	repo.Description = &desc
	err = UpdateRepo(repo)
	// Assert
	assert.Equal(t, true, err == nil)
}

/**
sql:

INSERT INTO repo (created_at, updated_at, owner, repo, full_name, html_url, description) VALUES (Now(),Now(),'pingcap','tiflow','pingcap/tiflow', 'https://github.com/pingcap/tiflow', 'tiflow源码库');
INSERT INTO repo (created_at, updated_at, owner, repo, full_name, html_url, description) VALUES (Now(),Now(),'pingcap','tidb','pingcap/tidb', 'https://github.com/pingcap/tidb', 'tidb源码库');
INSERT INTO repo (created_at, updated_at, owner, repo, full_name, html_url, description) VALUES (Now(),Now(),'tikv','tikv','tikv/tikv', 'https://github.com/tikv/tikv', 'tikv源码库');

**/
