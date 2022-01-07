package entity

import (
	"time"
)

type VersionTriage struct {
	ID           int64               `json:"id,omitempty"`
	VersionName  string              `json:"version_name,omitempty"`
	IssueID      int64               `json:"issue_id,omitempty"`
	TriageResult VersionTriageResult `json:"triage_result,omitempty"`

	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	DueTime    time.Time `json:"due_time,omitempty"`
	Comment    string    `json:"comment,omitempty"`
}

// Enum type
type VersionTriageResult string

const (
	VersionTriageResultUnKnown  = VersionTriageResult("UnKnown")
	VersionTriageResultAccept   = VersionTriageResult("Accept")
	VersionTriageResultLater    = VersionTriageResult("Later")
	VersionTriageResultWontFix  = VersionTriageResult("Won't Fix")
	VersionTriageResultReleased = VersionTriageResult("Released")
)

// List Option
type VersionTriageOption struct {
	ID           int64               `json:"id"`
	VersionName  string              `json:"version_name,omitempty"`
	IssueID      int64               `json:"issue_id,omitempty"`
	TriageResult VersionTriageResult `json:"triage_result,omitempty"`
}

// DB-Table
func (VersionTriage) TableName() string {
	return "version_triage"
}

/**

CREATE TABLE IF NOT EXISTS version_triage (
	id INT(11) NOT NULL COMMENT '全局ID',
	version_name VARCHAR(255) NOT NULL COMMENT '版本号',
	issue_id INT(11) COMMENT 'IssueID',
	triage_result VARCHAR(32) COMMENT 'Triage状态',

	create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	due_time TIMESTAMP COMMENT '延期时间',
	comment VARCHAR(1024) COMMENT '评论',

	PRIMARY KEY (id),
	UNIQUE KEY uk_versionname_issueid (version_name, issue_id),
	INDEX idx_issueid (issue_id)
)
ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT '版本Triage信息表';

**/
