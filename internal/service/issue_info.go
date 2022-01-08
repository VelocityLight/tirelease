package service

import (
	"strings"
	"time"
	"tirelease/internal/entity"
	"tirelease/internal/repository"

	"github.com/google/go-github/v41/github"
)

var MinorVersionList = []string{"4.0", "5.0", "5.1", "5.2", "5.3", "5.4"}

func ListIssueInfo() ([]*IssueInfo, error) {
	resp := []*IssueInfo{}
	status := entity.ReleaseVersionStatusOpen
	typeV := entity.ReleaseVersionTypePatch
	optionRV := entity.ReleaseVersionOption{
		Status: &status,
		Type:   &typeV,
	}
	patchVersions, err := repository.SelectReleaseVersion(&optionRV)
	if err != nil {
		return resp, err
	}
	minorPatchVersionMap := map[string]string{}
	for _, patchVersion := range *patchVersions {
		minorPatchVersionMap[patchVersion.FatherReleaseVersionName] = patchVersion.Name
	}

	optionI := entity.IssueOption{}
	issues, err := repository.SelectIssue(&optionI)
	if err != nil {
		return resp, err
	}

	for _, issue := range *issues {
		if issue.Labels == nil {
			issue.Labels = &[]github.Label{}
		}
		if issue.Assignees == nil {
			issue.Assignees = &[]github.User{}
		}
		if issue.ClosedAt == nil {
			issue.ClosedAt = &time.Time{}
		}
		affects, err := ListAffected(issue.IssueID, issue.ClosedByPullRequestID, minorPatchVersionMap)
		if err != nil {
			return resp, err
		}

		issueInfo := &IssueInfo{
			Number:    issue.Number,
			Title:     issue.Title,
			Url:       issue.HTMLURL,
			CreatedAt: issue.CreatedAt,
			ClosedAt:  *issue.ClosedAt,
			Severity:  getIssueType(*issue.Labels),
			State:     issue.State,
			Type:      getIssueSeverity(*issue.Labels),
			Assignee:  getAssignee(*issue.Assignees),
			Labels:    []string{},
			Affects:   affects,
		}
		resp = append(resp, issueInfo)
	}
	return resp, nil
}

func ListAffected(issueID string, closedPrID string, minorPatchVersionMap map[string]string) ([]*Affect, error) {
	resp := []*Affect{}
	issueAffects, err := repository.SelectIssueAffect(&entity.IssueAffectOption{IssueID: issueID})
	if err != nil {
		return resp, err
	}
	versionTraiges, err := repository.SelectVersionTriage(&entity.VersionTriageOption{IssueID: issueID})
	if err != nil {
		return resp, err
	}
	cherryPickToPrs := []entity.PullRequest{}
	if closedPrID != "" {
		results, err := repository.SelectPullRequest(&entity.PullRequestOption{SourcePullRequestID: closedPrID})
		if err != nil {
			return resp, err
		}
		cherryPickToPrs = *results
	}
	for _, issueAffect := range *issueAffects {
		triagestatus := ""
		for _, versionTraige := range *versionTraiges {
			if versionTraige.VersionName == minorPatchVersionMap[issueAffect.AffectVersion] {
				triagestatus = string(versionTraige.TriageResult)
				break
			}
		}
		release := Release{
			BaseVersion:  issueAffect.AffectVersion,
			TriageStatus: triagestatus,
			Patch:        minorPatchVersionMap[issueAffect.AffectVersion],
		}
		pr := Pr{}
		for _, cpr := range cherryPickToPrs {
			if cpr.MergedAt == nil {
				cpr.MergedAt = &time.Time{}
			}
			if strings.HasSuffix(cpr.HeadBranch, issueAffect.AffectVersion) {
				pr = Pr{
					Repository:  cpr.Repo,
					Number:      cpr.Number,
					Title:       cpr.Title,
					Url:         cpr.HTMLURL,
					State:       cpr.State,
					MergeTarget: cpr.HeadBranch,
					MergedAt:    *cpr.MergedAt,
				}
			}
		}

		affect := &Affect{
			Version: issueAffect.AffectVersion,
			Affect:  string(issueAffect.AffectResult),
			Release: &release,
			Pr:      &pr,
		}
		resp = append(resp, affect)
	}
	return resp, nil
}

func getIssueType(labels []github.Label) string {
	resp := ""
	for _, label := range labels {
		if strings.HasPrefix(*label.Name, "type/") {
			resp = strings.ReplaceAll(*label.Name, "type/", "")
			return resp
		}
	}
	return resp
}

func getIssueSeverity(labels []github.Label) string {
	resp := ""
	for _, label := range labels {
		if strings.HasPrefix(*label.Name, "severity/") {
			resp = strings.ReplaceAll(*label.Name, "severity/", "")
			return resp
		}
	}
	return resp
}

func getAssignee(users []github.User) string {
	resp := ""
	for _, user := range users {
		resp = *user.Name
	}
	return resp
}

type IssueInfo struct {
	Number    int
	Title     string
	Url       string
	CreatedAt time.Time
	ClosedAt  time.Time
	Severity  string
	State     string
	Type      string
	Assignee  string
	Labels    []string
	Affects   []*Affect
}

type Affect struct {
	Version string
	Affect  string
	Pr      *Pr
	Release *Release
}

type Pr struct {
	Repository  string
	Number      int
	Title       string
	Url         string
	State       string
	MergeTarget string
	MergedAt    time.Time
	Commit      string
	Refs        string
}

type Release struct {
	BaseVersion  string
	TriageStatus string
	Patch        string
}
