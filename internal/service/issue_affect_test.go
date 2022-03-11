package service

import (
	"testing"

	"tirelease/commons/database"
	"tirelease/commons/git"

	// "tirelease/internal/entity"

	"github.com/stretchr/testify/assert"
)

func TestComposeIssueAffectWithIssueID(t *testing.T) {
	// Init
	database.Connect(generateConfig())

	// Insert
	// issueAffect := &entity.IssueAffect{
	// 	IssueID: git.TestIssueNodeID,
	// 	AffectVersion: "5.4",
	// 	AffectResult: entity.AffectResultResultYes,
	// }
	// err := UpdateIssueAffect(issueAffect)
	// assert.Equal(t, true, err == nil)

	// Test
	issueAffects, err := ComposeIssueAffectWithIssueID(git.TestIssueNodeID)
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(*issueAffects) > 0)
}
