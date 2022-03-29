package service

import (
	"fmt"
	"strconv"
	"strings"

	"tirelease/internal/dto"
	"tirelease/internal/entity"
	"tirelease/internal/repository"

	"github.com/pkg/errors"
)

// ============================================================================
// ============================================================================ CURD Of IssueRelationInfo
func SelectIssueRelationInfo(option *dto.IssueRelationInfoQuery) (*[]dto.IssueRelationInfo, error) {
	// select join
	joins, err := repository.SelectIssueRelationInfoByJoin(option)
	if nil != err {
		return nil, err
	}

	// compose
	issueRelationInfos := make([]dto.IssueRelationInfo, 0)
	for _, join := range *joins {
		issueRelationInfo := &dto.IssueRelationInfo{}

		// issue
		issueOption := &entity.IssueOption{
			IssueID: join.IssueID,
		}
		issue, err := repository.SelectIssueUnique(issueOption)
		if nil != err {
			return nil, err
		}
		issueRelationInfo.Issue = issue

		// issue_affects
		if join.IssueAffectIDs != "" {
			idList := make([]int64, 0)
			ids := strings.Split(join.IssueAffectIDs, ",")
			for _, id := range ids {
				idint, _ := strconv.Atoi(id)
				idList = append(idList, int64(idint))
			}

			issueAffectOption := &entity.IssueAffectOption{
				IDs: idList,
			}
			issueAffects, err := repository.SelectIssueAffect(issueAffectOption)
			if nil != err {
				return nil, err
			}
			issueRelationInfo.IssueAffects = issueAffects
		}

		// issue_pr_relations
		issuePrRelationOption := &entity.IssuePrRelationOption{
			IssueID: issue.IssueID,
		}
		issuePrRelations, err := repository.SelectIssuePrRelation(issuePrRelationOption)
		if nil != err {
			return nil, err
		}
		issueRelationInfo.IssuePrRelations = issuePrRelations

		// prs
		if issuePrRelations != nil && len(*issuePrRelations) > 0 {
			pullRequestIDs := make([]string, 0)
			for _, issuePrRelation := range *issuePrRelations {
				pullRequestIDs = append(pullRequestIDs, string(issuePrRelation.PullRequestID))
			}

			pullRequestOption := &entity.PullRequestOption{
				PullRequestIDs: pullRequestIDs,
			}
			if option.BaseBranch != "" {
				pullRequestOption.BaseBranch = option.BaseBranch
			}
			pullRequests, err := repository.SelectPullRequest(pullRequestOption)
			if nil != err {
				return nil, err
			}
			issueRelationInfo.PullRequests = pullRequests
		}

		// return
		issueRelationInfos = append(issueRelationInfos, *issueRelationInfo)
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
	if issueRelationInfo.PullRequests != nil {
		for _, pullRequest := range *issueRelationInfo.PullRequests {
			if err := repository.CreateOrUpdatePullRequest(&pullRequest); nil != err {
				return err
			}
		}
	}

	return nil
}
