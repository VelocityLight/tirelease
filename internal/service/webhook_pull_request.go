package service

import (
	"time"

	"tirelease/commons/git"
	"tirelease/internal/entity"
	"tirelease/internal/repository"

	"github.com/google/go-github/v41/github"
)

// Cron Job
func CronRefreshPullRequestV4() error {
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
		prs, err := git.ClientV4.GetPullRequestsFromV4(
			repo.Owner, repo.Repo,
			time.Now().Add(time.Duration(fromTimeBefore)*time.Hour),
			batchLimit, totalLimit)
		if err != nil {
			return err
		}

		for _, pr := range prs {
			err = repository.CreateOrUpdatePullRequest(entity.ComposePullRequestFromV4(&pr))
			if err != nil {
				return err
			}
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
