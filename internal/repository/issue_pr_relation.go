package repository

import (
	"fmt"
	"time"

	"tirelease/commons/database"
	"tirelease/internal/entity"

	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
)

func CreateIssuePrRelation(issuePrRelation *entity.IssuePrRelation) error {
	if issuePrRelation.CreateTime.IsZero() {
		issuePrRelation.CreateTime = time.Now()
	}
	if issuePrRelation.UpdateTime.IsZero() {
		issuePrRelation.UpdateTime = time.Now()
	}
	// 存储
	if err := database.DBConn.DB.Clauses(
		clause.OnConflict{DoNothing: true}).Create(&issuePrRelation).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("create issue_pr_relation: %+v failed", issuePrRelation))
	}
	return nil
}

func SelectIssuePrRelation(option *entity.IssuePrRelationOption) (*[]entity.IssuePrRelation, error) {
	// 查询
	var issuePrRelations []entity.IssuePrRelation
	if err := database.DBConn.DB.Where(option).Find(&issuePrRelations).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("find issue_pr_relation: %+v failed", option))
	}
	return &issuePrRelations, nil
}

// func DeleteIssuePrRelation(issuePrRelation *entity.IssuePrRelation) error {
// 	if err := database.DBConn.DB.Delete(issuePrRelation).Error; err != nil {
// 		return errors.Wrap(err, fmt.Sprintf("delete issue_pr_relation: %+v failed", issuePrRelation))
// 	}
// 	return nil
// }
