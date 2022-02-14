package repository

import (
	"fmt"

	"tirelease/commons/database"
	"tirelease/internal/entity"

	"github.com/pkg/errors"
)

func CreateRepo(repo *entity.Repo) error {
	// 存储
	if err := database.DBConn.DB.Create(&repo).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("create repo: %+v failed", repo))
	}
	return nil
}

func UpdateRepo(repo *entity.Repo) error {
	// 更新
	if err := database.DBConn.DB.Save(&repo).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("update repo: %+v failed", repo))
	}
	return nil
}
