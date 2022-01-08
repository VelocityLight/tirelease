package repository

import (
	"fmt"

	"tirelease/commons/database"
	"tirelease/internal/entity"

	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
)

func CreateIssueAffect(issueAffect *entity.IssueAffect) error {
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

func UpdateIssueAffect(issueAffect *entity.IssueAffect) error {
	// 更新
	if err := database.DBConn.DB.Save(&issueAffect).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("update issue affect: %+v failed", issueAffect))
	}
	return nil
}
