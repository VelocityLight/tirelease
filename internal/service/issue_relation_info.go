package service

import (
	"time"

	"tirelease/commons/git"
	"tirelease/internal/dto"
	"tirelease/internal/entity"
	"tirelease/internal/repository"
)

// Get All IssueRelationInfo by Issue
func GetIssueRelationInfoByIssueNumber(owner, repo string, number int) (*dto.IssueRelationInfo, error) {
	issue, err := GetIssueByNumberFromV3(owner, repo, number)
	if nil != err {
		return nil, err
	}

	issueAffects, err := ConsistIssueAffectsByIssue(issue)
	if nil != err {
		return nil, err
	}
	issuePrRelations, pullRequests, err := ConsistIssuePrRelationsByIssue(issue)
	if nil != err {
		return nil, err
	}

	triageRelationInfo := &dto.IssueRelationInfo{
		Issue:            issue,
		IssueAffects:     issueAffects,
		IssuePrRelations: issuePrRelations,
		PullRequests:     pullRequests,
	}
	return triageRelationInfo, nil
}

// Save TriageRelationInfo Every Detail
func SaveIssueRelationInfo(triageRelationInfo *dto.IssueRelationInfo) error {
	// Save Issue
	if err := repository.CreateOrUpdateIssue(triageRelationInfo.Issue); nil != err {
		return err
	}

	// Save IssueAffects
	for _, issueAffect := range triageRelationInfo.IssueAffects {
		if err := repository.CreateIssueAffect(issueAffect); nil != err {
			return err
		}
	}

	// Save IssuePrRelations
	for _, issuePrRelation := range triageRelationInfo.IssuePrRelations {
		if err := repository.CreateIssuePrRelation(issuePrRelation); nil != err {
			return err
		}
	}

	// Save PullRequests
	for _, pullRequest := range triageRelationInfo.PullRequests {
		if err := repository.CreateOrUpdatePullRequest(pullRequest); nil != err {
			return err
		}
	}

	return nil
}

// Consist IssueAffects by Issue
func ConsistIssueAffectsByIssue(issue *entity.Issue) ([]*entity.IssueAffect, error) {
	issueAffects := make([]*entity.IssueAffect, 0)
	// todo: minorVersion is fixed now, need to search later - 2022.2.7
	for _, minorVersion := range MinorVersionList {
		issueAffect := &entity.IssueAffect{
			CreateTime:    time.Now(),
			UpdateTime:    time.Now(),
			IssueID:       issue.IssueID,
			AffectVersion: minorVersion,
			AffectResult:  entity.AffectResultResultUnKnown,
		}
		issueAffects = append(issueAffects, issueAffect)
	}
	return issueAffects, nil
}

// Consist IssuePrRelations & PullRequests by Issue
func ConsistIssuePrRelationsByIssue(issue *entity.Issue) ([]*entity.IssuePrRelation, []*entity.PullRequest, error) {
	// Query timeline
	issueWithTimeline, err := git.ClientV4.GetIssueByNumber(issue.Owner, issue.Repo, issue.Number)
	if nil != err {
		return nil, nil, err
	}
	edges := issueWithTimeline.TimelineItems.Edges
	if nil == edges || len(edges) == 0 {
		return nil, nil, nil
	}

	// Analyze timeline to consist IssuePrRelations & PullRequests
	issuePrRelations := make([]*entity.IssuePrRelation, 0)
	pullRequests := make([]*entity.PullRequest, 0)
	for _, edge := range edges {
		if nil == &edge.Node || nil == &edge.Node.CrossReferencedEvent ||
			nil == &edge.Node.CrossReferencedEvent.Source || nil == &edge.Node.CrossReferencedEvent.Source.PullRequest {
			continue
		}
		if git.CrossReferencedEvent != edge.Node.Typename {
			continue
		}
		pr := edge.Node.CrossReferencedEvent.Source.PullRequest
		if nil == pr.ID {
			continue
		}

		var issuePrRelation = &entity.IssuePrRelation{
			CreateTime:    time.Now(),
			UpdateTime:    time.Now(),
			IssueID:       issueWithTimeline.ID.(string),
			PullRequestID: pr.ID.(string),
		}
		pullRequest, err := GetPullRequestByNumberFromV3(
			string(pr.Repository.Owner.Login), string(pr.Repository.Name), int(pr.Number))
		if nil != err {
			return nil, nil, err
		}
		issuePrRelations = append(issuePrRelations, issuePrRelation)
		pullRequests = append(pullRequests, pullRequest)
	}

	// Return
	return issuePrRelations, pullRequests, nil
}
