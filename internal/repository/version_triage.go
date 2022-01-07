package repository

import (
	"fmt"

	"tirelease/commons/database"
	"tirelease/internal/entity"

	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
)

func CreateVersionTriage(versionTriage *entity.VersionTriage) error {
	// 存储
	if err := database.DBConn.DB.Clauses(
		clause.OnConflict{DoNothing: true}).Create(&versionTriage).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("create version triage: %+v failed", versionTriage))
	}
	return nil
}

func SelectVersionTriage(option *entity.VersionTriageOption) (*[]entity.VersionTriage, error) {
	// 查询
	var versionTriages []entity.VersionTriage
	if err := database.DBConn.DB.Find(&versionTriages).Where(option).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("find version triage: %+v failed", option))
	}
	return &versionTriages, nil
}

func UpdateVersionTriage(versionTriage *entity.VersionTriage) error {
	// 更新
	if err := database.DBConn.DB.Save(&versionTriage).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("update version triage: %+v failed", versionTriage))
	}
	return nil
}
