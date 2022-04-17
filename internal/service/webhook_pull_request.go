package service

import (
	"strings"
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

	// handler approve later

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

	// find close or ref issue numbers
	prV4, err := git.ClientV4.GetPullRequestByID(pullRequestID)
	if err != nil {
		return err
	}
	baseBranch := prV4.BaseRefName
	if baseBranch == "" || !strings.HasPrefix(string(baseBranch), git.ReleaseBranchPrefix) {
		return nil
	}
	issueNumbers, err := GetPullRequestRefIssuesByRegexFromV4(prV4)
	if err != nil {
		return err
	}

	// refresh cross-referenced issue
	if len(issueNumbers) > 0 {
		for _, issueNumber := range issueNumbers {
			issueOption := &entity.IssueOption{
				Number: issueNumber,
			}
			issues, err := repository.SelectIssue(issueOption)
			if err != nil {
				return err
			}
			if len(*issues) == 0 {
				continue
			}

			for _, issue := range *issues {
				err := WebhookRefreshIssueV4ByIssueID(issue.IssueID)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
