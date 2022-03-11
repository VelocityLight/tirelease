package service

import (
	"fmt"

	"tirelease/commons/git"
	"tirelease/internal/entity"
	"tirelease/internal/repository"
)

func CreateOrUpdateIssueAffect(issueAffect *entity.IssueAffect) error {
	// create or update
	err := repository.CreateOrUpdateIssueAffect(issueAffect)
	if err != nil {
		return err
	}

	// result operation
	err = IssueAffectResultOperation(issueAffect)
	if err != nil {
		return err
	}

	return nil
}

func IssueAffectResultOperation(issueAffect *entity.IssueAffect) error {
	// param protection
	if issueAffect.AffectResult != entity.AffectResultResultYes {
		return nil
	}

	// select latest version & insert cherry-pick
	releaseVersionOption := &entity.ReleaseVersionOption{
		FatherReleaseVersionName: issueAffect.AffectVersion,
		Status:                   entity.ReleaseVersionStatusOpen,
	}
	releaseVersion, err := repository.SelectReleaseVersionUnique(releaseVersionOption)
	if err != nil {
		return err
	}
	versionTriage := &entity.VersionTriage{
		IssueID:      issueAffect.IssueID,
		VersionName:  releaseVersion.Name,
		TriageResult: entity.VersionTriageResultUnKnown,
	}
	_, err = CreateOrUpdateVersionTriageInfo(versionTriage)
	if err != nil {
		return err
	}

	// git add label
	label := fmt.Sprintf(git.AffectsLabel, releaseVersion.Name)
	err = AddLabelByIssueID(issueAffect.IssueID, label)
	if err != nil {
		return err
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
