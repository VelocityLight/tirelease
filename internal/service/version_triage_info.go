package service

import (
	"fmt"

	// "tirelease/commons/git"
	"tirelease/internal/dto"
	"tirelease/internal/entity"
	"tirelease/internal/repository"

	"github.com/pkg/errors"
)

func CreateOrUpdateVersionTriageInfo(versionTriage *entity.VersionTriage) (*dto.VersionTriageInfo, error) {
	// Check
	releaseVersionOption := &entity.ReleaseVersionOption{
		Name: versionTriage.VersionName,
	}
	releaseVersion, err := CheckReleaseVersion(releaseVersionOption)
	if err != nil {
		return nil, err
	}

	// Create Or Update
	var isFrozen bool = releaseVersion.Status == entity.ReleaseVersionStatusFrozen
	var isAccept bool = versionTriage.TriageResult == entity.VersionTriageResultAccept
	if isFrozen && isAccept {
		versionTriage.TriageResult = entity.VersionTriageResultAcceptFrozen
	}
	if err := repository.CreateOrUpdateVersionTriage(versionTriage); err != nil {
		return nil, err
	}

	// Operate Git
	// if !isFrozen && isAccept {
	// 	// TODO by tony: git label
	// }

	// Return
	return &dto.VersionTriageInfo{
		ReleaseVersion: releaseVersion,

		VersionTriage: versionTriage,
		IsFrozen:      isFrozen,
		IsAccept:      isAccept,
	}, nil
}

func SelectVersionTriageInfo() {

}

func CheckReleaseVersion(option *entity.ReleaseVersionOption) (*entity.ReleaseVersion, error) {
	releaseVersion, err := repository.SelectReleaseVersionUnique(option)
	if err != nil {
		return nil, err
	}
	if releaseVersion == nil {
		return nil, errors.Wrap(err, fmt.Sprintf("find release version unique is nil: %+v failed", option))
	}
	if releaseVersion.Status == entity.ReleaseVersionStatusReleased || releaseVersion.Status == entity.ReleaseVersionStatusClosed {
		return nil, errors.Wrap(err, fmt.Sprintf("find release version is already closed or released: %+v failed", releaseVersion))
	}
	return releaseVersion, nil
}
