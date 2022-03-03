package dto

import (
	"tirelease/internal/entity"
)

// VersionTriageInfo Query Struct
type VersionTriageInfoQuery struct {
	VersionName string `json:"version_name,omitempty"`
	IssueID     string `json:"issue_id,omitempty"`
}

// VersionTriage ReturnBack Struct
type VersionTriageInfo struct {
	VersionTriage     *entity.VersionTriage  `json:"version_triage,omitempty"`
	ReleaseVersion    *entity.ReleaseVersion `json:"release_version,omitempty"`
	IssueRelationInfo *IssueRelationInfo     `json:"issue_relation_info,omitempty"`
}

type VersionTriageInfoWrap struct {
	ReleaseVersion     *entity.ReleaseVersion `json:"release_version,omitempty"`
	VersionTriageInfos *[]VersionTriageInfo   `json:"version_triage_infos,omitempty"`
}
