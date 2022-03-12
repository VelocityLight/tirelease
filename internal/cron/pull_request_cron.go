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
	cron.Create("* */1 * * * *", func() { service.CronRefreshPullRequestV4(repos) })
}
