package service

import (
	"fmt"
	"strings"

	// "tirelease/commons/git"
	"tirelease/commons/git"
	"tirelease/internal/dto"
	"tirelease/internal/entity"
	"tirelease/internal/repository"

	"github.com/pkg/errors"
)

func CreateOrUpdateVersionTriageInfo(versionTriage *entity.VersionTriage) (*dto.VersionTriageInfo, error) {
	// check
	shortType := ComposeVersionShortType(versionTriage.VersionName)
	major, minor, patch, _ := ComposeVersionAtom(versionTriage.VersionName)
	releaseVersionOption := &entity.ReleaseVersionOption{}
	if shortType == entity.ReleaseVersionShortTypeMinor {
		releaseVersionOption.Major = major
		releaseVersionOption.Minor = minor
		releaseVersionOption.Status = entity.ReleaseVersionStatusUpcoming
	} else if shortType == entity.ReleaseVersionShortTypePatch || shortType == entity.ReleaseVersionShortTypeHotfix {
		releaseVersionOption.Major = major
		releaseVersionOption.Minor = minor
		releaseVersionOption.Patch = patch
	} else {
		return nil, errors.New(fmt.Sprintf("CreateOrUpdateVersionTriageInfo params invalid: %+v failed", versionTriage))
	}
	releaseVersion, err := CheckReleaseVersion(releaseVersionOption)
	if err != nil {
		return nil, err
	}
	releaseBranch := releaseVersion.ReleaseBranch
	versionTriage.VersionName = releaseVersion.Name

	// basic info
	issueRelationInfo, err := SelectIssueRelationInfoUnique(&dto.IssueRelationInfoQuery{
		IssueOption: entity.IssueOption{
			IssueID: versionTriage.IssueID,
		},
		BaseBranch: releaseBranch,
	})
	if err != nil {
		return nil, err
	}

	// create or update
	var isFrozen bool = releaseVersion.Status == entity.ReleaseVersionStatusFrozen
	var isAccept bool = versionTriage.TriageResult == entity.VersionTriageResultAccept
	if isFrozen && isAccept {
		versionTriage.TriageResult = entity.VersionTriageResultAcceptFrozen
	}
	if issueRelationInfo.Issue.SeverityLabel == git.SeverityCriticalLabel {
		versionTriage.BlockVersionRelease = entity.BlockVersionReleaseResultBlock
	}
	if err := repository.CreateOrUpdateVersionTriage(versionTriage); err != nil {
		return nil, err
	}

	// remote operation
	if len(*issueRelationInfo.PullRequests) > 0 {
		for i := range *issueRelationInfo.PullRequests {
			pr := (*issueRelationInfo.PullRequests)[i]
			if !isFrozen && isAccept {
				err := RemoveLabelByPullRequestID(pr.PullRequestID, git.NotCheryyPickLabel)
				if err != nil {
					return nil, err
				}

				err = AddLabelByPullRequestID(pr.PullRequestID, git.CherryPickLabel)
				if err != nil {
					return nil, err
				}
			}
			var isRelease bool = releaseVersion.Status == entity.ReleaseVersionStatusReleased
			if !isAccept && !isRelease {
				err := RemoveLabelByPullRequestID(pr.PullRequestID, git.CherryPickLabel)
				if err != nil {
					return nil, err
				}

				err = AddLabelByPullRequestID(pr.PullRequestID, git.NotCheryyPickLabel)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	// status
	var status entity.VersionTriageMergeStatus
	if len(*issueRelationInfo.PullRequests) == 0 {
		status = entity.VersionTriageMergeStatusPr
	} else {
		allMerge := true
		for _, pr := range *issueRelationInfo.PullRequests {
			if !pr.CherryPickApproved {
				status = entity.VersionTriageMergeStatusApprove
				break
			} else if !pr.AlreadyReviewed {
				status = entity.VersionTriageMergeStatusReview
				break
			} else if !pr.Merged {
				allMerge = false
				break
			}
		}
		if allMerge {
			status = entity.VersionTriageMergeStatusMerged
		} else {
			status = entity.VersionTriageMergeStatusCITesting
		}
	}

	// return
	return &dto.VersionTriageInfo{
		ReleaseVersion: releaseVersion,
		IsFrozen:       isFrozen,
		IsAccept:       isAccept,

		VersionTriage:            versionTriage,
		VersionTriageMergeStatus: status,

		IssueRelationInfo: issueRelationInfo,
	}, nil
}

func SelectVersionTriageInfo(query *dto.VersionTriageInfoQuery) (*dto.VersionTriageInfoWrap, error) {
	// Option
	versionTriageOption := &entity.VersionTriageOption{
		ID:           query.ID,
		IssueID:      query.IssueID,
		VersionName:  query.Version,
		TriageResult: query.TriageResult,
	}

	// Select
	versionTriages, err := repository.SelectVersionTriage(versionTriageOption)
	if err != nil {
		return nil, err
	}

	releaseVersion, err := repository.SelectReleaseVersionLatest(&entity.ReleaseVersionOption{
		Name: query.Version,
	})
	if err != nil {
		return nil, err
	}

	// Compose
	versionTriageInfos := make([]dto.VersionTriageInfo, 0)
	for i := range *versionTriages {
		versionTriage := (*versionTriages)[i]
		issueRelationInfo, err := SelectIssueRelationInfoUnique(&dto.IssueRelationInfoQuery{
			IssueOption: entity.IssueOption{
				IssueID: versionTriage.IssueID,
			},
			BaseBranch: releaseVersion.ReleaseBranch,
		})
		if err != nil {
			return nil, err
		}

		versionTriageInfo := dto.VersionTriageInfo{}
		versionTriageInfo.VersionTriage = &versionTriage
		versionTriageInfo.IssueRelationInfo = issueRelationInfo
		versionTriageInfo.ReleaseVersion = releaseVersion
		versionTriageInfos = append(versionTriageInfos, versionTriageInfo)
	}

	return &dto.VersionTriageInfoWrap{
		ReleaseVersion:     releaseVersion,
		VersionTriageInfos: &versionTriageInfos,
	}, nil
}

func InheritVersionTriage(fromVersion string, toVersion string) error {
	// Select
	versionTriageOption := &entity.VersionTriageOption{
		VersionName: fromVersion,
	}
	versionTriages, err := repository.SelectVersionTriage(versionTriageOption)
	if err != nil {
		return err
	}
	if len(*versionTriages) == 0 {
		return nil
	}

	// Migrate
	for i := range *versionTriages {
		versionTriage := (*versionTriages)[i]
		switch versionTriage.TriageResult {
		case entity.VersionTriageResultAccept:
			versionTriage.TriageResult = entity.VersionTriageResultReleased
		case entity.VersionTriageResultUnKnown:
			versionTriage.VersionName = toVersion
		case entity.VersionTriageResultLater, entity.VersionTriageResultAcceptFrozen:
			versionTriage.VersionName = toVersion
			versionTriage.TriageResult = entity.VersionTriageResultAccept
		}
		if err := repository.CreateOrUpdateVersionTriage(&versionTriage); err != nil {
			return err
		}
	}

	return nil
}

func CheckReleaseVersion(option *entity.ReleaseVersionOption) (*entity.ReleaseVersion, error) {
	if option == nil || option.Name == "" {
		return nil, errors.New(fmt.Sprintf("CheckReleaseVersion params invalid: %+v failed", option))
	}
	releaseVersion, err := repository.SelectReleaseVersionLatest(option)
	if err != nil {
		return nil, err
	}
	if releaseVersion.Status == entity.ReleaseVersionStatusReleased || releaseVersion.Status == entity.ReleaseVersionStatusCancelled {
		return nil, errors.Wrap(err, fmt.Sprintf("find release version is already released or cancelled: %+v failed", releaseVersion))
	}
	return releaseVersion, nil
}

// Export history data (Only database operation, no remote operation)
func ExportHistoryVersionTriageInfo(info *dto.IssueRelationInfo, releaseVersions *[]entity.ReleaseVersion) error {
	// param check
	if info == nil || releaseVersions == nil {
		return errors.New("ExportHistoryVersionTriageInfo params invalid")
	}
	if info.PullRequests == nil || len(*info.PullRequests) == 0 {
		return nil
	}

	// insert version triage
	for i := range *info.PullRequests {
		pr := (*info.PullRequests)[i]
		if !pr.Merged || !pr.CherryPickApproved ||
			!strings.HasPrefix(pr.BaseBranch, git.ReleaseBranchPrefix) {
			continue
		}
		releaseBranch := string(pr.BaseBranch)
		branchVersion := strings.Replace(pr.BaseBranch, git.ReleaseBranchPrefix, "", -1)
		major, minor, _, _ := ComposeVersionAtom(branchVersion)

		// search version in time section
		// release version is already sorted desc
		for i := len(*releaseVersions) - 1; i >= 0; i-- {
			releaseVersion := (*releaseVersions)[i]
			if releaseVersion.Status != entity.ReleaseVersionStatusReleased {
				continue
			}
			if releaseVersion.Major != major || releaseVersion.Minor != minor || releaseVersion.ReleaseBranch != releaseBranch {
				continue
			}
			if releaseVersion.ActualReleaseTime.After(*(pr.MergedAt)) {
				versionTriage := &entity.VersionTriage{
					IssueID:      info.Issue.IssueID,
					VersionName:  releaseVersion.Name,
					TriageResult: entity.VersionTriageResultReleased,
					CreateTime:   *(pr.MergedAt),
					UpdateTime:   *(pr.MergedAt),
				}
				if err := repository.CreateOrUpdateVersionTriage(versionTriage); err != nil {
					return err
				}

				issueAffect := &entity.IssueAffect{
					IssueID:       info.Issue.IssueID,
					AffectVersion: branchVersion,
					AffectResult:  entity.AffectResultResultYes,
					CreateTime:    *(pr.MergedAt),
					UpdateTime:    *(pr.MergedAt),
				}
				if err := repository.CreateOrUpdateIssueAffect(issueAffect); err != nil {
					return err
				}

				break
			}
		}
	}

	return nil
}
