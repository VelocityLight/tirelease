package repository

import (
	"testing"

	"tirelease/commons/database"
	"tirelease/commons/git"
	"tirelease/internal/dto"

	"github.com/stretchr/testify/assert"
)

func TestSelectIssueRelationInfo(t *testing.T) {
	// Init
	var config = generateConfig()
	database.Connect(config)

	// Select
	option := &dto.IssueRelationInfoQuery{
		Owner: git.TestOwner,
		Repo:  git.TestRepo,
	}
	issueRelationInfos, err := SelectIssueRelationInfo(option)
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(*issueRelationInfos) > 0)
}
