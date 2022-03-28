package repository

import (
	"fmt"

	"tirelease/commons/database"
	"tirelease/internal/dto"
	"tirelease/internal/entity"

	"github.com/pkg/errors"
)

func SelectIssueRelationInfoByJoin(option *dto.IssueRelationInfoQuery) (*[]dto.IssueRelationInfoByJoin, error) {
	issueAffect := &entity.IssueAffectOption{
		AffectVersion: option.AffectVersion,
		AffectResult:  option.AffectResult,
	}
	// 查询
	var issueRelationInfoJoin []dto.IssueRelationInfoByJoin
	sql := "select issue.issue_id, group_concat(issue_affect.id) as issue_affect_ids from "
	sql += " ( "
	sql += "select * from issue where 1=1 " + IssueWhere(&option.IssueOption)
	sql += " ) as issue "
	sql += "left join "
	sql += " ( "
	sql += "select * from issue_affect where 1=1 " + IssueAffectWhere(issueAffect)
	sql += " ) as issue_affect "
	sql += "on issue.issue_id = issue_affect.issue_id "
	sql += "where issue_affect.issue_id is not null "
	sql += "group by issue.issue_id "
	sql += IssueOrderBy(&option.IssueOption) + IssueLimit(&option.IssueOption)

	if err := database.DBConn.RawWrapper(sql, option).Find(&issueRelationInfoJoin).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("select issue_relation by raw by join failed, option: %+v", option))
	}

	return &issueRelationInfoJoin, nil
}
