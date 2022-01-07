package repository

import (
	"encoding/json"
	"fmt"

	"tirelease/commons/database"
	"tirelease/internal/entity"

	"github.com/pkg/errors"
)

func CreateReleaseVersion(version *entity.ReleaseVersion) error {
	// 加工
	serializeReleaseVersion(version)

	// 存储
	if err := database.DBConn.DB.Omit("Repos", "Labels").Create(&version).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("create release version: %+v failed", version))
	}
	return nil
}

func UpdateReleaseVersion(version *entity.ReleaseVersion) error {
	// 加工
	serializeReleaseVersion(version)

	// 更新
	if err := database.DBConn.DB.Omit("Repos", "Labels").Save(&version).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("update release version: %+v failed", version))
	}
	return nil
}

func SelectReleaseVersion(option *entity.ReleaseVersionOption) (*[]entity.ReleaseVersion, error) {
	// 查询
	var releaseVersions []entity.ReleaseVersion
	if err := database.DBConn.DB.Find(&releaseVersions).Where(option).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("find release version: %+v failed", option))
	}

	// 加工
	for i := 0; i < len(releaseVersions); i++ {
		unSerializeReleaseVersion(&releaseVersions[i])
	}
	return &releaseVersions, nil
}

// 序列化和反序列化
func serializeReleaseVersion(version *entity.ReleaseVersion) {
	if nil != version.Repos {
		reposString, _ := json.Marshal(version.Repos)
		version.ReposString = string(reposString)
	}
	if nil != version.Labels {
		labelsString, _ := json.Marshal(version.Labels)
		version.LabelsString = string(labelsString)
	}
}

func unSerializeReleaseVersion(version *entity.ReleaseVersion) {
	if version.ReposString != "" {
		var repos []string
		json.Unmarshal([]byte(version.ReposString), &repos)
		version.Repos = &repos
	}
	if version.LabelsString != "" {
		var labels []string
		json.Unmarshal([]byte(version.LabelsString), &labels)
		version.Labels = &labels
	}
}
