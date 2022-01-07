package entity

import (
	"time"

	"github.com/google/go-github/v41/github"
)

// Struct of Issue
type Issue struct {
	// DataBase Column
	ID      int64  `json:"id,omitempty"`
	Number  int    `json:"number,omitempty"`
	State   string `json:"state,omitempty"`
	Title   string `json:"title,omitempty"`
	Repo    string `json:"repo,omitempty"`
	HTMLURL string `json:"html_url,omitempty"`

	ClosedAt  time.Time `json:"closed_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	LabelsString    string `json:"labels_string,omitempty"`
	AssigneeString  string `json:"assignee_string,omitempty"`
	AssigneesString string `json:"assignees_string,omitempty"`

	ClosedByPullRequestID int64 `json:"closed_by_pull_request_id,omitempty"`

	// OutPut-Serial
	Labels    *[]github.Label `json:"labels,omitempty" gorm:"-"`
	Assignee  *github.User    `json:"assignee,omitempty" gorm:"-"`
	Assignees *[]github.User  `json:"assignees,omitempty" gorm:"-"`
}

// List Option
type IssueOption struct {
	ID     int64  `json:"id"`
	Number int    `json:"number,omitempty"`
	State  string `json:"state,omitempty"`
	Repo   string `json:"repo,omitempty"`
}

// DB-Table
func (Issue) TableName() string {
	return "issue"
}

/**

CREATE TABLE IF NOT EXISTS issue (
	id INT(11) NOT NULL COMMENT '全局ID',
	number INT(11) NOT NULL COMMENT '当前库ID',
	state VARCHAR(32) NOT NULL COMMENT '状态',
	title VARCHAR(1024) COMMENT '标题',
	repo VARCHAR(255) COMMENT '仓库',
	html_url VARCHAR(1024) COMMENT '链接',

	closed_at TIMESTAMP COMMENT '关闭时间',
	created_at TIMESTAMP COMMENT '创建时间',
	updated_at TIMESTAMP COMMENT '更新时间',

	labels_string TEXT COMMENT '标签',
	assignee_string TEXT COMMENT '处理人',
	assignees_string TEXT COMMENT '处理人列表',

	closed_by_pull_request_id INT(11) COMMENT '处理的PR',

	PRIMARY KEY (id),
	INDEX idx_state (state),
	INDEX idx_repo (repo),
	INDEX idx_createdat (created_at)
)
ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT 'issue信息表';

**/
