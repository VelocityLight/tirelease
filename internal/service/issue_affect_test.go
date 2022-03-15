package service

import (
	"testing"

	"tirelease/commons/database"
	"tirelease/commons/git"
	"tirelease/internal/entity"

	"github.com/stretchr/testify/assert"
)

func TestComposeIssueAffectWithIssueID(t *testing.T) {
	// Init
	database.Connect(generateConfig())

	// Test
	issueAffects, err := ComposeIssueAffectWithIssueID(git.TestIssueNodeID, nil)
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(*issueAffects) > 0)
}

func TestCreateOrUpdateIssueAffect(t *testing.T) {
	// init
	git.Connect(git.TestToken)
	git.ConnectV4(git.TestToken)
	database.Connect(generateConfig())

	// Test
	issueAffect := &entity.IssueAffect{
		IssueID:       git.TestIssueNodeID2,
		AffectVersion: "5.4",
		AffectResult:  entity.AffectResultResultYes,
	}
	err := CreateOrUpdateIssueAffect(issueAffect)
	assert.Equal(t, true, err == nil)

	issueAffect.AffectResult = entity.AffectResultResultNo
	err = CreateOrUpdateIssueAffect(issueAffect)
	assert.Equal(t, true, err == nil)
}