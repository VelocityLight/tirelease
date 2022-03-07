package service

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"tirelease/internal/entity"
	"tirelease/internal/repository"
)

func UpdateIssueAffect(issueAffect *entity.IssueAffect) error {
	err := repository.UpdateIssueAffect(issueAffect)
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
			CreateTime: time.Now(),
			UpdateTime: time.Now(),

			IssueID:      issueAffect.IssueID,
			VersionName:  (*releaseVersions)[0].Name,
			TriageResult: entity.VersionTriageResultAccept,
		}
		_, err := CreateOrUpdateVersionTriageInfo(versionTriage)
		if err != nil {
			return err
		}
	}
	return nil
}
