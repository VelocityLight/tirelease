package service

import (
	"time"

	"tirelease/commons/git"
	"tirelease/internal/dto"
	"tirelease/internal/entity"
)

// ================================================================ Compose Function From Remote Query
func GetIssueRelationInfoByIssueNumber(owner, repo string, number int) (*dto.IssueRelationInfo, error) {
	issue, err := GetIssueByNumberFromV3(owner, repo, number)
	if nil != err {
		return nil, err
	}

	return ComposeTriageRelationInfoByIssue(issue)
}

func ComposeTriageRelationInfoByIssue(issue *entity.Issue) (*dto.IssueRelationInfo, error) {
	issueAffects, err := ComposeIssueAffectsByIssue(issue)
	if nil != err {
		return nil, err
	}
	issuePrRelations, pullRequests, err := ComposeIssuePrRelationsByIssue(issue)
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

func ComposeIssueAffectsByIssue(issue *entity.Issue) (*[]entity.IssueAffect, error) {
	issueAffects := make([]entity.IssueAffect, 0)
	// todo: minorVersion is fixed now, need to search later - 2022.2.7
	for _, minorVersion := range MinorVersionList {
		issueAffect := &entity.IssueAffect{
			CreateTime:    time.Now(),
			UpdateTime:    time.Now(),
			IssueID:       issue.IssueID,
			AffectVersion: minorVersion,
			AffectResult:  entity.AffectResultResultUnKnown,
		}
		issueAffects = append(issueAffects, *issueAffect)
	}
	return &issueAffects, nil
}

func ComposeIssuePrRelationsByIssue(issue *entity.Issue) (*[]entity.IssuePrRelation, *[]entity.PullRequest, error) {
	// Query timeline
	issueWithTimeline, err := git.ClientV4.GetIssueByNumber(issue.Owner, issue.Repo, issue.Number)
	if nil != err {
		return nil, nil, err
	}
	edges := issueWithTimeline.TimelineItems.Edges
	if nil == edges || len(edges) == 0 {
		return nil, nil, nil
	}

	// Analyze timeline to compose IssuePrRelations & PullRequests
	issuePrRelations := make([]entity.IssuePrRelation, 0)
	pullRequests := make([]entity.PullRequest, 0)
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
		issuePrRelations = append(issuePrRelations, *issuePrRelation)
		pullRequests = append(pullRequests, *pullRequest)
	}

	// Return
	return &issuePrRelations, &pullRequests, nil
}
