package dto

import (
	"tirelease/internal/entity"
)

// VersionTriage Query Struct
type VersionTriageInfoQuery struct {
	ID           int64                      `json:"id" form:"id"`
	VersionName  string                     `json:"version_name,omitempty" form:"version_name"`
	IssueID      string                     `json:"issue_id,omitempty" form:"issue_id"`
	TriageResult entity.VersionTriageResult `json:"triage_result,omitempty" form:"triage_result"`
}

// VersionTriage ReturnBack Struct
type VersionTriageInfo struct {
	ReleaseVersion *entity.ReleaseVersion `json:"release_version,omitempty"`
	IsFrozen       bool                   `json:"is_frozen,omitempty"`
	IsAccept       bool                   `json:"is_accept,omitempty"`

	VersionTriage     *entity.VersionTriage `json:"version_triage,omitempty"`
	IssueRelationInfo *IssueRelationInfo    `json:"issue_relation_info,omitempty"`
}

type VersionTriageInfoWrap struct {
	ReleaseVersion     *entity.ReleaseVersion `json:"release_version,omitempty"`
	VersionTriageInfos *[]VersionTriageInfo   `json:"version_triage_infos,omitempty"`
}
