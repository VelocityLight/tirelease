package service

import (
	"strings"
	"time"
	"tirelease/commons/git"
	"tirelease/internal/entity"
	"tirelease/internal/repository"

	"github.com/google/go-github/v41/github"
	"github.com/shurcooL/githubv4"
)

func UpdatePrAndIssue(webhookPayload WebhookPayload) error {
	if webhookPayload.Issue != nil {
		issue, err := git.ClientV4.GetIssueByNumber(webhookPayload.Repository.Owner.Login, webhookPayload.Repository.Name, webhookPayload.Issue.Number)
		// create issue affect
		for _, minorVersion := range MinorVersionList {
			issueAffect := entity.IssueAffect{
				CreateTime:    time.Now(),
				UpdateTime:    time.Now(),
				IssueID:       issue.ID.(string),
				AffectVersion: minorVersion,
				AffectResult:  entity.AffectResultResultUnKnown,
			}
			if err := repository.CreateIssueAffect(&issueAffect); err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
		if err := repository.CreateOrUpdateIssue(IssueNodeToIssue(issue)); err != nil {
			return err
		}
	}
	if webhookPayload.PullRequest != nil {
		pullRequest, err := git.ClientV4.GetPullRequestsByNumber(webhookPayload.Repository.Owner.Login, webhookPayload.Repository.Name, webhookPayload.PullRequest.Number)
		if err != nil {
			return err
		}
		if err := repository.CreateOrUpdatePullRequest(PullRequestNodeToPullRequest(pullRequest)); err != nil {
			return err
		}
	}
	return nil
}

func IssueNodeToIssue(issueNode *git.IssueNode) *entity.Issue {
	labels := &[]github.Label{}
	for _, labelNode := range issueNode.Labels.Nodes {
		label := github.Label{}
		label.Name = github.String(string(labelNode.Name))
		*labels = append(*labels, label)
	}

	assignees := &[]github.User{}
	for _, userNode := range issueNode.Assignees.Nodes {
		user := github.User{
			Login:     (*string)(&userNode.Login),
			CreatedAt: (*github.Timestamp)(&userNode.CreatedAt),
		}
		*assignees = append(*assignees, user)
	}

	closedByPrID := ""
	if issueNode.State == githubv4.IssueStateClosed {
		for _, edge := range issueNode.TimelineItems.Edges {
			closer := edge.Node.ClosedEvent.Closer.PullRequest
			if closer.Number != 0 {
				closedByPrID = closer.ID.(string)
			}
		}
	}

	resp := &entity.Issue{
		IssueID:               issueNode.ID.(string),
		Number:                int(issueNode.Number),
		State:                 string(issueNode.State),
		Title:                 string(issueNode.Title),
		Repo:                  strings.Join([]string{string(issueNode.Repository.Owner.Login), string(issueNode.Repository.Name)}, "/"),
		ClosedAt:              issueNode.ClosedAt.Time,
		CreatedAt:             issueNode.CreatedAt.Time,
		UpdatedAt:             issueNode.UpdatedAt.Time,
		Labels:                labels,
		Assignees:             assignees,
		ClosedByPullRequestID: closedByPrID,
	}
	return resp
}

func PullRequestNodeToPullRequest(pullRequestNode *git.PullRequest) *entity.PullRequest {
	labels := &[]github.Label{}
	for _, labelNode := range pullRequestNode.Labels.Nodes {
		label := github.Label{}
		label.Name = github.String(string(labelNode.Name))
		*labels = append(*labels, label)
	}
	assignees := &[]github.User{}
	for _, userNode := range pullRequestNode.Assignees.Nodes {
		user := github.User{
			Login:     (*string)(&userNode.Login),
			CreatedAt: (*github.Timestamp)(&userNode.CreatedAt),
		}
		*assignees = append(*assignees, user)
	}
	SourcePrID := ""
	for _, edge := range pullRequestNode.TimelineItems.Edges {
		sourcePr := edge.Node.CrossReferencedEvent.Source.PullRequest
		if sourcePr.Number != 0 && strings.HasPrefix(string(pullRequestNode.Title), string(sourcePr.Title)) {
			SourcePrID = sourcePr.ID.(string)
		}
	}

	resp := &entity.PullRequest{
		PullRequestID:       pullRequestNode.ID.(string),
		Number:              int(pullRequestNode.Number),
		State:               string(pullRequestNode.State),
		Title:               string(pullRequestNode.Title),
		Repo:                strings.Join([]string{string(pullRequestNode.Repository.Owner.Login), string(pullRequestNode.Repository.Name)}, "/"),
		HeadBranch:          string(pullRequestNode.HeadRefName),
		MergedAt:            pullRequestNode.MergedAt.Time,
		CreatedAt:           pullRequestNode.CreatedAt.Time,
		UpdatedAt:           pullRequestNode.UpdatedAt.Time,
		Merged:              bool(pullRequestNode.Merged),
		Labels:              labels,
		Assignees:           assignees,
		SourcePullRequestID: SourcePrID,
	}
	return resp
}

type WebhookPayload struct {
	PullRequest *PullRequest `json:"pull_request,omitempty"`
	Issue       *Issue       `json:"issue,omitempty"`
	Repository  *Repository  `json:"repository,omitempty"`
}

type PullRequest struct {
	ID     int64 `json:"id"`
	Number int   `json:"number"`
}

type Issue struct {
	ID     int64 `json:"id"`
	Number int   `json:"number"`
}

type Repository struct {
	Name  string `json:"name"`
	Owner Owner  `json:"owner"`
}

type Owner struct {
	Login string `json:"login"`
}
