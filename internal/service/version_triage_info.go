package service

import (
	"fmt"

	// "tirelease/commons/git"
	"tirelease/commons/git"
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
	if !isFrozen && isAccept {
		err := AddLabelByIssueID(versionTriage.IssueID, git.CherryPickLabel)
		if err != nil {
			return nil, err
		}
	}

	// Return
	return &dto.VersionTriageInfo{
		ReleaseVersion: releaseVersion,

		VersionTriage: versionTriage,
		IsFrozen:      isFrozen,
		IsAccept:      isAccept,
	}, nil
}

func SelectVersionTriageInfo(query *dto.VersionTriageInfoQuery) (*dto.VersionTriageInfoWrap, error) {
	// Option
	versionTriageOption := &entity.VersionTriageOption{
		ID:           query.ID,
		IssueID:      query.IssueID,
		VersionName:  query.VersionName,
		TriageResult: query.TriageResult,
	}

	// Select
	versionTriages, err := repository.SelectVersionTriage(versionTriageOption)
	if err != nil {
		return nil, err
	}
	releaseVersion, err := repository.SelectReleaseVersionUnique(&entity.ReleaseVersionOption{
		Name: query.VersionName,
	})
	if err != nil {
		return nil, err
	}

	// Compose
	versionTriageInfos := make([]dto.VersionTriageInfo, 0)
	for _, versionTriage := range *versionTriages {
		issueRelationInfos, err := SelectIssueRelationInfo(&dto.IssueRelationInfoQuery{
			IssueID:    versionTriage.IssueID,
			BaseBranch: releaseVersion.ReleaseBranch,
		})
		if err != nil {
			return nil, err
		}

		versionTriageInfo := dto.VersionTriageInfo{}
		versionTriageInfo.VersionTriage = &versionTriage
		if len(*issueRelationInfos) == 1 {
			versionTriageInfo.IssueRelationInfo = &(*issueRelationInfos)[0]
		}
		versionTriageInfo.ReleaseVersion = releaseVersion
		versionTriageInfos = append(versionTriageInfos, versionTriageInfo)
	}

	return &dto.VersionTriageInfoWrap{
		ReleaseVersion:     releaseVersion,
		VersionTriageInfos: &versionTriageInfos,
	}, nil
}

func CheckReleaseVersion(option *entity.ReleaseVersionOption) (*entity.ReleaseVersion, error) {
	releaseVersion, err := repository.SelectReleaseVersionUnique(option)
	if err != nil {
		return nil, err
	}
	if releaseVersion.Status == entity.ReleaseVersionStatusReleased || releaseVersion.Status == entity.ReleaseVersionStatusClosed {
		return nil, errors.Wrap(err, fmt.Sprintf("find release version is already closed or released: %+v failed", releaseVersion))
	}
	return releaseVersion, nil
}
