package entity

import (
	"time"

	"github.com/google/go-github/v41/github"
)

// Struct of Pull Request
type PullRequest struct {
	// DataBase columns
	ID            int64  `json:"id,omitempty"`
	PullRequestID string `json:"pull_request_id,omitempty"`
	Number        int    `json:"number,omitempty"`
	State         string `json:"state,omitempty"`
	Title         string `json:"title,omitempty"`
	Repo          string `json:"repo,omitempty"`
	HTMLURL       string `json:"html_url,omitempty"`
	HeadBranch    string `json:"head_branch,omitempty"`

	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	ClosedAt  *time.Time `json:"closed_at,omitempty"`
	MergedAt  *time.Time `json:"merged_at,omitempty"`

	Merged         bool   `json:"merged,omitempty"`
	Mergeable      bool   `json:"mergeable,omitempty"`
	MergeableState string `json:"mergeable_state,omitempty"`

	SourcePullRequestID string `json:"source_pull_request_id,omitempty"`

	LabelsString             string `json:"labels_string,omitempty"`
	AssigneeString           string `json:"assignee_string,omitempty"`
	AssigneesString          string `json:"assignees_string,omitempty"`
	RequestedReviewersString string `json:"requested_reviewers_string,omitempty"`

	// OutPut-Serial
	Labels             *[]github.Label `json:"labels,omitempty" gorm:"-"`
	Assignee           *github.User    `json:"assignee,omitempty" gorm:"-"`
	Assignees          *[]github.User  `json:"assignees,omitempty" gorm:"-"`
	RequestedReviewers *[]github.User  `json:"requested_reviewers,omitempty" gorm:"-"`
}

// List Option
type PullRequestOption struct {
	ID                  int64  `json:"id"`
	PullRequestID       string `json:"pull_request_id,omitempty"`
	Number              int    `json:"number,omitempty"`
	State               string `json:"state,omitempty"`
	Repo                string `json:"repo,omitempty"`
	HeadBranch          string `json:"head_branch,omitempty"`
	SourcePullRequestID string `json:"source_pull_request_id,omitempty"`
}

// DB-Table
func (PullRequest) TableName() string {
	return "pull_request"
}

/**

CREATE TABLE IF NOT EXISTS pull_request (
	id INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
	pull_request_id VARCHAR(255) COMMENT 'Pr全局ID',
	number INT(11) NOT NULL COMMENT '当前库ID',
	state VARCHAR(32) NOT NULL COMMENT '状态',
	title VARCHAR(1024) COMMENT '标题',
	repo VARCHAR(255) COMMENT '仓库',
	html_url VARCHAR(1024) COMMENT '链接',
	head_branch VARCHAR(255) COMMENT '链接',

	closed_at TIMESTAMP COMMENT '关闭时间',
	created_at TIMESTAMP COMMENT '创建时间',
	updated_at TIMESTAMP COMMENT '更新时间',
	merged_at TIMESTAMP COMMENT '合入时间',

	merged BOOLEAN COMMENT '是否已合入',
	mergeable BOOLEAN COMMENT '是否可合入',
	mergeable_state VARCHAR(32) COMMENT '可合入状态',

	source_pull_request_id VARCHAR(255) COMMENT '来源ID',

	labels_string TEXT COMMENT '标签',
	assignee_string TEXT COMMENT '处理人',
	assignees_string TEXT COMMENT '处理人列表',
	requested_reviewers_string TEXT COMMENT '处理人列表',

	PRIMARY KEY (id),
	UNIQUE KEY uk_prid (pull_request_id),
	INDEX idx_state (state),
	INDEX idx_repo (repo),
	INDEX idx_createdat (created_at),
	INDEX idx_sourceprid (source_pull_request_id)
)
ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT 'pull_request信息表';

**/
