package dto

import (
	"tirelease/internal/entity"
)

// IssueRelationInfo Query Struct
type IssueRelationInfoQuery struct {
	// Issue
	ID      int64  `form:"id,omitempty"`
	IssueID string `form:"issue_id,omitempty"`
	Number  int    `form:"number,omitempty"`
	State   string `form:"state,omitempty"`
	Owner   string `form:"owner,omitempty"`
	Repo    string `form:"repo,omitempty"`

	SeverityLabel string `form:"severity_label,omitempty"`
	TypeLabel     string `form:"type_label,omitempty"`

	// Filter Option
	AffectVersion string `form:"affect_version,omitempty"`
	BaseBranch    string `form:"base_branch,omitempty"`
}

// IssueRelationInfo ReturnBack Struct
type IssueRelationInfo struct {
	Issue            *entity.Issue
	IssueAffects     *[]entity.IssueAffect
	IssuePrRelations *[]entity.IssuePrRelation
	PullRequests     *[]entity.PullRequest
}
