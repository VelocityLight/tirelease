package dto

import (
	"tirelease/internal/entity"
)

// IssueRelationInfo Query Struct
type IssueRelationInfoQuery struct {
	// Issue
	ID      int64  `json:"id,omitempty"`
	IssueID string `json:"issue_id,omitempty"`
	Number  int    `json:"number,omitempty"`
	State   string `json:"state,omitempty"`
	Owner   string `json:"owner,omitempty"`
	Repo    string `json:"repo,omitempty"`

	SeverityLabel string `json:"severity_label,omitempty"`
	TypeLabel     string `json:"type_label,omitempty"`

	// Filter Option
	AffectVersion string `json:"affect_version,omitempty"`
	BaseBranch    string `json:"base_branch,omitempty"`
}

// IssueRelationInfo ReturnBack Struct
type IssueRelationInfo struct {
	Issue            *entity.Issue
	IssueAffects     *[]entity.IssueAffect
	IssuePrRelations *[]entity.IssuePrRelation
	PullRequests     *[]entity.PullRequest
}
