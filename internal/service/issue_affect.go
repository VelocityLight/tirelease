package service

import (
	"fmt"

	"github.com/pkg/errors"
	"tirelease/internal/entity"
	"tirelease/internal/repository"
)

func UpdateIssueAffect(issueAffect *entity.IssueAffect) error {
	err := repository.CreateOrUpdateIssueAffect(issueAffect)
	if err != nil {
		return err
	}

	// get lastest release_version
	releaseVersionOption := &entity.ReleaseVersionOption{
		FatherReleaseVersionName: issueAffect.AffectVersion,
		Status:                   entity.ReleaseVersionStatusOpen,
	}
	releaseVersions, err := repository.SelectReleaseVersion(releaseVersionOption)
	if err != nil {
		return err
	}
	if len(*releaseVersions) == 0 {
		return errors.Wrap(err, fmt.Sprintf("no active release version: %+v failed", releaseVersionOption))
	}

	// accept operate
	if issueAffect.AffectResult == entity.AffectResultResultYes {
		// todo update git label

		// insert cherry-pick
		versionTriage := &entity.VersionTriage{
			IssueID:      issueAffect.IssueID,
			VersionName:  (*releaseVersions)[0].Name,
			TriageResult: entity.VersionTriageResultUnKnown,
		}
		_, err := CreateOrUpdateVersionTriageInfo(versionTriage)
		if err != nil {
			return err
		}
	}
	return nil
}

func ComposeIssueAffectWithIssueID(issueID string) (*[]entity.IssueAffect, error) {
	// Select Exist Issue Affect
	issueAffectOption := &entity.IssueAffectOption{
		IssueID: issueID,
	}
	issueAffects, err := repository.SelectIssueAffect(issueAffectOption)
	if err != nil {
		return nil, err
	}

	// Implement New Issue Affect
	releaseVersionOption := &entity.ReleaseVersionOption{
		Type:   entity.ReleaseVersionTypeMinor,
		Status: entity.ReleaseVersionStatusOpen,
	}
	releaseVersions, err := repository.SelectReleaseVersion(releaseVersionOption)
	if nil != err {
		return nil, err
	}
	for _, releaseVersion := range *releaseVersions {
		var isExist bool = false
		for _, issueAffect := range *issueAffects {
			if issueAffect.AffectVersion == releaseVersion.Name {
				isExist = true
				break
			}
		}

		if !isExist {
			newAffect := &entity.IssueAffect{
				IssueID:       issueID,
				AffectVersion: releaseVersion.Name,
				AffectResult:  entity.AffectResultResultUnKnown,
			}
			(*issueAffects) = append((*issueAffects), *newAffect)
		}
	}
	return issueAffects, nil
}
