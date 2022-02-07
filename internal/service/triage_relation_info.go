package service

import (
	"time"

	"tirelease/commons/git"
	"tirelease/internal/entity"
	"tirelease/internal/repository"
)

// Consist IssueAffects by Issue
func ConsistIssueAffectsByIssue(issue *entity.Issue) ([]*entity.IssueAffect, error) {
	return nil, nil
}

// Consist IssuePrRelations & PullRequests by Issue
func ConsistIssuePrRelationsByIssue(owner, repo string, number int) ([]*entity.IssuePrRelation, []*entity.PullRequest, error) {
	// Query timeline
	issueWithTimeline, err := git.ClientV4.GetIssueByNumber(owner, repo, number)
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

// Consist All TriageRelationInfo by Issue
func ConsistTriageRelationInfoByIssue(issue *entity.Issue) (*entity.TriageRelationInfo, error) {
	issueAffects, err := ConsistIssueAffectsByIssue(issue)
	if nil != err {
		return nil, err
	}
	issuePrRelations, pullRequests, err := ConsistIssuePrRelationsByIssue(issue.Owner, issue.Repo, issue.Number)
	if nil != err {
		return nil, err
	}

	triageRelationInfo := &entity.TriageRelationInfo{
		Issue:            issue,
		IssueAffects:     issueAffects,
		IssuePrRelations: issuePrRelations,
		PullRequests:     pullRequests,
	}
	return triageRelationInfo, nil
}

// Save TriageRelationInfo Every Detail
func SaveTriageRelationInfo(triageRelationInfo *entity.TriageRelationInfo) error {
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