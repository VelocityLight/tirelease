package repository

import (
	"testing"
	"time"

	"tirelease/commons/database"
	"tirelease/internal/entity"

	"github.com/stretchr/testify/assert"
)

func TestIssueAffect(t *testing.T) {
	// Init
	var config = generateConfig()
	database.Connect(config)

	// Create
	var issueAffect = &entity.IssueAffect{
		ID:         100,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),

		AffectVersion: "5.4.1",
		IssueID:       100,
		AffectResult:  entity.AffectResultResultUnKnown,
	}
	err := CreateIssueAffect(issueAffect)
	// Assert
	assert.Equal(t, true, err == nil)

	// Update
	issueAffect.AffectResult = entity.AffectResultResultYes
	err = UpdateIssueAffect(issueAffect)
	// Assert
	assert.Equal(t, true, err == nil)

	// Select
	var option = &entity.IssueAffectOption{
		AffectVersion: "5.4.1",
	}
	issueAffects, err := SelectIssueAffect(option)
	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(*issueAffects) > 0)
}
