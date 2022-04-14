package service

import (
	"time"

	"tirelease/commons/git"
	"tirelease/internal/entity"
	"tirelease/internal/repository"

	"github.com/google/go-github/v41/github"
)

// Cron Job
type RefreshPullRequestParams struct {
	Repos       *[]entity.Repo `json:"repos"`
	BeforeHours int64          `json:"before_hours"`
	Batch       int            `json:"batch"`
	Total       int            `json:"total"`
}

func CronRefreshPullRequestV4(params *RefreshPullRequestParams) error {
	// get repos
	if params.Repos == nil || len(*params.Repos) == 0 {
		return nil
	}

	// multi-batch refresh
	for _, repo := range *params.Repos {
		request := &git.RemoteIssueRangeRequest{
			Owner:      repo.Owner,
			Name:       repo.Repo,
			From:       time.Now().Add(time.Duration(params.BeforeHours) * time.Hour),
			BatchLimit: params.Batch,
			TotalLimit: params.Total,
		}
		prs, err := git.ClientV4.GetPullRequestsFromV4(request)
		if err != nil {
			return err
		}

		for i := range prs {
			err = repository.CreateOrUpdatePullRequest(entity.ComposePullRequestFromV4(&(prs[i])))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func CronMergeRetryPullRequestV3() error {
	// select no merge PRs
	merged := false
	cherryPickApproved := true
	alreadyReviewed := true
	option := &entity.PullRequestOption{
		State:              git.OpenStatus,
		Merged:             &merged,
		MergeableState:     &git.MergeableStateMergeable,
		CherryPickApproved: &cherryPickApproved,
		AlreadyReviewed:    &alreadyReviewed,
	}
	prs, err := repository.SelectPullRequest(option)
	if err != nil {
		return err
	}

	// retry
	for _, pr := range *prs {
		_, _, err := git.Client.CreateCommentByNumber(pr.Owner, pr.Repo, pr.Number, git.MergeRetryComment)
		if err != nil {
			return err
		}
	}
	return nil
}

// Git Webhook
// Webhook param only support v3 (v4 has no webhook right now)
func WebhookRefreshPullRequestV3(pr *github.PullRequest) error {
	// params
	if pr == nil {
		return nil
	}

	// handler
	err := repository.CreateOrUpdatePullRequest(entity.ComposePullRequestFromV3(pr))
	if err != nil {
		return err
	}

	return nil
}

func WebHookRefreshPullRequestRefIssue(pr *github.PullRequest) error {
	// params
	if pr == nil {
		return nil
	}
	pullRequestID := *(pr.NodeID)
	if pullRequestID == "" {
		return nil
	}

	// handler
	prWithTimeline, err := git.ClientV4.GetPullRequestByID(pullRequestID)
	if err != nil {
		return err
	}
	edges := prWithTimeline.TimelineItems.Edges
	if nil == edges || len(edges) == 0 {
		return nil
	}
	for i := range edges {
		edge := edges[i]
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

		// find pr relation issue and refresh
		option := &entity.IssuePrRelationOption{
			PullRequestID: pr.ID.(string),
		}
		issuePrRelations, err := repository.SelectIssuePrRelation(option)
		if err != nil {
			return err
		}
		if nil == issuePrRelations || len(*issuePrRelations) == 0 {
			continue
		}
		for _, issuePrRelation := range *issuePrRelations {
			err := WebhookRefreshIssueV4ByIssueID(issuePrRelation.IssueID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
