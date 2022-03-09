package entity

import (
	"strings"
	"time"

	"tirelease/commons/git"

	"github.com/google/go-github/v41/github"
	"github.com/shurcooL/githubv4"
)

// Struct of Issue
type Issue struct {
	// DataBase Column
	ID      int64  `json:"id,omitempty"`
	IssueID string `json:"issue_id,omitempty"`
	Number  int    `json:"number,omitempty"`
	State   string `json:"state,omitempty"`
	Title   string `json:"title,omitempty"`
	Owner   string `json:"owner,omitempty"`
	Repo    string `json:"repo,omitempty"`
	HTMLURL string `json:"html_url,omitempty"`

	ClosedAt  *time.Time `json:"closed_at,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`

	LabelsString    string `json:"labels_string,omitempty"`
	AssigneesString string `json:"assignees_string,omitempty"`

	ClosedByPullRequestID string `json:"closed_by_pull_request_id,omitempty"`

	// OutPut-Serial
	Labels    *[]github.Label `json:"labels,omitempty" gorm:"-"`
	Assignees *[]github.User  `json:"assignees,omitempty" gorm:"-"`
}

// List Option
type IssueOption struct {
	ID      int64  `json:"id"`
	IssueID string `json:"issue_id,omitempty"`
	Number  int    `json:"number,omitempty"`
	State   string `json:"state,omitempty"`
	Owner   string `json:"owner,omitempty"`
	Repo    string `json:"repo,omitempty"`
}

// DB-Table
func (Issue) TableName() string {
	return "issue"
}

// ComposeIssueFromV3
func ComposeIssueFromV3(issue *github.Issue) *Issue {
	labels := &[]github.Label{}
	for _, node := range issue.Labels {
		label := &github.Label{
			ID:   node.ID,
			Name: node.Name,
		}
		*labels = append(*labels, *label)
	}
	assignees := &[]github.User{}
	for _, node := range issue.Assignees {
		user := &github.User{
			Login: node.Login,
		}
		*assignees = append(*assignees, *user)
	}
	url := strings.Split(*issue.RepositoryURL, "/")
	owner := url[len(url)-2]
	repo := url[len(url)-1]

	return &Issue{
		IssueID: *issue.NodeID,
		Number:  *issue.Number,
		State:   *issue.State,
		Title:   *issue.Title,
		Owner:   owner,
		Repo:    repo,
		HTMLURL: *issue.HTMLURL,

		CreatedAt: *issue.CreatedAt,
		UpdatedAt: *issue.UpdatedAt,
		ClosedAt:  issue.ClosedAt,

		Labels:    labels,
		Assignees: assignees,
	}
}

// ComposeIssueFromV4
// TODO: v4 implement by tony at 2022/02/14
func ComposeIssueFromV4(issueFiled *git.IssueField) *Issue {
	labels := &[]github.Label{}
	for _, labelNode := range issueFiled.Labels.Nodes {
		label := github.Label{
			Name: github.String(string(labelNode.Name)),
		}
		*labels = append(*labels, label)
	}
	assignees := &[]github.User{}
	for _, userNode := range issueFiled.Assignees.Nodes {
		user := github.User{
			Login: (*string)(&userNode.Login),
		}
		*assignees = append(*assignees, user)
	}
	closedByPrID := ""
	if issueFiled.State == githubv4.IssueStateClosed {
		for _, edge := range issueFiled.TimelineItems.Edges {
			closer := edge.Node.ClosedEvent.Closer.PullRequest
			if closer.Number != 0 {
				closedByPrID = closer.ID.(string)
			}
		}
	}

	return &Issue{
		IssueID: issueFiled.ID.(string),
		Number:  int(issueFiled.Number),
		State:   string(issueFiled.State),
		Title:   string(issueFiled.Title),
		Owner:   string(issueFiled.Repository.Owner.Login),
		Repo:    string(issueFiled.Repository.Name),
		HTMLURL: string(issueFiled.Url),

		CreatedAt: issueFiled.CreatedAt.Time,
		UpdatedAt: issueFiled.UpdatedAt.Time,
		ClosedAt:  &issueFiled.ClosedAt.Time,

		Labels:    labels,
		Assignees: assignees,

		ClosedByPullRequestID: closedByPrID,
	}
}

/**

CREATE TABLE IF NOT EXISTS issue (
	id INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
	issue_id VARCHAR(255) NOT NULL COMMENT 'Issue全局ID',
	number INT(11) NOT NULL COMMENT '当前库ID',
	state VARCHAR(32) NOT NULL COMMENT '状态',
	title VARCHAR(1024) COMMENT '标题',
	owner VARCHAR(255) COMMENT '仓库所有者',
	repo VARCHAR(255) COMMENT '仓库名称',
	html_url VARCHAR(1024) COMMENT '链接',

	closed_at TIMESTAMP COMMENT '关闭时间',
	created_at TIMESTAMP COMMENT '创建时间',
	updated_at TIMESTAMP COMMENT '更新时间',

	labels_string TEXT COMMENT '标签',
	assignees_string TEXT COMMENT '处理人列表',

	closed_by_pull_request_id VARCHAR(255) COMMENT '处理的PR',

	PRIMARY KEY (id),
	UNIQUE KEY uk_issueid (issue_id),
	INDEX idx_state (state),
	INDEX idx_owner_repo (owner, repo),
	INDEX idx_createdat (created_at)
)
ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT 'issue信息表';

**/
