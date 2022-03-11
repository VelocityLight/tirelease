package service

import (
	"time"

	"tirelease/commons/git"
	"tirelease/internal/entity"
	"tirelease/internal/repository"

	"github.com/google/go-github/v41/github"
)

// Cron Job
func CronRefreshIssuesV4() error {
	// get repos
	repoOption := &entity.RepoOption{}
	repos, err := repository.SelectRepo(repoOption)
	if err != nil {
		return err
	}

	// multi-batch refresh
	fromTimeBefore := -96
	batchLimit := 20
	totalLimit := 500
	for _, repo := range *repos {
		issues, err := git.ClientV4.GetIssuesByTimeRangeV4(
			repo.Owner, repo.Repo, []string{git.BugTypeLabel},
			time.Now().Add(time.Duration(fromTimeBefore)*time.Hour), time.Now(),
			batchLimit, totalLimit)
		if err != nil {
			return err
		}

		for _, issue := range issues {
			issueRelation, err := ComposeIssueRelationInfoByIssueV4(&issue)
			if err != nil {
				return err
			}
			err = SaveIssueRelationInfo(issueRelation)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Git Webhook
// Webhook param only support v3 (v4 has no webhook right now)
func WebhookRefreshIssueV4(issue *github.Issue) error {
	// params
	if issue == nil {
		return nil
	}
	issueID := *issue.NodeID

	// handler
	issueRelationInfo, err := GetIssueRelationInfoByIssueIDV4(issueID)
	if err != nil {
		return err
	}
	if issueRelationInfo == nil {
		return nil
	}
	err = SaveIssueRelationInfo(issueRelationInfo)
	if err != nil {
		return err
	}

	return nil
}
