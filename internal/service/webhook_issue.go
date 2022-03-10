package service

import (
	"time"

	"tirelease/internal/entity"
	"tirelease/internal/repository"

	"tirelease/commons/git"
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
			repo.Owner, repo.Repo, []string{git.BugLabel},
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
func WebhookRefreshIssueV4() error {
	return nil
}
