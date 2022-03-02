package dto

import (
	"tirelease/internal/entity"
)

// IssueRelationInfo Query Struct
type IssueRelationInfoQuery struct {
	IssueID string `json:"issue_id,omitempty"`
	Number  int    `json:"number,omitempty"`
	State   string `json:"state,omitempty"`
	Repo    string `json:"repo,omitempty"`

	AffectVersion string `json:"affect_version,omitempty"`
}

// IssueRelationInfo ReturnBack Struct
type IssueRelationInfo struct {
	Issue            *entity.Issue
	IssueAffects     []*entity.IssueAffect
	IssuePrRelations []*entity.IssuePrRelation
	PullRequests     []*entity.PullRequest
}
