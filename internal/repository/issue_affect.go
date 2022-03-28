package repository

import (
	"fmt"
	"time"

	"tirelease/commons/database"
	"tirelease/internal/entity"

	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
)

// func CreateIssueAffect(issueAffect *entity.IssueAffect) error {
// 	issueAffect.CreateTime = time.Now()
// 	issueAffect.UpdateTime = time.Now()
// 	// 存储
// 	if err := database.DBConn.DB.Clauses(
// 		clause.OnConflict{DoNothing: true}).Create(&issueAffect).Error; err != nil {
// 		return errors.Wrap(err, fmt.Sprintf("create issue affect: %+v failed", issueAffect))
// 	}
// 	return nil
// }

func SelectIssueAffect(option *entity.IssueAffectOption) (*[]entity.IssueAffect, error) {
	// search
	var issueAffects []entity.IssueAffect
	if err := database.DBConn.DB.Where(option).Find(&issueAffects).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("find issue affect: %+v failed", option))
	}
	return &issueAffects, nil
}

func CreateOrUpdateIssueAffect(issueAffect *entity.IssueAffect) error {
	// ignore useless create or update
	// if issueAffect.AffectResult == "" || issueAffect.AffectResult == entity.AffectResultResultUnKnown {
	// 	return nil
	// }

	// update
	if issueAffect.CreateTime.IsZero() {
		issueAffect.CreateTime = time.Now()
	}
	if issueAffect.UpdateTime.IsZero() {
		issueAffect.UpdateTime = time.Now()
	}
	if err := database.DBConn.DB.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"update_time", "affect_result"}),
	}).Create(&issueAffect).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("create or update issue affect: %+v failed", issueAffect))
	}
	return nil
}

func IssueAffectWhere(option *entity.IssueAffectOption) string {
	sql := ""

	if option.ID != 0 {
		sql += " and issue_affect.id = @ID"
	}
	if option.IssueID != "" {
		sql += " and issue_affect.issue_id = @IssueID"
	}
	if option.AffectVersion != "" {
		sql += " and issue_affect.affect_version = @AffectVersion"
	}

	return sql
}
