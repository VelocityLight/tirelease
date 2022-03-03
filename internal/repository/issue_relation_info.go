package repository

import (
	"tirelease/internal/dto"
	"tirelease/internal/entity"
)

// ============================================================================ CURD Of IssueRelationInfo
func SaveIssueRelationInfo(triageRelationInfo *dto.IssueRelationInfo) error {
	// Save Issue
	if err := CreateOrUpdateIssue(triageRelationInfo.Issue); nil != err {
		return err
	}

	// Save IssueAffects
	for _, issueAffect := range *triageRelationInfo.IssueAffects {
		if err := CreateIssueAffect(&issueAffect); nil != err {
			return err
		}
	}

	// Save IssuePrRelations
	for _, issuePrRelation := range *triageRelationInfo.IssuePrRelations {
		if err := CreateIssuePrRelation(&issuePrRelation); nil != err {
			return err
		}
	}

	// Save PullRequests
	for _, pullRequest := range *triageRelationInfo.PullRequests {
		if err := CreateOrUpdatePullRequest(&pullRequest); nil != err {
			return err
		}
	}

	return nil
}

func SelectIssueRelationInfo(option *dto.IssueRelationInfoQuery) (*[]dto.IssueRelationInfo, error) {
	// Select Issues
	issueOption := &entity.IssueOption{
		ID:      option.ID,
		IssueID: option.IssueID,
		Number:  option.Number,
		State:   option.State,
		Owner:   option.Owner,
		Repo:    option.Repo,
	}
	issues, err := SelectIssue(issueOption)
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

	return &issueRelationInfos, nil
}

// ============================================================================ Inner Function
func ComposeRelationInfoByIssue(issue *entity.Issue) (*dto.IssueRelationInfo, error) {
	// Find IssueAffects
	issueAffectOption := &entity.IssueAffectOption{
		IssueID: issue.IssueID,
	}
	issueAffects, err := SelectIssueAffect(issueAffectOption)
	if nil != err {
		return nil, err
	}

	// Find IssuePrRelations
	issuePrRelationOption := &entity.IssuePrRelationOption{
		IssueID: issue.IssueID,
	}
	issuePrRelations, err := SelectIssuePrRelation(issuePrRelationOption)
	if nil != err {
		return nil, err
	}

	// Find PullRequests
	pullRequests := make([]entity.PullRequest, 0)
	for _, issuePrRelation := range *issuePrRelations {
		pullRequestOption := &entity.PullRequestOption{
			PullRequestID: issuePrRelation.PullRequestID,
		}
		pullRequest, err := SelectPullRequestUnique(pullRequestOption)
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
