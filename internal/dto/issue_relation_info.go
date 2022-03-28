package dto

import (
	"time"
	"tirelease/internal/entity"
)

// IssueRelationInfo Query Struct
type IssueRelationInfoQuery struct {
	// Issue
	entity.IssueOption

	// Filter Option
	AffectVersion string                    `json:"affect_version,omitempty" form:"affect_version" uri:"affect_version"`
	AffectResult  entity.AffectResultResult `json:"affect_result,omitempty" form:"affect_result" uri:"affect_result"`
	BaseBranch    string                    `json:"base_branch,omitempty" form:"base_branch" uri:"base_branch"`
}

// IssueRelationInfo ReturnBack Struct
type IssueRelationInfo struct {
	Issue            *entity.Issue
	IssueAffects     *[]entity.IssueAffect
	IssuePrRelations *[]entity.IssuePrRelation
	PullRequests     *[]entity.PullRequest
	VersionTriages   *[]entity.VersionTriage
}

// Join IssueRelationInfo
type IssueRelationInfoByJoin struct {
	// issue
	entity.Issue

	// issue_affect
	IssueAffectJoin
}

type IssueAffectJoin struct {
	AffectID      int64                     `json:"affect_id,omitempty"`
	CreateTime    time.Time                 `json:"create_time"`
	UpdateTime    time.Time                 `json:"update_time"`
	AffectVersion string                    `json:"affect_version,omitempty"`
	AffectResult  entity.AffectResultResult `json:"affect_result,omitempty"`
}

func ComposeIssueAffectFromJoin(join *IssueRelationInfoByJoin) *entity.IssueAffect {
	return &entity.IssueAffect{
		ID:         join.IssueAffectJoin.AffectID,
		CreateTime: join.IssueAffectJoin.CreateTime,
		UpdateTime: join.IssueAffectJoin.UpdateTime,

		IssueID:       join.Issue.IssueID,
		AffectVersion: join.IssueAffectJoin.AffectVersion,
		AffectResult:  join.IssueAffectJoin.AffectResult,
	}
}
