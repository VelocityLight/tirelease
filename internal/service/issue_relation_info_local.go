package service

import (
	"fmt"

	"tirelease/internal/dto"
	"tirelease/internal/entity"
	"tirelease/internal/repository"

	"github.com/pkg/errors"
)

// ============================================================================
// ============================================================================ CURD Of IssueRelationInfo

func SelectIssueRelationInfo(option *dto.IssueRelationInfoQuery) (*[]dto.IssueRelationInfo, error) {
	// Select Issues
	issueOption := &entity.IssueOption{
		ID:      option.ID,
		IssueID: option.IssueID,
		Number:  option.Number,
		State:   option.State,
		Owner:   option.Owner,
		Repo:    option.Repo,

		SeverityLabel: option.SeverityLabel,
		TypeLabel:     option.TypeLabel,
	}
	issues, err := repository.SelectIssue(issueOption)
	if nil != err {
		return nil, err
	}

	// From Issue to select IssueAffects & IssuePrRelations & PullRequests
	alls := make([]dto.IssueRelationInfo, 0)
	for _, issue := range *issues {
		issueRelationInfo, err := ComposeRelationInfoByIssue(&issue)
		if nil != err {
			return nil, err
		}
		alls = append(alls, *issueRelationInfo)
	}

	// Filter & Result
	issueRelationInfos := make([]dto.IssueRelationInfo, 0)
	for _, issueRelationInfo := range alls {
		var filter bool = false
		for _, issueAffect := range *issueRelationInfo.IssueAffects {
			if option.AffectVersion == "" || issueAffect.AffectVersion == option.AffectVersion {
				filter = true
				break
			}
		}
		if filter {
			issueRelationInfos = append(issueRelationInfos, issueRelationInfo)
		}
	}
	for i := range issueRelationInfos {
		pullRequests := make([]entity.PullRequest, 0)
		for _, pr := range *(issueRelationInfos[i].PullRequests) {
			if option.BaseBranch == "" || pr.BaseBranch == option.BaseBranch {
				pullRequests = append(pullRequests, pr)
			}
		}
		issueRelationInfos[i].PullRequests = &pullRequests
	}

	return &issueRelationInfos, nil
}

func SelectIssueRelationInfoUnique(option *dto.IssueRelationInfoQuery) (*dto.IssueRelationInfo, error) {
	infos, err := SelectIssueRelationInfo(option)
	if nil != err {
		return nil, err
	}
	if len(*infos) != 1 {
		return nil, errors.New(fmt.Sprintf("more than one issue_relation found: %+v", option))
	}
	return &((*infos)[0]), nil
}

func SaveIssueRelationInfo(issueRelationInfo *dto.IssueRelationInfo) error {

	if issueRelationInfo == nil {
		return nil
	}

	// Save Issue
	if issueRelationInfo.Issue != nil {
		if err := repository.CreateOrUpdateIssue(issueRelationInfo.Issue); nil != err {
			return err
		}
	}

	// Save IssueAffects
	if issueRelationInfo.IssueAffects != nil {
		for _, issueAffect := range *issueRelationInfo.IssueAffects {
			if err := repository.CreateOrUpdateIssueAffect(&issueAffect); nil != err {
				return err
			}
		}
	}

	// Save IssuePrRelations
	if issueRelationInfo.IssuePrRelations != nil {
		for _, issuePrRelation := range *issueRelationInfo.IssuePrRelations {
			if err := repository.CreateIssuePrRelation(&issuePrRelation); nil != err {
				return err
			}
		}
	}

	// Save PullRequests
	for _, pullRequest := range *issueRelationInfo.PullRequests {
		if err := repository.CreateOrUpdatePullRequest(&pullRequest); nil != err {
			return err
		}
	}

	return nil
}

// ============================================================================
// ============================================================================ Inner Function

func ComposeRelationInfoByIssue(issue *entity.Issue) (*dto.IssueRelationInfo, error) {
	// Find IssueAffects
	issueAffects, err := ComposeIssueAffectWithIssueID(issue.IssueID)
	if nil != err {
		return nil, err
	}

	// Find IssuePrRelations
	issuePrRelationOption := &entity.IssuePrRelationOption{
		IssueID: issue.IssueID,
	}
	issuePrRelations, err := repository.SelectIssuePrRelation(issuePrRelationOption)
	if nil != err {
		return nil, err
	}

	// Find PullRequests
	pullRequests := make([]entity.PullRequest, 0)
	for _, issuePrRelation := range *issuePrRelations {
		pullRequestOption := &entity.PullRequestOption{
			PullRequestID: issuePrRelation.PullRequestID,
		}
		pullRequest, err := repository.SelectPullRequestUnique(pullRequestOption)
		if nil != err {
			return nil, err
		}
		pullRequests = append(pullRequests, *pullRequest)
	}

	// Construct IssueRelationInfo
	issueRelationInfo := &dto.IssueRelationInfo{
		Issue:            issue,
		IssueAffects:     issueAffects,
		IssuePrRelations: issuePrRelations,
		PullRequests:     &pullRequests,
	}
	return issueRelationInfo, nil
}
