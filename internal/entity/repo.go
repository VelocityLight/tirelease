package entity

import (
	"time"
)

// Struct of Issue
type Repo struct {
	// DataBase Column
	ID      int64  `json:"id,omitempty"`
	CreatedAt           time.Time      `json:"created_at,omitempty"`
	UpdatedAt           time.Time      `json:"updated_at,omitempty"`

	Owner   string `json:"owner,omitempty"`
	Repo    string `json:"repo,omitempty"`
	FullName string `json:"full_name,omitempty"`
	HTMLURL *string `json:"html_url,omitempty"`
	Description         *string         `json:"description,omitempty"`
}

// List Option
type RepoOption struct {
	ID      int64  `json:"id"`
	Owner   string `json:"owner,omitempty"`
	Repo    string `json:"repo,omitempty"`
	FullName string `json:"full_name,omitempty"`
}

// DB-Table
func (Repo) TableName() string {
	return "repo"
}

/**

CREATE TABLE IF NOT EXISTS repo (
	id INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
	created_at TIMESTAMP COMMENT '创建时间',
	updated_at TIMESTAMP COMMENT '更新时间',
	
	owner VARCHAR(255) COMMENT '仓库所有者',
	repo VARCHAR(255) COMMENT '仓库名称',
	full_name VARCHAR(255) COMMENT '仓库全称',
	html_url VARCHAR(1024) COMMENT '链接',
	description VARCHAR(1024) COMMENT '描述',

	PRIMARY KEY (id),
	UNIQUE KEY uk_owner_repo (owner, repo),
	INDEX idx_fullname (full_name)
)
ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT 'repo信息表';

**/
