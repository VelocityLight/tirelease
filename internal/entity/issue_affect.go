package entity

import (
	"time"
)

type IssueAffect struct {
	ID         int64     `json:"id,omitempty"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`

	IssueID       int64              `json:"issue_id,omitempty"`
	AffectVersion string             `json:"affect_version,omitempty"`
	AffectResult  AffectResultResult `json:"affect_result,omitempty"`
}

// Enum type
type AffectResultResult string

const (
	AffectResultResultUnKnown = AffectResultResult("UnKnown")
	AffectResultResultYes     = AffectResultResult("Yes")
	AffectResultResultNo      = AffectResultResult("No")
)

// List Option
type IssueAffectOption struct {
	ID            int64  `json:"id"`
	IssueID       int64  `json:"issue_id,omitempty"`
	AffectVersion string `json:"affect_version,omitempty"`
}

// DB-Table
func (IssueAffect) TableName() string {
	return "issue_affect"
}

/**

CREATE TABLE IF NOT EXISTS issue_affect (
	id INT(11) NOT NULL COMMENT '全局ID',
	create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',

	issue_id INT(11) COMMENT 'IssueID',
	affect_version VARCHAR(255) NOT NULL COMMENT '版本号',
	affect_result VARCHAR(32) COMMENT 'Triage状态',

	PRIMARY KEY (id),
	UNIQUE KEY uk_issueid_affectversion (issue_id, affect_version)
)
ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT 'Issue影响版本信息表';

**/
