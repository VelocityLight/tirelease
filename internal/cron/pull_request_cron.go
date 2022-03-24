package cron

import (
	"tirelease/commons/cron"
	"tirelease/internal/entity"
	"tirelease/internal/repository"
	"tirelease/internal/service"
)

func PullRequestCron() {
	// Cron 表达式及功能方法
	repoOption := &entity.RepoOption{}
	repos, err := repository.SelectRepo(repoOption)
	if err != nil {
		return
	}
	params := &service.RefreshPullRequestParams{
		Repos:       repos,
		BeforeHours: -8760,
		Batch:       20,
		Total:       12000,
	}
	cron.Create("* */1 * * * *", func() { service.CronRefreshPullRequestV4(params) })

	cron.Create("* */2 * * * *", func() { service.CronMergeRetryPullRequestV3() })
}
