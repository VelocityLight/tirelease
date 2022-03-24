package dto

import (
	"tirelease/internal/entity"
)

// IssueRelationInfo Query Struct
type IssueRelationInfoQuery struct {
	// Issue
	entity.IssueOption

	// Filter Option
	AffectVersion string `json:"affect_version,omitempty" form:"affect_version" uri:"affect_version"`
	BaseBranch    string `json:"base_branch,omitempty" form:"base_branch" uri:"base_branch"`
}

// IssueRelationInfo ReturnBack Struct
type IssueRelationInfo struct {
	Issue            *entity.Issue
	IssueAffects     *[]entity.IssueAffect
	IssuePrRelations *[]entity.IssuePrRelation
	PullRequests     *[]entity.PullRequest
	VersionTriages   *[]entity.VersionTriage
}
