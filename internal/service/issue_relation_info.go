package service

import (
	"time"

	"tirelease/commons/git"
	"tirelease/internal/dto"
	"tirelease/internal/entity"
	"tirelease/internal/repository"
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
	// search all minor version
	releaseVersionOption := &entity.ReleaseVersionOption{
		Type:   entity.ReleaseVersionTypeMinor,
		Status: entity.ReleaseVersionStatusOpen,
	}
	releaseVersions, err := repository.SelectReleaseVersion(releaseVersionOption)
	if nil != err {
		return nil, err
	}

	// init all affects of issue
	issueAffects := make([]entity.IssueAffect, 0)
	for _, releaseVersion := range *releaseVersions {
		issueAffect := &entity.IssueAffect{
			CreateTime:    time.Now(),
			UpdateTime:    time.Now(),
			IssueID:       issue.IssueID,
			AffectVersion: releaseVersion.Name,
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

// ============================================================================ CURD Of IssueRelationInfo
func SaveIssueRelationInfo(triageRelationInfo *dto.IssueRelationInfo) error {
	// Save Issue
	if err := repository.CreateOrUpdateIssue(triageRelationInfo.Issue); nil != err {
		return err
	}

	// Save IssueAffects
	for _, issueAffect := range *triageRelationInfo.IssueAffects {
		if err := repository.CreateOrUpdateIssueAffect(&issueAffect); nil != err {
			return err
		}
	}

	// Save IssuePrRelations
	for _, issuePrRelation := range *triageRelationInfo.IssuePrRelations {
		if err := repository.CreateIssuePrRelation(&issuePrRelation); nil != err {
			return err
		}
	}

	// Save PullRequests
	for _, pullRequest := range *triageRelationInfo.PullRequests {
		if err := repository.CreateOrUpdatePullRequest(&pullRequest); nil != err {
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
			if option.HeadBranch == "" || pr.HeadBranch == option.HeadBranch {
				pullRequests = append(pullRequests, pr)
			}
		}
		issueRelationInfos[i].PullRequests = &pullRequests
	}

	return &issueRelationInfos, nil
}

// ============================================================================ Inner Function
func ComposeRelationInfoByIssue(issue *entity.Issue) (*dto.IssueRelationInfo, error) {
	// Find IssueAffects
	issueAffectOption := &entity.IssueAffectOption{
		IssueID: issue.IssueID,
	}
	issueAffects, err := repository.SelectIssueAffect(issueAffectOption)
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
