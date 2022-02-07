package dto

import (
	"tirelease/internal/entity"
)

// IssueRelationInfo Struct
type IssueRelationInfo struct {
	Issue            *entity.Issue
	IssueAffects     []*entity.IssueAffect
	IssuePrRelations []*entity.IssuePrRelation
	PullRequests     []*entity.PullRequest
}
