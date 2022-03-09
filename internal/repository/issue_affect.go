package repository

import (
	"fmt"
	"time"

	"tirelease/commons/database"
	"tirelease/internal/entity"

	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
)

func CreateIssueAffect(issueAffect *entity.IssueAffect) error {
	issueAffect.CreateTime = time.Now()
	issueAffect.UpdateTime = time.Now()
	// 存储
	if err := database.DBConn.DB.Clauses(
		clause.OnConflict{DoNothing: true}).Create(&issueAffect).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("create issue affect: %+v failed", issueAffect))
	}
	return nil
}

func SelectIssueAffect(option *entity.IssueAffectOption) (*[]entity.IssueAffect, error) {
	// 查询
	var issueAffects []entity.IssueAffect
	if err := database.DBConn.DB.Where(option).Find(&issueAffects).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("find issue affect: %+v failed", option))
	}
	return &issueAffects, nil
}

func CreateOrUpdateIssueAffect(issueAffect *entity.IssueAffect) error {
	// 更新
	issueAffect.CreateTime = time.Now()
	issueAffect.UpdateTime = time.Now()
	if err := database.DBConn.DB.Clauses(clause.OnConflict{
		// DoUpdates: clause.AssignmentColumns([]string{"update_time", "affect_result"}),
		UpdateAll: true,
	}).Create(&issueAffect).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("create or update issue affect: %+v failed", issueAffect))
	}
	return nil
}
