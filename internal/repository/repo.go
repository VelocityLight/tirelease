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

func SelectRepo(option *entity.RepoOption) (*[]entity.Repo, error) {
	// 查询
	var repos []entity.Repo
	if err := database.DBConn.DB.Where(option).Find(&repos).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("find repo: %+v failed", option))
	}
	return &repos, nil
}
