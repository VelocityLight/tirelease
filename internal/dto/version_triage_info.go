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
	ReleaseVersion *entity.ReleaseVersion `json:"release_version,omitempty"`

	VersionTriage     *entity.VersionTriage `json:"version_triage,omitempty"`
	IsFrozen          bool                  `json:"is_frozen,omitempty"`
	IsAccept          bool                  `json:"is_accept,omitempty"`
	IssueRelationInfo *IssueRelationInfo    `json:"issue_relation_info,omitempty"`
}

type VersionTriageInfoWrap struct {
	ReleaseVersion     *entity.ReleaseVersion `json:"release_version,omitempty"`
	VersionTriageInfos *[]VersionTriageInfo   `json:"version_triage_infos,omitempty"`
}
