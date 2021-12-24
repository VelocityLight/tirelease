package entity

import (
	"time"
)

/**
mysql --host 172.16.4.36 --port 3306 -u cicd_online -pwGEXq8a4MeCw6G

CREATE TABLE IF NOT EXISTS test_cases (
	id INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
	create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	update_time DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	job_name VARCHAR(255) COMMENT '任务名称',
	job_url VARCHAR(1024) COMMENT '任务链接',
	repo VARCHAR(64) NOT NULL COMMENT '代码仓库',
	branch VARCHAR(64) COMMENT '代码分支',
	pull_request_id INT(11) COMMENT '合入请求ID',
	commit_id VARCHAR(255) COMMENT '代码提交ID',
	suite_name VARCHAR(255) COMMENT '组件名称',
	case_name VARCHAR(64) NOT NULL COMMENT '用例名称',
	case_class VARCHAR(1024) COMMENT '用例类名',
	execution_time VARCHAR(255) COMMENT '执行时长',
	status VARCHAR(32) NOT NULL COMMENT '用例结果',
	error_detail TEXT COMMENT '错误信息',
	stack_trace TEXT COMMENT '错误堆栈',

	PRIMARY KEY (id),
	// UNIQUE KEY uk_xxx (xxx)
	INDEX idx_createtime (create_time),
	INDEX idx_jobname_repo_branch (job_name, repo, branch)
)
ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT 'ci_detail任务详情表';
**/

// Struct of ci_detail
type TestEntity struct {
	ID            int64       `json:"id"`
	CreateTime    time.Time   `json:"create_time"`
	UpdateTime    time.Time   `json:"update_time"`
	JobName       string      `json:"job_name"`
	JobURL        string      `json:"job_url"`
	Repo          string      `json:"repo"`
	Branch        string      `json:"branch"`
	PullRequestID string      `json:"pull_request_id"`
	CommitID      string      `json:"commit_id"`
	SuiteName     string      `json:"suite_name"`
	CaseName      string      `json:"case_name"`
	CaseClass     string      `json:"case_class"`
	ExecutionTime string      `json:"execution_time"`
	Status        *CaseStatus `json:"status"`
	ErrorDetail   string      `json:"error_detail"`
	StackTrace    string      `json:"stack_trace"`
}

// Enum type
type CaseStatus string

// Enum list...
const (
	CaseStatusPassed = CaseStatus("passed")
	CaseStatusFailed = CaseStatus("failed")
	CaseStatusSkiped = CaseStatus("skipped")
)

// List Option
type ListOption struct {
}
