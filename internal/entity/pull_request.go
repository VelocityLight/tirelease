package entity

import (
	"github.com/shurcooL/githubv4"
)

// Struct of Pull Request
type PullRequest struct {
	ID         int64  `json:"id,omitempty"`
	Number     int    `json:"number,omitempty"`
	State      string `json:"state,omitempty"`
	Title      string `json:"title,omitempty"`
	Repo       string `json:"repo,omitempty"`
	HTMLURL    string `json:"html_url,omitempty"`
	HeadBranch string `json:"head_branch,omitempty"`

	CreatedAt githubv4.DateTime `json:"created_at,omitempty"`
	UpdatedAt githubv4.DateTime `json:"updated_at,omitempty"`
	ClosedAt  githubv4.DateTime `json:"closed_at,omitempty"`
	MergedAt  githubv4.DateTime `json:"merged_at,omitempty"`

	Merged         bool   `json:"merged,omitempty"`
	Mergeable      bool   `json:"mergeable,omitempty"`
	MergeableState string `json:"mergeable_state,omitempty"`

	LabelsString             string `json:"labels_string,omitempty"`
	AssigneeString           string `json:"assignee_string,omitempty"`
	AssigneesString          string `json:"assignees_string,omitempty"`
	RequestedReviewersString string `json:"requested_reviewers_string,omitempty"`
}

// DB-Table
func (PullRequest) TableName() string {
	return "pull_request"
}

/**
CREATE TABLE IF NOT EXISTS pull_request (
	id INT(11) NOT NULL COMMENT '全局ID',
	number INT(11) NOT NULL COMMENT '当前库ID',
	state VARCHAR(32) NOT NULL COMMENT '状态',
	title VARCHAR(1024) COMMENT '标题',
	repo VARCHAR(255) COMMENT '仓库',
	html_url VARCHAR(1024) COMMENT '链接',
	head_branch VARCHAR(255) COMMENT '链接',

	close_at TIMESTAMP COMMENT '关闭时间',
	create_at TIMESTAMP COMMENT '创建时间',
	update_at TIMESTAMP COMMENT '更新时间',
	merged_at TIMESTAMP COMMENT '更新时间',

	merged BOOLEAN COMMENT '是否已合入',
	mergeable BOOLEAN COMMENT '是否可合入',
	mergeable_state BOOLEAN COMMENT '可合入状态',

	labels_string TEXT COMMENT '标签',
	assignee_string TEXT COMMENT '处理人',
	assignees_string TEXT COMMENT '处理人列表',
	requested_reviewers_string TEXT COMMENT '处理人列表',

	PRIMARY KEY (id),
	INDEX idx_state (state),
	INDEX idx_repo (repo),
	INDEX idx_createat (create_at)
)
ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT 'pull_request信息表';
**/
